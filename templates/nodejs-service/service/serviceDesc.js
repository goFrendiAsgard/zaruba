module.exports = {
    serviceName: 'servicename', // Name of the service
    httpPort: process.env.SERVICENAME_HTTP_PORT || 3000, // HTTP Port, retrieved from envvar. By default, the value is `3000`
    logger: console,
    status: { // The status of the service. Typically related to kubernetes probe (https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)
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