FROM alpine:3.6
LABEL NAME=docker-first-python
LABEL AUTHOR=DavyJ0nes

# Add non root user to run app as
RUN adduser -D -s /bin/sh dockmaster

# Install curl and Python3
RUN apk update && \
  apk add curl && \
  apk add python3 && \
  ln -s /usr/bin/python3 /usr/bin/python && \
  ln -s /usr/bin/pip3 /usr/bin/pip

WORKDIR /home/dockmaster

# Adding pip installing requirements.txt so not not break build cache
ADD ./requirements.txt /home/dockmaster
RUN pip install -r requirements.txt

# all commands after this will be carried out by the dockmaster user
USER dockmaster
ADD . /home/dockmaster

# Healthcheck to ensure app code starts up correctly
HEALTHCHECK --interval=10s --timeout=2s \
  CMD curl -f http://localhost:5000/health || exit 1

EXPOSE 5000
CMD ["python", "app.py"]

