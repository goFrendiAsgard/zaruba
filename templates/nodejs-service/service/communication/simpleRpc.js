const request = require("request");
const { EnvelopedMessage } = require("./envelopedMessage");
const { BaseRPC } = require("./baseClasses")

class SimpleRPC extends BaseRPC {

    constructor(engine, serviceMap) {
        super();
        this.engine = engine;
        this.serviceMap = serviceMap;
        this.handlers = {};
        this.logger = console;
    }

    setLogger(logger) {
        this.logger = logger;
        return this;
    }

    registerHandler(functionName, handler) {
        this.handlers[functionName] = handler;
    }

    serve() {
        for (let key in this.handlers) {
            const functionName = key;
            const handler = this.handlers[functionName];
            this.engine.post(`/api/${functionName}`, function (req, res) {
                let correlationId = "";
                try {
                    const envelopedInput = new EnvelopedMessage(req.body);
                    correlationId = envelopedInput.correlationId;
                    const output = handler(envelopedInput.message);
                    const envelopedOutput = new EnvelopedMessage().setCorrelationId(correlationId).setMessage(output);
                    res.send(envelopedOutput.toJson())
                } catch (err) {
                    const errorMessage = typeof err == "string" ? err : err.message;
                    res.json({ correlation_id: correlationId, error: errorMessage })
                }
            });
        }
    }

    async call(serviceName, functionName, input) {
        const url = this.serviceMap[serviceName]
        const envelopedInput = new EnvelopedMessage().setCorrelationId().setMessage(input);
        return new Promise(function (resolve, reject) {
            request.post({
                headers: { "content-type": "application/json" },
                url: `${url}/api/${functionName}`,
                body: envelopedInput.toJson(),
            }, (err, res, body) => {
                if (err) {
                    return reject(err);
                }
                try {
                    const envelopedOutput = new EnvelopedMessage(body);
                    resolve(envelopedOutput.message);
                } catch (err) {
                    reject(err);
                }
            });
        })
    }

}

module.exports = {
    SimpleRPC
};
