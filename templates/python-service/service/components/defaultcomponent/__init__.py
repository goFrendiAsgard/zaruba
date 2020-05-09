from core import App, SetupComponent
from config import Config


def create_setup(app: App, config: Config) -> SetupComponent:

    def index():
        return {"service_name": config.service_name}

    def setup():
        r = app.router()
        r.add_url_rule("/", "index", index)

    return setup
