FROM golang:1.20-alpine as buildbase

WORKDIR /go/src/github.com/rarimo/broadcaster-svc
COPY vendor .
COPY . .

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"

RUN apk add build-base
RUN go build -o /usr/local/bin/broadcaster-svc github.com/rarimo/broadcaster-svc

###

FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/broadcaster-svc /usr/local/bin/broadcaster-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["broadcaster-svc"]

