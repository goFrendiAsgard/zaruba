class BaseRPC {

    call(serviceName, functionName, input) {
        throw new Error("Not implemented");
    }

    serve() {
        throw new Error("Not implemented");
    }

    registerHandler(functionName, handler) {
        throw new Error("Not impemented");
    }

}

class BasePubSub {

    publish(eventName, message) {
        throw new Error("Not implemented");
    }

    start() {
        throw new Error("Not implemented");
    }

    registerHandler(eventName, handler) {
        throw new error("Not implemented");
    }

}

module.exports = {
    BaseRPC, BasePubSub
}