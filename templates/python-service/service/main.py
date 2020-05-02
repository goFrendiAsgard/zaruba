from config import Config
from core import MainApp
import components.defaultcomponent as defaultcomponent


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
        defaultcomponent.create_setup(app, config),
    ])

    # run
    app.run()


if __name__ == "__main__":
    main()
