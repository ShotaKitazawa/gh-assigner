## Build stage ##
FROM golang:1.13.3
WORKDIR /go/src/github.com/ShotaKitazawa/gh-assigner/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

## Run stage ##
FROM alpine:3.10.3
# Install github.com/jwilder/dockerize
RUN apk add --no-cache openssl
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
# Copy app binary from above image (golang:1.13.3)
COPY --from=0 /go/src/github.com/tamac-io/sre-sampleapp/app .
# Run
ENTRYPOINT ["./app"]
