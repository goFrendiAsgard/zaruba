from config import Config
from core import MainApp
import components.defaultcomponent as defaultComponent
import components.monitoring as monitoring
import components.example as example


def main():

    # create config and app
    config = Config()
    print("CONFIG: {}".format(config))
    app = MainApp(
        http_port=config.http_port,
        global_rmq_connection_string=config.global_rmq_connection_string,
        local_rmq_connection_string=config.local_rmq_connection_string
    )

    # setup components
    app.setup([
        defaultComponent.create_setup(app, config),
        monitoring.create_setup(app, config),
        example.Component(app, config).setup
    ])

    # run
    app.run()


if __name__ == "__main__":
    main()
