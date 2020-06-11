from core import App, Comp
from config import Config
from typing import Callable
from flask import Flask, Response
import json


class Component(Comp):

    def __init__(self, config: Config, app: App, router: Flask):
        self.config = config
        self.app = app
        self.router = router

    def setup(self):
        service_name = self.config.service_name

        @self.router.route("/liveness", methods=["GET"])
        def liveness():
            liveness = self.app.liveness()
            http_code = 200 if liveness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_alive": liveness})
            return Response(json_response, status=http_code, mimetype="application/json")

        @self.router.route("/readiness", methods=["GET"])
        def readiness():
            readiness = self.app.readiness()
            http_code = 200 if readiness else 500
            json_response = json.dumps(
                {"service_name": service_name, "is_ready": readiness})
            return Response(json_response, status=http_code, mimetype="application/json")
