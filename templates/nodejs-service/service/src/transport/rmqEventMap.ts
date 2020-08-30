import { RmqEventConfig } from "./interfaces";

export class RmqEventMap {
    constructor(private map: {[eventName: string]: RmqEventConfig}) {}

    public getExchangeName(eventName: string) {
        return this.map[eventName] ? this.map[eventName].exchangeName : eventName;
    }

    public getQueueName(eventName: string) {
        return this.map[eventName] ? this.map[eventName].queueName : eventName;
    }

}