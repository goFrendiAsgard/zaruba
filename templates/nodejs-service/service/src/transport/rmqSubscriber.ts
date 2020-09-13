import amqplib from "amqplib";
import { EventHandler, Subscriber } from "./interfaces";
import { EnvelopedMessage } from "./envelopedMessage";
import { rmqConsume, rmqDeclareAndBindQueue, rmqDeclareQueue } from "./rmqHelper";
import { RmqEventMap } from "./rmqEventMap";

export class RmqSubscriber implements Subscriber {
    private handlers: { [functionName: string]: EventHandler };

    constructor(private logger: Console, private connection: amqplib.Connection, private eventMap: RmqEventMap) {
        this.handlers = {};
    }

    public registerHandler(eventName: string, handler: EventHandler): Subscriber {
        this.handlers[eventName] = handler;
        return this;
    }

    public async subscribe(): Promise<void> {
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

    private async pSubscribe(ch: amqplib.Channel) {
        const self = this;
        for (let key in this.handlers) {
            const eventName = key;
            // declare dlx
            const args: amqplib.Options.AssertQueue = {};
            if (this.eventMap.getTtl(eventName) > 0) {
                const deadLetterExchange = this.eventMap.getDeadLetterExchange(eventName);
                const deadLetterQueue = this.eventMap.getDeadLetterQueue(eventName);
                await rmqDeclareAndBindQueue(ch, deadLetterExchange, deadLetterQueue, {durable: true});
                args.deadLetterExchange = deadLetterExchange;
                args.messageTtl = this.eventMap.getTtl(eventName);
            }
            // declare queue
            const exchangeName = self.eventMap.getExchangeName(eventName);
            const queueName = self.eventMap.getQueueName(eventName);
            const handler = this.handlers[eventName];
            await rmqDeclareAndBindQueue(ch, exchangeName, queueName, args);
            this.logger.log(`[INFO RmqSubscriber] Subscribe ${eventName}`);
            const autoAck = this.eventMap.getAutoAck(eventName);
            // handle message
            rmqConsume(ch, queueName, args, autoAck, async (rmqMessageOrNull) => {
                const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                try {
                    const jsonEnvelopedInput = rmqMessage.content.toString();
                    const envelopedInput = new EnvelopedMessage(jsonEnvelopedInput);
                    this.logger.log(`[INFO RmqSubscriber] Get Event ${eventName}: `, JSON.stringify(envelopedInput.getMessage()));
                    try {
                        await handler(envelopedInput.getMessage());
                        ch.ack(rmqMessage);
                    } catch(err) {
                        ch.nack(rmqMessage);
                    }
                } catch (err) {
                    this.logger.log(`[ERROR RmqSubscriber] Get Event ${eventName}: `, err);
                    self.logger.error(err);
                    ch.ack(rmqMessage)
                }
            })
        }
    }

}