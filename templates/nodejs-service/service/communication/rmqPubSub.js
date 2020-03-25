const { BasePubSub } = require("./baseClasses");
const { EnvelopedMessage } = require("./envelopedMessage");
const { createRmqConnectionAndChannel, closeRmqConnectionAndChannel, declareAndBindRmqQueueToExchange, declareRmqQueue, declareRmqFanoutExchange, rmqConsume, rmqPublish } = require("./rmqHelper");

class RmqPubSub extends BasePubSub {

    constructor(connectionString) {
        super();
        this.connectionString = connectionString;
        this.handlers = {};
        this.logger = console;
    }

    setLogger(logger) {
        this.logger = logger;
        return this;
    }

    registerHandler(eventName, handler) {
        this.handlers[eventName] = handler;
    }

    async start() {
        const self = this
        try {
            const { ch } = await createRmqConnectionAndChannel(this.connectionString);
            for (let key in this.handlers) {
                const eventName = key;
                await declareAndBindRmqQueueToExchange(ch, eventName)
                const handler = this.handlers[eventName];
                await rmqConsume(ch, eventName, async function (rmqMessage) {
                    try {
                        const jsonMessage = rmqMessage.content.toString()
                        const envelopedInput = new EnvelopedMessage(jsonMessage);
                        await handler(envelopedInput.message);
                    } catch (err) {
                        self.logger.error(err);
                    }
                })
            }
        } catch (err) {
            throw (err);
        }
    }

    async publish(eventName, message) {
        try {
            const { conn, ch } = await createRmqConnectionAndChannel(this.connectionString);
            await declareRmqFanoutExchange(ch, eventName);
            const envelopedMessage = new EnvelopedMessage().setCorrelationId().setMessage(message);
            await rmqPublish(ch, eventName, "", Buffer.from(envelopedMessage.toJson()));
            closeRmqConnectionAndChannel(conn, ch);
        } catch (err) {
            throw (err);
        }
    }

}

module.exports = {
    RmqPubSub
}