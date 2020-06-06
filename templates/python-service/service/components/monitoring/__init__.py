from core import App
from config import Config
from typing import Callable
from flask import Flask, Response
import json


def create_setup(config: Config, app: App, router: Flask) -> Callable[[], None]:
    def setup():
        service_name = config.service_name

        @router.route("/liveness", methods=["GET"])
        def liveness():
            liveness = app.liveness()
            http_code = 200 if liveness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_alive": liveness})
            return Response(json_response, status=http_code, mimetype="application/json")

        @router.route("/readiness", methods=["GET"])
        def readiness():
            readiness = app.readiness()
            http_code = 200 if readiness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_ready": readiness})
            return Response(json_response, status=http_code, mimetype="application/json")

    return setup
