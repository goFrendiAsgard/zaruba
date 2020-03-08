// imports
const amqplib = require('amqplib');
const express = require('express');
const bodyParser = require('body-parser');

const serviceDesc = require('./serviceDesc');
const { createApp, startApp } = require('./helpers/express/expressHelper');
const { createRmq, sendToQueue, consume } = require('./helpers/amqp/amqpHelper');
const { logger, rmqEvent } = serviceDesc;

async function main() {
    try {
        // Variables
        const app = createApp(serviceDesc, express, bodyParser);
        const rmq = await createRmq(serviceDesc, serviceDesc.rmq, amqplib);

        // Http routes
        app.all('/', (req, res) => {
            logger.log('QUERY', req.query)
            logger.log('BODY', req.body)
            return res.status(200).send('Hello World !!!');
        });

        // Start Listening to rmqEvent
        await consume(rmq, rmqEvent, (message) => {
            logger.log("GET MESSAGE:", message);
        });
        // Send Message to rmqEvent
        await sendToQueue(rmq, rmqEvent, "A message");

        // Start HTTP Server
        startApp(serviceDesc, app);
    } catch (error) {
        logger.error(error);
    }
}

if (require.main == module) {
    main();
}