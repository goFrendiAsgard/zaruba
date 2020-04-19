export class Config {
    httpPort: number;
    serviceName: string;
    globalRmqConnectionString: string;
    localRmqConnectionString: string;

    constructor() {
        this.httpPort = (process.env.SERVICENAME_HTTP_PORT || 3000) as number;
        this.serviceName = "servicename";
        this.globalRmqConnectionString = process.env.GLOBAL_RMQ_CONNECTION_STRING || "amqp://localhost:5672/";
        this.localRmqConnectionString = process.env.LOCAL_RMQ_CONNECTION_STRING || "amqp://localhost:5672/";
    }

}