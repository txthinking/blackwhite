FROM alpine:latest

ADD pac /pac
ADD white.list /white.list
ADD black.list /black.list
ADD customize.map /customize.map

ENTRYPOINT ["/pac"]


