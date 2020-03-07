module.exports = {
    createApp,
    startApp,
};

function createApp(serviceDesc, express, bodyParser) {
    const { logger } = serviceDesc;

    // initiate app
    const app = express();
    app.use(bodyParser.urlencoded({ extended: false }));
    app.use(bodyParser.json());
    app.use((req, res, next) => {
        logger.log('HTTP REQUEST:', req.method, req.url, '\nQUERY:', req.query, '\nBODY:', req.body, '\nHEADER:', req.headers);
        try {
            next();
        } catch (error) {
            logger.error(error);
            return res.sendStatus(500);
        }
    });

    // add liveness
    app.all('/liveness', (_, res) => {
        const { serviceName, status } = serviceDesc;
        if (status.getLiveness()) {
            return res.status(200).send(serviceName + ' is alive');
        }
        return res.status(500).send(serviceName + ' is not alive');
    });


    // add readiness
    app.all('/readiness', (_, res) => {
        const { serviceName, status } = serviceDesc;
        if (status.getReadiness()) {
            return res.status(200).send(serviceName + ' is ready');
        }
        return res.status(500).send(serviceName + ' is not ready');
    });

    return app;
}


function startApp(serviceDesc, app) {
    const { httpPort, logger, serviceName } = serviceDesc;
    return app.listen(httpPort, () => {
        logger.log(`${serviceName} is running at port ${httpPort}`);
    });
}