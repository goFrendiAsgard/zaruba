module.exports = {
    serviceName: 'servicename',

    // Configurations
    httpPort: process.env.SERVICENAME_HTTP_PORT || 3000,
    rmq: {
        host: process.env.RMQ_HOST || "localhost",
        port: process.env.RMQ_PORT || 5672,
        user: process.env.RMQ_USER || "root",
        password: process.env.RMQ_PASSWORD || "toor",
        vhost: process.env.RMQ_VHOST || "/",
    },
    rmqEvent: process.env.SERVICENAME_EVENT || "servicename",

    // Components
    logger: console,

    // Status of the service. Typically related to kubernetes probe (https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)
    status: {
        _isAlive: true,
        _isReady: true,
        setLiveness: function (liveness) {
            this._isAlive = liveness;
        },
        setReadiness: function (readiness) {
            this._isReady = readiness;
        },
        getLiveness: function () {
            return this._isAlive;
        },
        getReadiness: function () {
            return this._isReady;
        }
    }

}