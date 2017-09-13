from flask import Flask, jsonify
app = Flask(__name__)


@app.route('/v1/data', methods=['GET'])
def dataHandler():

    response = jsonify({
        'Name': 'Python API',
        'RandomString': 'alsdjf;aoiwu2403982098sdf',
        'Version': 'PythonAPI - v0.0.1'
    })
    response.status_code = 200
    return response


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
