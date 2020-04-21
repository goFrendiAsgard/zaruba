import { Message, Publisher } from "./interfaces";
import { rmqCreateConnectionAndChannel, rmqDeclareQueueAndBindToDefaultExchange, rmqPublish, rmqCloseConnectionAndChannel } from "./helpers";
import { EnvelopedMessage } from "./envelopedMessage";

export class RmqPublisher implements Publisher {
    connectionString: string
    logger: Console;

    constructor(connectionString: string) {
        this.connectionString = connectionString;
        this.logger = console;
    }

    setLogger(logger: Console): Publisher {
        this.logger = logger;
        return this;
    }

    async publish(eventName: string, msg: Message) {
        this.logger.log("[INFO RmqPublisher] Publish", eventName, JSON.stringify(msg));
        const { conn, ch } = await rmqCreateConnectionAndChannel(this.connectionString);
        await rmqDeclareQueueAndBindToDefaultExchange(ch, eventName);
        const envelopedMessage = new EnvelopedMessage().setCorrelationId().setMessage(msg);
        await rmqPublish(ch, eventName, "", Buffer.from(envelopedMessage.toJson()));
        rmqCloseConnectionAndChannel(conn, ch);
    };

}