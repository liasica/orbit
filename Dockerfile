FROM golang:1-alpine AS builder

RUN apk --no-cache --no-progress add make git

WORKDIR /go/orbit
COPY . .
RUN go mod download && make build

FROM alpine:latest
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/* \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /app

WORKDIR /app

EXPOSE 80

ENV ADDRESS=":80"

COPY --from=builder /go/orbit/build/orbit /app/orbit

ENTRYPOINT ["/bin/sh", "-c", "/app/orbit server run --config ./configs/config.yaml --address ${ADDRESS}"]
