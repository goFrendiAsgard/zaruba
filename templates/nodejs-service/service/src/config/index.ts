export class Config {
    httpPort: number;
    serviceName: string;
    logger: Console;
    rmqConnectionString: string;
    localServiceAddress: string;

    constructor() {
        const servicePort = (process.env.SERVICENAME_HTTP_PORT || 3000) as number;
        const rmqConnectionString = getRmqConnectionString(
            process.env.RMQ_HOST || "localhost",
            (process.env.RMQ_PORT || 5672) as number,
            process.env.RMQ_USER || "root",
            process.env.RMQ_PASSWORD || "toor",
            process.env.RMQ_VHOST || "/",
        );
        this.httpPort = servicePort;
        this.serviceName = "servicename";
        this.logger = console;
        this.rmqConnectionString = rmqConnectionString;
        this.localServiceAddress = `http://localhost:${servicePort}`;
    }

}

export function getRmqConnectionString(host: string, port: number, user: string, password: string, vhost: string): string {
    return `amqp://${user}:${password}@${host}:${port}${vhost}`;
}