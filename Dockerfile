# Builder
FROM golang:alpine as builder

WORKDIR /go/src/github.com/twitter-clone
RUN apk add --no-cache make gcc musl-dev linux-headers git make ca-certificates
COPY . .
RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata ca-certificates && \
    mkdir /app

WORKDIR /app
COPY properties.json /app

EXPOSE 8090

COPY --from=builder /go/src/github.com/twitter-clone/twitter-clone  /app

CMD /app/twitter-clone -configfile=properties.json

