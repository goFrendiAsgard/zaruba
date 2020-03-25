
function createLivenessHandler(context) {
    return function (req, res) {
        const httpCode = context.getLiveness() ? 200 : 500;
        const config = context.getConfig();
        res.status(httpCode).send({
            service_name: config.serviceName,
            is_alive: context.getLiveness(),
        });
    }
}

module.exports = { createLivenessHandler }