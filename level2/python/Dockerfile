FROM alpine:3.6
LABEL NAME=docker-first-python
LABEL AUTHOR=DavyJ0nes

# Install curl and Python3
RUN apk update && \
  apk add python3 && \
  ln -s /usr/bin/python3 /usr/bin/python && \
  ln -s /usr/bin/pip3 /usr/bin/pip && \
  mkdir -p /src/app

WORKDIR /src/app

# Adding pip installing requirements.txt so not not mess with build cache
ADD ./requirements.txt /src/app
RUN pip install -r requirements.txt

ADD . /src/app

EXPOSE 5000
CMD ["python", "app.py"]

