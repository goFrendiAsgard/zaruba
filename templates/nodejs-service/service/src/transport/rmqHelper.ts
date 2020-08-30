import amqplib from "amqplib";

export async function rmqDeclareAndBindQueue(ch: amqplib.Channel, exchangeName: string, queueName: string): Promise<amqplib.Replies.AssertQueue> {
    await rmqDeclareFanoutExchange(ch, exchangeName);
    const q = await rmqDeclareQueue(ch, queueName);
    await ch.bindQueue(queueName, exchangeName, "");
    return q
}

export async function rmqDeclareQueue(ch: amqplib.Channel, queueName: string): Promise<amqplib.Replies.AssertQueue> {
    return await ch.assertQueue(queueName, { durable: false });
}

export async function rmqDeclareFanoutExchange(ch: amqplib.Channel, exchangeName: string): Promise<amqplib.Replies.AssertExchange> {
    return await ch.assertExchange(exchangeName, "fanout", {});
}

export async function rmqConsume(ch: amqplib.Channel, queueName: string, handler: (msq: amqplib.ConsumeMessage | null) => any): Promise<amqplib.Replies.Consume> {
    return await ch.consume(
        queueName,
        handler,
        {
            noAck: true,
        },
    );
}

export async function rmqPublish(ch: amqplib.Channel, exchangeName: string, routingKey: string, data: Buffer, options: amqplib.Options.Publish = {}): Promise<boolean> {
    return ch.publish(
        exchangeName, routingKey,
        data,
        options
    );
}