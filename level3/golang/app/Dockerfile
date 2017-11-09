# Multistage Build

## IMAGE 0
FROM alpine:3.6 AS alpine
RUN adduser -D -u 10001 dockmaster

## IMAGE 1
FROM scratch
LABEL Name=DockerFirstDevDemo-Level3-Golang
LABEL Author=DavyJ0nes AuthorEmail=davy.jones@me.com

COPY --from=alpine /etc/passwd /etc/passwd

ADD level3-goapp /
USER dockmaster

EXPOSE 3000
CMD ["./level3-goapp"]
