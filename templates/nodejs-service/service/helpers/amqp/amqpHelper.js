module.exports = {
    createConnection,
    sendToQueue,
    publish,
    consume,
    bindQueue,
};

async function createConnection(serviceDescRmq, amqplib) {
    const { host, port, user, password, vhost } = serviceDescRmq;
    const connectionString = 'amqp://' + user + ':' + password + '@' + host + ':' + port + vhost;
    return await amqplib.connect(connectionString);
}

async function sendToQueue(serviceDesc, connection, queue, message) {
    const { logger } = serviceDesc;
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


async function publish(serviceDesc, connection, exchange, message) {
    const { logger } = serviceDesc;
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


async function consume(serviceDesc, connection, queue, onMessage) {
    const { logger } = serviceDesc;
    try {
        channel = await connection.createChannel(connection);
        await channel.assertQueue(queue);
        return channel.consume(queue, (messageBuffer) => {
            try {
                onMessage(messageBuffer.content.toString());
            } catch (error) {
                logger.error(error);
            }
        });
    } catch (error) {
        return logger.error(error);
    }
}

async function bindQueue(serviceDesc, connection, queue, exchange, pattern) {
    const { logger } = serviceDesc;
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