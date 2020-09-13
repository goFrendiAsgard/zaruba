import amqplib from "amqplib";
import { Config } from "./config";
import { MainApp, createRouter } from "./core";
import { RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient } from "./transport";

import * as mainComponent from "./components/main/component";
import * as monitoring from "./components/monitoring/component";
import * as example from "./components/example/component";

async function main() {

    // app component definitions
    const logger = console;
    const config = new Config();
    logger.log("CONFIG:", JSON.stringify(config));
    const router = createRouter(logger);
    const rmq = await amqplib.connect(config.defaultRmqConnectionString);
    const rpcServer = new RmqRPCServer(logger, rmq, config.rmqEventMap);
    const rpcClient = new RmqRPCClient(logger, rmq, config.rmqEventMap);
    const subscriber = new RmqSubscriber(logger, rmq, config.rmqEventMap);
    const publisher = new RmqPublisher(logger, rmq, config.rmqEventMap);

    // app creation
    const app = new MainApp({
        logger,
        router,
        subscribers: [subscriber],
        rpcServers: [rpcServer],
        httpPort: config.httpPort,
    });

    // app setup
    app.setup([
        new mainComponent.Component(config, router),
        new monitoring.Component(config, app, router),
        new example.Component(config, app, router, publisher, subscriber, rpcServer, rpcClient),
    ]);

    // app execution
    app.run();

}

if (require.main === module) {
    main().catch((err) => {
        console.error(err);
    });
}