// imports
const express = require('express');
const bodyParser = require('body-parser');
const serviceDesc = require('./serviceDesc');
const createApp = require('./helpers/createApp');

// global variables
const { logger } = serviceDesc;
const app = createApp(serviceDesc, express, bodyParser);

app.all('/', (req, res) => {
    logger.log('QUERY', req.query)
    logger.log('BODY', req.body)
    return res.status(200).send('Hello World !!!');
});

app.get('/kill', (req, res) => {
    serviceDesc.status.setLiveness(false);
    return res.send('kill');
});

app.get('/revive', (req, res) => {
    serviceDesc.status.setLiveness(true);
    return res.send('revive');
});

if (require.main == module) {
    // Serve HTTP Server
    app.startService();
}