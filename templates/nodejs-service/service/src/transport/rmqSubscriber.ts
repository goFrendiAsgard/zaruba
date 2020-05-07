import amqplib from "amqplib";
import { EventHandler, Subscriber } from "./interfaces";
import { EnvelopedMessage } from "./envelopedMessage";
import { rmqCreateConnectionAndChannel, rmqConsume, rmqDeclareQueueAndBindToDefaultExchange } from "./helpers";

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
        return new Promise(async (_, reject) => {
            try {
                const { conn, ch } = await rmqCreateConnectionAndChannel(this.connectionString);
                conn.on("error", (err) => {
                    reject(err);
                });
                conn.on("close", (err) => {
                    reject(err);
                });
                this.pSubscribe(ch);
            } catch (err) {
                reject(err);
            }
        });
    }

    async pSubscribe(ch: amqplib.Channel) {
        const self = this;
        for (let key in this.handlers) {
            const eventName = key;
            const handler = this.handlers[eventName];
            await rmqDeclareQueueAndBindToDefaultExchange(ch, eventName);
            this.logger.log(`[INFO RmqSubscriber] Subscribe ${eventName}`);
            rmqConsume(ch, eventName, async (rmqMessageOrNull) => {
                try {
                    const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                    const jsonEnvelopedInput = rmqMessage.content.toString();
                    const envelopedInput = new EnvelopedMessage(jsonEnvelopedInput);
                    this.logger.log(`[INFO RmqSubscriber] Get Event ${eventName}: `, JSON.stringify(envelopedInput.message));
                    await handler(envelopedInput.message);
                } catch (err) {
                    this.logger.log(`[ERROR RmqSubscriber] Get Event ${eventName}: `, err);
                    self.logger.error(err);
                }
            })
        }
    }

}