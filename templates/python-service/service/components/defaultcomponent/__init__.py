from flask import Flask
from core import App, Comp
from config import Config


class Component(Comp):

    def __init__(self, config: Config, router: Flask):
        self.config = config
        self.router = router

    def setup(self):
        @self.router.route("/")
        def index():
            return {"service_name": self.config.service_name}
