FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

ENTRYPOINT ["/app/main"]
