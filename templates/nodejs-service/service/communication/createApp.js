const express = require('express');
const bodyParser = require('body-parser');

function createApp(logger = null) {
    if (!logger) {
        logger = console;
    }
    const app = express();
    app.use(bodyParser.urlencoded({ extended: false }));
    app.use(bodyParser.json());
    app.use(async (req, res, next) => {
        const processName = `HTTP REQUEST: ${req.method} ${req.url}`;
        logger.time(processName);
        try {
            await next();
        } catch (err) {
            logger.error(error);
            res.sendStatus(500);
        }
        logger.timeEnd(processName);
    });
    return app;
}

module.exports = { createApp };