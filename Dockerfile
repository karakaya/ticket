FROM golang:1.19-alpine AS builder

WORKDIR /src

COPY go.sum go.mod ./
COPY . .

RUN go mod download
RUN go build cmd/main.go 

FROM alpine
COPY --from=builder /src/main /bin/app

EXPOSE 8080

ENTRYPOINT ["bin/app"]
