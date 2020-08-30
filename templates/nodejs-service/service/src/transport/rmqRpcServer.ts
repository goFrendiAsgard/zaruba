import amqplib from "amqplib";
import { RPCServer, RPCHandler } from "./interfaces";
import { EnvelopedMessage } from "./envelopedMessage";
import { rmqDeclareAndBindQueue, rmqConsume } from "./rmqHelper";
import { rmqRpcReplyOutput, rmqRpcReplyError } from "./rmqRpcHelper";
import { RmqEventMap } from "./rmqEventMap";

export class RmqRPCServer implements RPCServer {
    private handlers: { [functionName: string]: RPCHandler };

    constructor(private logger: Console, private connection: amqplib.Connection, private eventMap: RmqEventMap) {
        this.handlers = {};
    }

    public registerHandler(functionName: string, handler: RPCHandler): RPCServer {
        this.handlers[functionName] = handler;
        return this;
    }

    public serve(): Promise<void> {
        const self = this;
        return new Promise(async (_, reject) => {
            try {
                const ch = await self.connection.createChannel();
                self.connection.on("error", (err) => {
                    reject(err);
                });
                self.connection.on("close", (err) => {
                    reject(err);
                })
                this.pServe(ch);
            } catch (err) {
                reject(err);
            }
        });
    }

    private async pServe(ch: amqplib.Channel) {
        for (let key in this.handlers) {
            const functionName = key;
            const exchangeName = this.eventMap.getExchangeName(functionName);
            const queueName = this.eventMap.getExchangeName(functionName);
            const handler = this.handlers[functionName];
            await rmqDeclareAndBindQueue(ch, exchangeName, queueName);
            this.logger.log(`[INFO RmqRPCServer] Serve ${functionName}`);
            rmqConsume(ch, queueName, async (rmqMessageOrNull) => {
                try {
                    const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                    const { replyTo } = rmqMessage.properties;
                    const jsonEnvelopedInput = rmqMessage.content.toString();
                    const envelopedInput = new EnvelopedMessage(jsonEnvelopedInput);
                    try {
                        const inputs = envelopedInput.getMessage().inputs;
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