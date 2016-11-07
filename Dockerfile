FROM alpine:latest

ADD pac /pac
ADD white.list /white.list
ADD black.list /black.list
ADD customize.map /customize.map
RUN apk --no-cache add ca-certificates && update-ca-certificates

ENTRYPOINT ["/pac"]


