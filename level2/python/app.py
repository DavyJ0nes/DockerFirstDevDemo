from flask import Flask, jsonify


def create_app():
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

    @app.errorhandler(404)
    def fourohfourHandler(err):
        response = jsonify({
            'Error': 'Page Not Found',
            'ErrorCode': 404
        })
        response.status_code = 404
        return response

    @app.route('/health', methods=['GET'])
    def healthHandler():
        return 'ok'

    return app


flask_app = create_app()
flask_app.run(debug=True, host='0.0.0.0')
