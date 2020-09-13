import { RmqEventMap } from "../transport";

export class Config {
    public httpPort: number;
    public serviceName: string;
    public defaultRmqConnectionString: string;
    public rmqEventMap: RmqEventMap;

    constructor() {
        this.httpPort = (process.env.SERVICENAME_HTTP_PORT || 3000) as number;
        this.serviceName = "servicename";
        this.defaultRmqConnectionString = process.env.DEFAULT_RMQ_CONNECTION_STRING || "amqp://localhost:5672/";
        this.rmqEventMap = new RmqEventMap({
            helloRPC: {
                exchangeName: "servicename.exchange.helloRPC",
                queueName: "servicename.queue.helloRPC",
                deadLetterExchange: "servicename.exchange.helloRPC.dlx",
                deadLetterQueue: "servicename.queue.helloRPC.dlx",
                ttl: 10000,
                autoAck: true,
                rpcTimeout: 20000,
            },
            hello: {
                exchangeName: "servicename.exchange.helloEvent",
                queueName: "servicename.queue.helloEvent",
                deadLetterExchange: "servicename.exchange.helloEvent.dlx",
                deadLetterQueue: "servicname.queue.helloEvent.dlx",
                ttl: 60000,
                autoAck: false,
            }
        });
    }

}