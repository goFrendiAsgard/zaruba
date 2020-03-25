const { createLivenessHandler } = require("./httphandlers/liveness");
const { createReadinessHandler } = require("./httphandlers/readiness");
const { hello } = require("./httphandlers/hello");
const { createHelloAllHandler } = require("./httphandlers/helloAll");
const { createHelloRPCHandler } = require("./httphandlers/helloRpc");
const { createHelloPubSubHandler } = require("./httphandlers/helloPubSub");

function registerHTTPHandlers(context, app, rpc, pubSub) {
    // liveness and readiness handlers
    app.get("/liveness", createLivenessHandler(context));
    app.get("/readiness", createReadinessHandler(context));

    // Default route
    app.all("/", (req, res) => { res.send("servicename") });

    // EXAMPLE: hello
    app.get("/hello", hello);
    app.post("/hello", hello);
    app.get("/hello/:name", hello);
    app.post("/hello/:name", hello);

    // EXAMPLE: hello-rpc
    helloRPCHandler = createHelloRPCHandler(rpc);
    app.get("/hello-rpc", helloRPCHandler)
    app.post("/hello-rpc", helloRPCHandler)
    app.get("/hello-rpc/:name", helloRPCHandler)
    app.post("/hello-rpc:name", helloRPCHandler)

    // EXAMPLE: hello-pubsub
    helloPubSubHandler = createHelloPubSubHandler(context, pubSub);
    app.get("/hello-pubsub", helloPubSubHandler)
    app.post("/hello-pubsub", helloPubSubHandler)
    app.get("/hello-pubsub/:name", helloPubSubHandler)
    app.post("/hello-pubsub:name", helloPubSubHandler)

    // EXAMPLE: hello-all
    app.get("/hello-all", createHelloAllHandler(context));

}

module.exports = { registerHTTPHandlers }