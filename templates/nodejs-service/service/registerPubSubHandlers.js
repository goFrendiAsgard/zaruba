const { createHelloHandler } = require("./pubsubhandlers/hello")

function registerPubSubHandlers(context, app, rpc, pubSub) {
    const config = context.getConfig();
    pubSub.registerHandler(config.defaultRmqEvent, createHelloHandler(context));
}

module.exports = { registerPubSubHandlers }
