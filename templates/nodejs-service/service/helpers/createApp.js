
module.exports = function (serviceDesc, express, bodyParser) {
    const { httpPort, logger, serviceName } = serviceDesc;

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

    // custom function to start http server
    app.startService = () => {
        app.listen(httpPort, () => {
            logger.log(`${serviceName} is running at port ${httpPort}`);
        });
    }

    return app;
}