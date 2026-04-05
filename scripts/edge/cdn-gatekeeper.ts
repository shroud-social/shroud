import * as BunnySDK from "https://esm.sh/@bunny.net/edgescript-sdk@0.11.2";
import { jwtVerify } from "https://esm.sh/jose";

/**
* @file Bunny EdgeScript for JWT-secured file uploads
 * * Required environment variables:
 * - UPLOAD_JWT_SECRET: HMAC secret for HS256 verification
 * - STORAGE_ZONE: Name of the Bunny storage zone
 * - STORAGE_API_KEY: Api key for above storage zone.
 * - COLD_STORAGE_ZONE: Name of cold Bunny storage zone
 * - COLD_STORAGE_API_KEY: Api key for above storage zone
*/

/**
 * Retries GET requests to try cold storage if file does not exist on hot storage
 * Intercepts PUT requests to verify JWT and authorized path
 * and rewrites the request to Bunny Storage API
 * @param {Object} context - Edgescript context object
 * @param {Request} context.request - Incoming HTTP request
 * @returns {Promise<Response|Request>} - Handled response of original request
 * @throws {Response} 401 if JWT is invalid or missing
 * @throws {Response} 403 if JWT encoded path doesn't match url
 * @throws {Repsonse} 400 if the request fails to rewrite
 */
async function onOriginRequest(context: { request: Request }): Promise<Response> | Response | Promise<Request> | Request | void {
    if (context.request.method == "GET") {
        let response = await fetch(context.request);
        if (response.status == 404) {
            const url = new URL(context.request.url);
            const coldStorageName = Deno.env.get("COLD_STORAGE_ZONE");
            const coldStorageKey = Deno.env.get("COLD_STORAGE_API_KEY");
            const coldStorageUrl = `https://storage.bunnycdn.com/${coldStorageName}${url.pathname}`;
            const tryCold = new Request(coldStorageUrl, {
                method: "GET",
                headers: {
                    "AccessKey": coldStorageKey
                }
            });
            response = await fetch(tryCold);
        }
        return response;
    } else if (context.request.method == "PUT") {
        const encoder = new TextEncoder();
        const secret = encoder.encode(Deno.env.get("UPLOAD_JWT_SECRET"));
        const auth = context.request.headers.get('Authorization');
        if (!auth || !auth.startsWith("Bearer ")) return new Response ('Unauthorized', { status: 401 });
        try {
            const token = auth.substring(7);
            const url = new URL(context.request.url);
            const { payload } = await jwtVerify(token, secret);
            if (payload.path !== url.pathname) return new Response('Wrong Path', { status: 403 });
            try {
                let bytesUploaded = 0;
                const sizeLimit = payload.sizeLimit;
                const stream = new TransformStream({
                    transform(chunk, controller) {
                        bytesUploaded += chunk.byteLength;
                        if (bytesUploaded as number > sizeLimit) {
                            controller.error(new Error("LIMIT_EXCEEDED")); // Currently doesn't reach the try catch block, appears as rewrite error.
                        } else {
                            controller.enqueue(chunk);
                        }
                    }
                })

                const storageZoneName = Deno.env.get("STORAGE_ZONE");
                const storageZoneKey = Deno.env.get("STORAGE_API_KEY");
                const storageUrl = `https://storage.bunnycdn.com/${storageZoneName}${url.pathname}`;
                const newHeaders = new Headers(context.request.headers);
                newHeaders.delete("Host");
                newHeaders.set("AccessKey", storageZoneKey);
                try {
                    const storageRequest = await fetch(storageUrl, {
                        method: "PUT",
                        headers: newHeaders,
                        body: context.request.body.pipeThrough(stream)
                    });
                    return storageRequest
                } catch (e) {
                    if (String(e).includes("LIMIT_EXCEEDED")) {
                        console.log("Upload error: " + e);
                        return new Response('File too large', { status: 413 });
                    }
                    throw e
                };
            } catch (e) {
                console.log("Rewrite Error: " + e);
                return new Response('Upload Failed', { status: 400 });
            }
        } catch (e) {
            console.log("Authorization error: " + e);
            return new Response('Unauthorized', { status: 401 });
        }

    }
    return new Response('Method Not Allowed', { status: 405 });
}

BunnySDK.net.http.servePullZone()
    .onOriginRequest(onOriginRequest);