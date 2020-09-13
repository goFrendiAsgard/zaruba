import { RmqEventConfig } from "./interfaces";

export class RmqEventMap {
    constructor(private map: {[eventName: string]: RmqEventConfig}) {}

    public getExchangeName(eventName: string): string {
        if (eventName in this.map && "exchangeName" in this.map[eventName] && this.map[eventName].exchangeName !== "") {
            return this.map[eventName].exchangeName;
        }
        return eventName;
    }

    public getQueueName(eventName: string): string {
        if (eventName in this.map && "queueName" in this.map[eventName] && this.map[eventName].queueName !== "") {
            return this.map[eventName].queueName;
        }
        return eventName;
    }

    public getDeadLetterExchange(eventName: string): string {
        if (eventName in this.map && "deadLetterExchange" in this.map[eventName] && this.map[eventName].deadLetterExchange !== "") {
            return this.map[eventName].deadLetterExchange;
        }
        const exchangeName = this.getExchangeName(eventName);
        return `${exchangeName}.dlx`;
    }

    public getDeadLetterQueue(eventName: string): string {
        if (eventName in this.map && "deadLetterQueue" in this.map[eventName] && this.map[eventName].deadLetterQueue !== "") {
            return this.map[eventName].deadLetterQueue;
        }
        const queueName = this.getQueueName(eventName);
        return `${queueName}.dlx`;
    }

    public getTtl(eventName: string): number {
        if (eventName in this.map && "ttl" in this.map[eventName] && this.map[eventName].ttl > 0) {
            return this.map[eventName].ttl;
        }
        return 0;
    }

    public getRpcTimeout(eventName: string): number {
        if (eventName in this.map && "rpcTimeout" in this.map[eventName] && this.map[eventName].rpcTimeout as number > 0) {
            return this.map[eventName].rpcTimeout as number;
        }
        return 30000;
    }

    public getAutoAck(eventName: string): boolean {
        if (eventName in this.map && "autoAck" in this.map[eventName]) {
            return this.map[eventName].autoAck;
        }
        return false;
    }


}