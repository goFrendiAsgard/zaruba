const { hello } = require("./rpchandlers/hello")

function registerRPCHandlers(context, app, rpc, pubSub) {
    rpc.registerHandler("helloRpc", hello);
}

module.exports = { registerRPCHandlers }
