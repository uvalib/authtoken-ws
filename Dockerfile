FROM alpine:3.3

RUN apk add --update bash && rm -rf /var/cache/apk/*

EXPOSE 8080
CMD ./authtoken-ws

COPY bin/authtoken-ws.linux authtoken-ws
