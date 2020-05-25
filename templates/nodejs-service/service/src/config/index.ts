export class Config {
    httpPort: number;
    serviceName: string;
    defaultRmqConnectionString: string;

    constructor() {
        this.httpPort = (process.env.SERVICENAME_HTTP_PORT || 3000) as number;
        this.serviceName = "servicename";
        this.defaultRmqConnectionString = process.env.DEFAULT_RMQ_CONNECTION_STRING || "amqp://localhost:5672/";
    }

}