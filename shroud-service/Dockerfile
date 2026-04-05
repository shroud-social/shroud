FROM golang:1.26-alpine AS builder

ARG SERVICE_NAME

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o binary ./shroud-service/cmd/${SERVICE_NAME}/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/binary .

CMD ["./binary"]