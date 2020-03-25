const { v4: uuid } = require("uuid");

class EnvelopedMessage {

    constructor(jsonMessage = null) {
        this.correlationId = "";
        this.message = {};
        this.errorMessage = "";
        if (jsonMessage) {
            const obj = typeof jsonMessage == "string" ? JSON.parse(jsonMessage) : jsonMessage;
            this.correlationId = obj.correlation_id;
            this.message = obj.message;
            this.errorMessage = obj.error;
        }
    }

    setCorrelationId(correlationId = null) {
        if (correlationId) {
            this.correlationId = correlationId;
            return this;
        }
        this.correlationId = uuid();
        return this;
    }

    setMessage(message) {
        this.message = message;
        return this;
    }

    toJson() {
        return JSON.stringify({
            correlation_id: this.correlationId,
            message: this.message,
            error: this.errorMessage,
        });
    }

}

module.exports = {
    EnvelopedMessage
};