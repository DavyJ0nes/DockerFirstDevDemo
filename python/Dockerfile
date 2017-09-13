FROM alpine:3.6
LABEL Name=docker-first-python Version=0.0.1

MAINTAINER DavyJ0nes <davy.jones@me.com>

RUN apk update && \
  apk add python3 && \
  ln -s /usr/bin/python3 /usr/bin/python && \
  ln -s /usr/bin/pip3 /usr/bin/pip

RUN mkdir -p /srv/app
ADD . /srv/app/
WORKDIR /srv/app/

RUN pip install -r requirements.txt

EXPOSE 5000
ENTRYPOINT ["python"]
CMD ["app.py"]
