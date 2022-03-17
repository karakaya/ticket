FROM golang:1.17-alpine AS builder

WORKDIR /src

COPY go.sum go.mod ./
COPY . .

RUN cp .env.example .env

RUN go mod download
RUN go build -o /bin/app

FROM alpine
COPY --from=builder /bin/app /bin/app

EXPOSE 8080

ENTRYPOINT ["bin/app"]
