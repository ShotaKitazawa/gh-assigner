FROM alpine:latest

EXPOSE 8080

COPY gh-assigner /gh-assigner

ENTRYPOINT ["/gh-assigner"]
