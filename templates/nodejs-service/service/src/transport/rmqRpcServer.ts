import amqplib from "amqplib";
import { RPCServer, RPCHandler } from "./interfaces";
import { EnvelopedMessage } from "./envelopedMessage";
import { rpcCreateEnvelopedErrorMessage, rpcCreateEnvelopedOutputMessage, rmqCreateConnectionAndChannel, rmqDeclareQueueAndBindToDefaultExchange, rmqConsume, rmqRpcReply, rmqRpcReplyError } from "./helpers";

export class RmqRPCServer implements RPCServer {
    logger: Console;
    connectionString: string;
    handlers: { [functionName: string]: RPCHandler };

    constructor(connectionString: string) {
        this.logger = console;
        this.connectionString = connectionString;
        this.handlers = {};
    }

    registerHandler(functionName: string, handler: RPCHandler): RPCServer {
        this.handlers[functionName] = handler;
        return this;
    }

    setLogger(logger: Console): RPCServer {
        this.logger = logger;
        return this;
    }

    async serve() {
        const { conn, ch } = await rmqCreateConnectionAndChannel(this.connectionString);
        for (let key in this.handlers) {
            const functionName = key;
            const handler = this.handlers[functionName];
            rmqDeclareQueueAndBindToDefaultExchange(ch, functionName);
            rmqConsume(ch, functionName, async (rmqMessageOrNull) => {
                const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                const { replyTo } = rmqMessage.properties;
                const jsonMessage = rmqMessage.content.toString();
                const envelopedInput = new EnvelopedMessage(jsonMessage);
                try {
                    const output = await handler(envelopedInput.message);
                    await rmqRpcReply(ch, replyTo, envelopedInput, output);
                } catch (err) {
                    await rmqRpcReplyError(ch, replyTo, envelopedInput, err);
                }
            });
        }
    }

}