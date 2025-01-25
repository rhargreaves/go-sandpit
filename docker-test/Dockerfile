FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod init github.com/rhargreaves/go-sandpit/docker-test && \
	go mod tidy && \
	go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]