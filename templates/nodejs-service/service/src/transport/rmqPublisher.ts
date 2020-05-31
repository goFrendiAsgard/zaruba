import amqplib from "amqplib";
import { Message, Publisher } from "./interfaces";
import { rmqDeclareQueueAndBindToDefaultExchange, rmqPublish } from "./helpers";
import { EnvelopedMessage } from "./envelopedMessage";

export class RmqPublisher implements Publisher {
    connection: amqplib.Connection;
    logger: Console;

    constructor(logger: Console, connection: amqplib.Connection) {
        this.connection = connection;
        this.logger = logger;
    }

    async publish(eventName: string, msg: Message): Promise<void> {
        this.logger.log("[INFO RmqPublisher] Publish", eventName, JSON.stringify(msg));
        const ch = await this.connection.createChannel();
        await rmqDeclareQueueAndBindToDefaultExchange(ch, eventName);
        const envelopedMessage = new EnvelopedMessage().setCorrelationId().setMessage(msg);
        await rmqPublish(ch, eventName, "", Buffer.from(envelopedMessage.toJson()));
        await ch.close();
    };

}