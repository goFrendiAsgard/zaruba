import { v4 } from "uuid";
import { Message } from "./interfaces";
const uuid = v4;

export class EnvelopedMessage {
    correlationId: string;
    message: Message;
    errorMessage: string;

    constructor(jsonEnvelopedMessage?: string) {
        this.correlationId = "";
        this.message = {};
        this.errorMessage = "";
        if (jsonEnvelopedMessage !== undefined) {
            const obj = JSON.parse(jsonEnvelopedMessage);
            this.correlationId = obj.correlation_id;
            this.message = obj.message;
            this.errorMessage = obj.error;
        }
    }

    setCorrelationId(correlationId?: string): EnvelopedMessage {
        this.correlationId == correlationId !== undefined ? correlationId : uuid();
        return this;
    }

    setMessage(message: Message): EnvelopedMessage {
        this.message = message;
        return this;
    }

    toJson(): string {
        return JSON.stringify({
            correlation_id: this.correlationId,
            message: this.message,
            error: this.errorMessage,
        });
    }

}