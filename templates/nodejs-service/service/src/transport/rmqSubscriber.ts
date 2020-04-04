import amqplib from "amqplib";
import { EventHandler, Subscriber } from "./interfaces";
import { EnvelopedMessage } from "./envelopedMessage";
import { rpcCreateEnvelopedErrorMessage, rpcCreateEnvelopedOutputMessage, rmqCreateConnectionAndChannel, rmqDeclareQueueAndBindToDefaultExchange, rmqConsume, rmqRpcReply, rmqRpcReplyError, rmqDeclareFanoutExchange } from "./helpers";

export class RmqSubscriber implements Subscriber {
    logger: Console;
    connectionString: string;
    handlers: { [functionName: string]: EventHandler };

    constructor(connectionString: string) {
        this.logger = console;
        this.connectionString = connectionString;
        this.handlers = {};
    }

    registerHandler(functionName: string, handler: EventHandler): Subscriber {
        this.handlers[functionName] = handler;
        return this;
    }

    setLogger(logger: Console): Subscriber {
        this.logger = logger;
        return this;
    }

    async subscribe() {
        const self = this;
        const { ch } = await rmqCreateConnectionAndChannel(this.connectionString);
        for (let key in this.handlers) {
            const eventName = key;
            const handler = this.handlers[eventName];
            rmqDeclareFanoutExchange(ch, eventName);
            rmqConsume(ch, eventName, async (rmqMessageOrNull) => {
                try {
                    const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                    const jsonMessage = rmqMessage.content.toString();
                    const envelopedInput = new EnvelopedMessage(jsonMessage);
                    await handler(envelopedInput.message);
                } catch (err) {
                    self.logger.error(err);
                }
            })
        }
    }

}