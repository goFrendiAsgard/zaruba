import amqplib from "amqplib";
import { EventHandler, Subscriber } from "./interfaces";
import { EnvelopedMessage } from "./envelopedMessage";
import { rmqConsume, rmqDeclareQueueAndBindToDefaultExchange } from "./helpers";

export class RmqSubscriber implements Subscriber {
    connection: amqplib.Connection;
    logger: Console;
    handlers: { [functionName: string]: EventHandler };

    constructor(logger: Console, connection: amqplib.Connection) {
        this.connection = connection;
        this.logger = console;
        this.handlers = {};
    }

    registerHandler(eventName: string, handler: EventHandler): Subscriber {
        this.handlers[eventName] = handler;
        return this;
    }

    async subscribe(): Promise<void> {
        const self = this;
        return new Promise(async (_, reject) => {
            try {
                const ch = await self.connection.createChannel();
                self.connection.on("error", (err) => {
                    reject(err);
                });
                self.connection.on("close", (err) => {
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