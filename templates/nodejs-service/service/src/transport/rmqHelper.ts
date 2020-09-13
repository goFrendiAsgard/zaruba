import amqplib from "amqplib";

export async function rmqDeclareAndBindQueue(ch: amqplib.Channel, exchangeName: string, queueName: string, options: amqplib.Options.AssertQueue): Promise<amqplib.Replies.AssertQueue> {
    await rmqDeclareFanoutExchange(ch, exchangeName);
    const q = await rmqDeclareQueue(ch, queueName, options);
    await ch.bindQueue(queueName, exchangeName, "");
    return q
}

export async function rmqDeclareQueue(ch: amqplib.Channel, queueName: string, options: amqplib.Options.AssertQueue): Promise<amqplib.Replies.AssertQueue> {
    if (!("durable" in options)) {
        options.durable = true;
    }
    return await ch.assertQueue(queueName, options);
}

export async function rmqDeclareFanoutExchange(ch: amqplib.Channel, exchangeName: string): Promise<amqplib.Replies.AssertExchange> {
    return await ch.assertExchange(exchangeName, "fanout", { durable: true });
}

export async function rmqConsume(ch: amqplib.Channel, queueName: string, options: amqplib.Options.AssertQueue, autoAck: boolean, handler: (msq: amqplib.ConsumeMessage | null) => any): Promise<amqplib.Replies.Consume> {
    await rmqDeclareQueue(ch, queueName, options);
    return await ch.consume(
        queueName,
        handler,
        {
            noAck: autoAck,
        },
    );
}

export async function rmqPublish(ch: amqplib.Channel, exchangeName: string, routingKey: string, data: Buffer, options: amqplib.Options.Publish = {}): Promise<boolean> {
    if (exchangeName != "") {
        await rmqDeclareFanoutExchange(ch, exchangeName);
    }
    return ch.publish(
        exchangeName, routingKey,
        data,
        options
    );
}