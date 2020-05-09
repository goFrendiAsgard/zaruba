import amqplib from "amqplib";
import { RPCServer, RPCHandler } from "./interfaces";
import { EnvelopedMessage } from "./envelopedMessage";
import { rmqCreateConnectionAndChannel, rmqDeclareQueueAndBindToDefaultExchange, rmqConsume, rmqRpcReplyOutput, rmqRpcReplyError } from "./helpers";

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

    serve(): Promise<void> {
        return new Promise(async (_, reject) => {
            try {
                const { conn, ch } = await rmqCreateConnectionAndChannel(this.connectionString);
                conn.on("error", (err) => {
                    reject(err);
                });
                conn.on("close", (err) => {
                    reject(err);
                })
                this.pServe(ch);
            } catch (err) {
                reject(err);
            }
        });
    }

    async pServe(ch: amqplib.Channel) {
        for (let key in this.handlers) {
            const functionName = key;
            const handler = this.handlers[functionName];
            await rmqDeclareQueueAndBindToDefaultExchange(ch, functionName);
            this.logger.log(`[INFO RmqRPCServer] Serve ${functionName}`);
            rmqConsume(ch, functionName, async (rmqMessageOrNull) => {
                try {
                    const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                    const { replyTo } = rmqMessage.properties;
                    const jsonEnvelopedInput = rmqMessage.content.toString();
                    const envelopedInput = new EnvelopedMessage(jsonEnvelopedInput);
                    try {
                        const inputs = envelopedInput.message.inputs;
                        const output = await handler(...inputs);
                        this.logger.log(`[INFO RmqRPCServer] Reply ${functionName}`, JSON.stringify(inputs), "output:", JSON.stringify(output));
                        await rmqRpcReplyOutput(ch, replyTo, envelopedInput, output);
                    } catch (err) {
                        this.logger.error(`[ERROR RmqRPCServer] Reply ${functionName}: `, err);
                        await rmqRpcReplyError(ch, replyTo, envelopedInput, err);
                    }
                } catch (err) {
                    this.logger.error(`[ERROR RmqRPCServer] Error Replying ${functionName}: `, err);
                }
            });
        }
    }

}