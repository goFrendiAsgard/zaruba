const amqplib = require('amqplib');

async function createRmqConnectionAndChannel(connectionString) {
    const conn = await amqplib.connect(connectionString);
    const ch = await conn.createChannel();
    return { conn, ch };
}

async function declareAndBindRmqQueueToExchange(ch, queueName) {
    const exchangeName = queueName
    await declareRmqFanoutExchange(ch, exchangeName)
    q = await declareRmqQueue(ch, queueName)
    await ch.bindQueue(queueName, exchangeName, "");
    return q
}

async function declareRmqQueue(ch, queueName) {
    return await ch.assertQueue(queueName, { durable: false });
}
async function declareRmqFanoutExchange(ch, exchangeName) {
    return await ch.assertExchange(exchangeName, "fanout", {});
}

async function rmqConsume(ch, queueName, handler) {
    return await ch.consume(
        queueName,
        handler,
        {
            noAck: true,
        },
    );
}

async function rmqPublish(ch, exchangeName, routingKey, data, options = {}) {
    return await ch.publish(
        exchangeName, routingKey,
        data,
        options
    );
}

async function closeRmqConnectionAndChannel(conn, ch) {
    try {
        await ch.close();
    } catch (err) { }
    try {
        await conn.close();
    } catch (err) { }
}

module.exports = {
    createRmqConnectionAndChannel, closeRmqConnectionAndChannel, declareAndBindRmqQueueToExchange, declareRmqQueue, declareRmqFanoutExchange, rmqConsume, rmqPublish
};
