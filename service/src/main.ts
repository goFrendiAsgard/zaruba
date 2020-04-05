import express from "express";
import bootstrap from "./bootstrap";
import { Context } from "./context";
import { RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient, SimpleRPCServer, SimpleRPCClient, createHttpLogger } from "./transport";
import { Setting } from "./components/setting";
import bodyParser from "body-parser";

function main() {

    const ctx = new Context();
    const logger = ctx.config.logger;
    const rmqConnectionString = ctx.config.rmqConnectionString;

    const router = express();
    router.use(bodyParser.urlencoded({ extended: false }));
    router.use(bodyParser.json());
    router.use(createHttpLogger(logger));

    const s: Setting = new Setting(
        ctx,
        router,
        // publishers
        {
            main: new RmqPublisher(rmqConnectionString).setLogger(logger)
        },
        // subscribers
        {
            main: new RmqSubscriber(rmqConnectionString).setLogger(logger)
        },
        // rpc servers
        {
            main: new RmqRPCServer(rmqConnectionString).setLogger(logger),
            secondary: new SimpleRPCServer(router).setLogger(logger)
        },
        // rpc clients
        {
            mainLoopBack: new RmqRPCClient(rmqConnectionString).setLogger(logger),
            secondaryLoopBack: new SimpleRPCClient(ctx.config.localServiceAddress).setLogger(logger)
        },
    );


    bootstrap.setup(s);
    bootstrap.run(s);

}

main();