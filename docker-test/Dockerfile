FROM golang:1.23.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk update && \
    apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["./main"]
