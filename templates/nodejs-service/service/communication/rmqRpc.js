const amqplib = require('amqplib');
const { v4: uuid } = require("uuid");
const { createRmqConnectionAndChannel, closeRmqConnectionAndChannel, declareAndBindRmqQueueToExchange, declareRmqQueue, declareRmqFanoutExchange, rmqConsume, rmqPublish } = require("./rmqHelper");

const { EnvelopedMessage } = require("./envelopedMessage");
const { BaseRPC } = require("./baseClasses")

class RmqRPC extends BaseRPC {

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

    registerHandler(functionName, handler) {
        this.handlers[functionName] = handler;
    }

    async serve() {
        const self = this;
        try {
            const { ch } = await createRmqConnectionAndChannel(this.connectionString);
            for (let key in this.handlers) {
                const eventName = key;
                await declareAndBindRmqQueueToExchange(ch, eventName)
                const handler = this.handlers[eventName];
                await rmqConsume(ch, eventName, async function (rmqMessage) {
                    try {
                        const jsonMessage = rmqMessage.content.toString()
                        const { replyTo, correlationId } = rmqMessage.properties;
                        const envelopedInput = new EnvelopedMessage(jsonMessage);
                        const output = await handler(envelopedInput.message);
                        const envelopedOutput = new EnvelopedMessage().setCorrelationId(correlationId).setMessage(output);
                        rmqPublish(ch, "", replyTo, Buffer.from(envelopedOutput.toJson()));
                    } catch (err) {
                        self.logger.error(err);
                        const errMessage = typeof err == "string" ? err : err.message;
                        rmqPublish(ch, "", replyTo, new EnvelopedMessage().setMessage(errMessage));
                    }
                })
            }
        } catch (err) {
            throw (err);
        }
    }

    async call(serviceName, functionName, input) {
        const self = this;
        return new Promise(async function (resolve, reject) {
            try {
                const replyTo = await self.generateReplyQueueName(functionName);
                const { conn, ch } = await createRmqConnectionAndChannel(self.connectionString);
                // consume reply
                await declareRmqQueue(ch, replyTo);
                let replyAccepted = false;
                rmqConsume(ch, replyTo, async (rmqMessage) => {
                    if (replyAccepted) {
                        return false;
                    }
                    try {
                        replyAccepted = true;
                        const jsonMessage = rmqMessage.content.toString()
                        const envelopedOutput = new EnvelopedMessage(jsonMessage)
                        if (envelopedOutput.errorMessage) {
                            return reject(new Error(envelopedOutput.errorMessage));
                        }
                        const message = envelopedOutput.message;
                        await self.deleteQueue(conn, ch, replyTo);
                        resolve(message);
                    } catch (err) {
                        self.logger.error(err);
                        await self.deleteQueue(conn, ch, replyTo);
                        reject(error);
                    }
                });
                // send message
                await declareRmqFanoutExchange(ch, functionName);
                const envelopedInput = new EnvelopedMessage().setCorrelationId().setMessage(input);
                await rmqPublish(ch, functionName, "", Buffer.from(envelopedInput.toJson()), {
                    contentType: "text/json",
                    correlationId: envelopedInput.correlationId,
                    replyTo: replyTo,
                });
            } catch (err) {
                self.logger.error(err);
                reject(err);
            }
        });
    }

    async deleteQueue(conn, ch, queueName) {
        try {
            await ch.deleteQueue(queueName, { ifUnused: false, ifEmpty: false, noWait: true });
            await closeRmqConnectionAndChannel(conn, ch);
        } catch (err) {
        }
    }

    generateReplyQueueName(functionName) {
        const rawRandomId = uuid();
        const randomId = rawRandomId.split("-").join("");
        return `${functionName}.reply.${randomId}`;
    }


}

module.exports = { RmqRPC };

