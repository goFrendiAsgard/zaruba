from flask import Flask
from core import App, SetupComponent
from config import Config


def create_setup(config: Config, router: Flask) -> SetupComponent:

    def setup():
        @router.route("/")
        def index():
            return {"service_name": config.service_name}

    return setup
