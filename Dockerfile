FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY . .
RUN cp .env.example .env
RUN go mod download
RUN CGO_ENABLED=0 go build -o main .
#RUN apk add upx
#RUN upx --ultra-brute /app/main
FROM alpine:3.14.0
COPY --from=builder /app .

CMD ["./main"]