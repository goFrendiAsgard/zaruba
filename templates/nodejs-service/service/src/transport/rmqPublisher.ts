import amqplib from "amqplib";
import { Message, Publisher } from "./interfaces";
import { rmqPublish, rmqDeclareFanoutExchange } from "./rmqHelper";
import { EnvelopedMessage } from "./envelopedMessage";
import { RmqEventMap } from "./rmqEventMap";

export class RmqPublisher implements Publisher {

    constructor(private logger: Console, private connection: amqplib.Connection, private eventMap: RmqEventMap) {}

    public async publish(eventName: string, msg: Message): Promise<void> {
        this.logger.log("[INFO RmqPublisher] Publish", eventName, JSON.stringify(msg));
        const ch = await this.connection.createChannel();
        const exchangeName = this.eventMap.getExchangeName(eventName);
        await rmqDeclareFanoutExchange(ch, exchangeName);
        const envelopedMessage = new EnvelopedMessage().setCorrelationId().setMessage(msg);
        await rmqPublish(ch, exchangeName, "", Buffer.from(envelopedMessage.toJson()));
        await ch.close();
    };

}