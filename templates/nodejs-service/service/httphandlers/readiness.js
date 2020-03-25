
function createReadinessHandler(context) {
    return function (req, res) {
        const httpCode = context.getReadiness() ? 200 : 500;
        const config = context.getConfig();
        res.status(httpCode).send({
            service_name: config.serviceName,
            is_ready: context.getReadiness(),
        });
    }
}

module.exports = { createReadinessHandler }