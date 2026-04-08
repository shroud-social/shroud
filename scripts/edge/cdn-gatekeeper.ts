import * as BunnySDK from "https://esm.sh/@bunny.net/edgescript-sdk@0.11.2";
import { jwtVerify, SignJWT } from "https://esm.sh/jose";

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

    const coldStorageName = Deno.env.get("COLD_STORAGE_ZONE");
    const coldStorageKey = Deno.env.get("COLD_STORAGE_API_KEY");
    const uploadJwtSecret = Deno.env.get("UPLOAD_JWT_SECRET");
    const storageZoneName = Deno.env.get("STORAGE_ZONE");
    const storageZoneKey = Deno.env.get("STORAGE_API_KEY");

    if (context.request.method == "GET") {
        let response = await fetch(context.request);
        if (response.status == 404) {
            const url = new URL(context.request.url);
            const coldStorageUrl = `https://storage.bunnycdn.com/${coldStorageName}${url.pathname}`;
            response = await fetch(new Request(coldStorageUrl, {
                method: "GET",
                headers: {
                    "AccessKey": coldStorageKey
                }
            }));
        }
        return response;
    }

    if (context.request.method == "PUT") {
        const encoder = new TextEncoder();
        const secret = encoder.encode(uploadJwtSecret);

        const auth = context.request.headers.get('Authorization');
        if (!auth || !auth.startsWith("Bearer ")) return new Response ('Unauthorized', { status: 401 });

        let payload: any;
        let url: any;
        try {
            const token = auth.substring(7);
            url = new URL(context.request.url);
            const result = await jwtVerify(token, secret);
            payload = result.payload
        } catch (e) {
            console.log("Authorization error: " + e);
            return new Response('Unauthorized', { status: 401 });
        }

        if (payload.path !== url.pathname) {
            return new Response('Wrong Path', { status: 403 });
        }

        const abortController = new AbortController();
        const size = payload.size;
        let bytesUploaded = 0;
        const stream = new TransformStream({
            transform(chunk, controller) {
                bytesUploaded += chunk.byteLength;
                if (bytesUploaded as number > size) {
                    abortController.abort("LIMIT_EXCEEDED");
                    controller.error(new Error("LIMIT_EXCEEDED"));
                } else {
                    controller.enqueue(chunk);
                }
            }
        });

        const storageUrl = `https://storage.bunnycdn.com/${storageZoneName}${url.pathname}`;
        const newHeaders = new Headers(context.request.headers)
        newHeaders.delete("Host");
        newHeaders.delete("Authorization");
        newHeaders.set("AccessKey", storageZoneKey);
        newHeaders.set("Checksum", String(payload.hash).toUpperCase());

        try {
            const storageRequest = await fetch(storageUrl, {
                method: "PUT",
                headers: newHeaders,
                body: context.request.body.pipeThrough(stream),
                signal: abortController.signal
            });
            if (storageRequest.ok) {
                const receipt = await new SignJWT({
                    upload_type: payload.upload_type,
                    guild_id: payload.guild_id,
                    channel_id: payload.channel_id,
                    file_name: payload.file_name,
                    size: payload.size,
                    hash: payload.hash,
                    upload_id: payload.upload_id,
                    user_id: payload.user_id,
                    path: payload.path
                })
                    .setProtectedHeader({ alg: 'HS256' })
                    .setIssuedAt()
                    .setExpirationTime('10m')
                    .sign(secret);
                return new Response(JSON.stringify({ receipt }), { status: 200 })
            }
            return storageRequest;
        } catch (e) {
            if (String(e).includes("LIMIT_EXCEEDED")) {
                console.log("Upload error: " + e);
                return new Response('File too large', { status: 413 });
            }
            console.log("Rewrite Error: " + e);
            return new Response('Upload Failed', { status: 400 });
        }
    }

    return new Response('Method Not Allowed', { status: 405 });
}

BunnySDK.net.http.servePullZone()
    .onOriginRequest(onOriginRequest);