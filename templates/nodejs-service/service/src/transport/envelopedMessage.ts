import { v4 } from "uuid";
const uuid = v4;

export class EnvelopedMessage {
    correlationId: string;
    message: { [key: string]: any };
    errorMessage: string;

    constructor(jsonMessage: string = "") {
        this.correlationId = "";
        this.message = {};
        this.errorMessage = "";
        if (jsonMessage !== "") {
            const obj = typeof jsonMessage == "string" ? JSON.parse(jsonMessage) : jsonMessage;
            this.correlationId = obj.correlation_id;
            this.message = obj.message;
            this.errorMessage = obj.error;
        }
    }

    setCorrelationId(correlationId: string = ""): EnvelopedMessage {
        if (correlationId === "") {
            this.correlationId = correlationId;
            return this;
        }
        this.correlationId = uuid();
        return this;
    }

    setMessage(message: { [key: string]: any }): EnvelopedMessage {
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