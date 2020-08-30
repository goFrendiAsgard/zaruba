import { v4 } from "uuid";
import { Message } from "./interfaces";
const uuid = v4;

export class EnvelopedMessage {
    private correlationId: string;
    private message: Message;
    private errorMessage: string;

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

    public setCorrelationId(correlationId?: string): EnvelopedMessage {
        this.correlationId == correlationId !== undefined ? correlationId : uuid();
        return this;
    }

    public setMessage(message: Message): EnvelopedMessage {
        this.message = message;
        return this;
    }

    public setErrorMessage(errorMessage: string): EnvelopedMessage {
        this.errorMessage = errorMessage;
        return this;
    }

    public toJson(): string {
        return JSON.stringify({
            correlation_id: this.correlationId,
            message: this.message,
            error: this.errorMessage,
        });
    }

    public getCorrelationId(): string {
        return this.correlationId;
    }

    public getMessage(): Message {
        return this.message;
    }

    public getErrorMessage(): string {
        return this.errorMessage;
    }


}