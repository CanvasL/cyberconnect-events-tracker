FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 1

RUN apk update --no-cache && apk add --no-cache tzdata && apk add --no-cache build-base  && apk add --no-cache curl

WORKDIR /build

COPY . /build
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/cyberconnect-events-tracker /build/

FROM alpine

RUN apk add --no-cache curl

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/cyberconnect-events-tracker /app/cyberconnect-events-tracker

CMD ["./cyberconnect-events-tracker"]
