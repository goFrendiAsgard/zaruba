module.exports = {
    createRmq,
    sendToQueue,
    publish,
    consume,
    bindQueue,
};

async function createRmq(serviceDesc, serviceDescRmq, amqplib) {
    const { logger } = serviceDesc;
    const { host, port, user, password, vhost } = serviceDescRmq;
    const connectionString = 'amqp://' + user + ':' + password + '@' + host + ':' + port + vhost;
    const connection = await amqplib.connect(connectionString);
    return { connection, logger };
}

async function sendToQueue(rmq, queue, message) {
    const { logger, connection } = rmq;
    try {
        channel = await connection.createChannel(connection);
        logger.log(`Sending to queue ${queue}: ${message}`)
        await channel.assertQueue(queue);
        await channel.sendToQueue(queue, Buffer.from(message));
        logger.log(`Message sent to queue ${queue}: ${message}`)
    } catch (error) {
        return logger.error(error);
    }
}


async function publish(rmq, exchange, message) {
    const { logger, connection } = rmq;
    try {
        channel = await connection.createChannel(connection);
        logger.log(`Sending to exchange ${exchange}: ${message}`)
        await channel.assertExchange(exchange, 'topic');
        await channel.publish(exchange, '', Buffer.from(message));
        logger.log(`Message sent to exchange ${exchange}: ${message}`)
    } catch (error) {
        return logger.error(error);
    }
}


async function consume(rmq, queue, onMessage) {
    const { logger, connection } = rmq;
    try {
        channel = await connection.createChannel(connection);
        await channel.assertQueue(queue);
        return channel.consume(queue, (messageBuffer) => {
            try {
                onMessage(messageBuffer.content.toString());
            } catch (error) {
                logger.error(error);
            }
        }, { noAck: true });
    } catch (error) {
        return logger.error(error);
    }
}

async function bindQueue(rmq, queue, exchange, pattern) {
    const { logger, connection } = rmq;
    pattern = pattern ? pattern : '#'
    try {
        channel = await connection.createChannel(connection);
        await channel.assertExchange(exchange, 'topic');
        await channel.assertQueue(queue);
        await channel.bindQueue(queue, exchange, pattern);
    } catch (error) {
        return logger.error(error);
    }
}