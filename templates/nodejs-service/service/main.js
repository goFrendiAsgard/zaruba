// imports
const { Context } = require("./context");
const { RmqPubSub, RmqRPC, SimpleRPC, createApp } = require("./communication");
const { registerHTTPHandlers } = require("./registerHttpHandlers");
const { registerRPCHandlers } = require("./registerRpcHandlers");
const { registerPubSubHandlers } = require("./registerPubSubHandlers");

async function main() {
    const context = new Context();
    const config = context.getConfig();
    const rmqConnectionString = config.defaultRmq.createConnectionString();
    const { logger } = config;

    const app = createApp(logger);
    const pubSub = new RmqPubSub(rmqConnectionString).setLogger(logger);
    const rpc = new RmqRPC(rmqConnectionString).setLogger(logger);
    // const rpc = new SimpleRPC(app, config.serviceUrlMap).setLogger(logger);

    registerHTTPHandlers(context, app, rpc, pubSub);
    registerRPCHandlers(context, app, rpc, pubSub);
    registerPubSubHandlers(context, app, rpc, pubSub);

    pubSub.start();
    rpc.serve();
    app.listen(config.httpPort, () => {
        logger.log(`Run at port ${config.httpPort}`);
    });
}

if (require.main == module) {
    main().catch(error => console.log(error));
}