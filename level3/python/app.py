from flask import Flask
from redis import Redis
#  import socket

app = Flask(__name__)
redis = Redis(host='redis', port=6379)


@app.route('/')
def hello():
    count = redis.incr('hits')
    #  hostname = socket.gethostname()
    return 'Hello from Docker! This page has been seen {} times.\n'.format(
        count
    )


if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True)
