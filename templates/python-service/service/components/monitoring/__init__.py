from core import App
from config import Config
from typing import Callable
from flask import Response
import json


def create_setup(app: App, config: Config) -> Callable[[], None]:
    def setup():
        service_name = config.service_name
        r = app.router()

        def liveness():
            liveness = app.liveness()
            http_code = 200 if liveness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_alive": liveness})
            return Response(json_response, status=http_code, mimetype="application/json")

        def readiness():
            readiness = app.readiness()
            http_code = 200 if readiness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_ready": readiness})
            return Response(json_response, status=http_code, mimetype="application/json")

        r.add_url_rule("/liveness", "liveness", liveness, methods=["GET"])
        r.add_url_rule("/readiness", "readiness", readiness, methods=["GET"])

    return setup
