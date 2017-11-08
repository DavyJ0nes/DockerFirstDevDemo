FROM alpine:3.6
RUN adduser -D -u 10001 dockmaster

FROM scratch
LABEL Name=DavyJ0nes-DockerFirstDevDemo-Golang
LABEL Author=DavyJ0nes AuthorEmail=davy.jones@me.com

ADD level2_go-api /
COPY --from=0 /etc/passwd /etc/passwd
USER dockmaster

EXPOSE 3000
CMD ["./level2_go-api"]
