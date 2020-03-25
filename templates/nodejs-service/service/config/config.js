class Config {

    constructor() {
        const httpPort = process.env.SERVICENAME_HTTP_PORT || 3000;
        this.httpPort = httpPort;
        this.serviceName = "servicename";
        this.logger = console;
        this.defaultRmqEvent = process.env.SERVICENAME_EVENT || "servicename";
        this.defaultRmq = new RmqConfig(
            process.env.RMQ_HOST || "localhost",
            process.env.RMQ_PORT || 5672,
            process.env.RMQ_USER || "root",
            process.env.RMQ_PASSWORD || "toor",
            process.env.RMQ_VHOST || "/"
        );
        this.serviceUrlMap = {
            "servicename": process.env.SERVICENAME_URL || `http://localhost:${httpPort}`,
        };
    }

}

class RmqConfig {

    constructor(host, port, user, password, vhost) {
        this.host = host;
        this.port = port;
        this.user = user;
        this.password = password;
        this.vhost = vhost;
    }

    createConnectionString() {
        return `amqp://${this.user}:${this.password}@${this.host}:${this.port}${this.vhost}`
    }

}

module.exports = {
    Config, RmqConfig
};
