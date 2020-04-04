import express from "express";
import bootstrap from "./bootstrap";
import { Context } from "./context";
import { RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient, SimpleRPCServer, SimpleRPCClient, createHttpLogger } from "./transport";
import { Setting, Publishers, Subscribers, RPCClients, RPCServers } from "./components/setting";
import example from "./example";
import bodyParser from "body-parser";

async function boot() {

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
        new Publishers(
            new RmqPublisher(rmqConnectionString).setLogger(logger)
        ),
        new Subscribers(
            new RmqSubscriber(rmqConnectionString).setLogger(logger)
        ),
        new RPCServers(
            new RmqRPCServer(rmqConnectionString).setLogger(logger),
            new SimpleRPCServer(router).setLogger(logger)
        ),
        new RPCClients(
            new RmqRPCClient(rmqConnectionString).setLogger(logger),
            new SimpleRPCClient(ctx.config.localServiceAddress).setLogger(logger)
        ),
    );

    // TODO: remove this.
    example.setup(s);

    bootstrap.setup(s);
    bootstrap.run(s);

}
boot();