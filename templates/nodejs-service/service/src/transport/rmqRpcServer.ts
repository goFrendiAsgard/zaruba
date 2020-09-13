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

    private async declareDlxAndBuildArgs(ch: amqplib.Channel, functionName: string): Promise<amqplib.Options.AssertQueue> {
        const args: amqplib.Options.AssertQueue = {};
        if (this.eventMap.getTtl(functionName) > 0) {
            const deadLetterExchange = this.eventMap.getDeadLetterExchange(functionName);
            const deadLetterQueue = this.eventMap.getDeadLetterQueue(functionName);
            await rmqDeclareAndBindQueue(ch, deadLetterExchange, deadLetterQueue, { durable: true });
            args.deadLetterExchange = deadLetterExchange;
            args.messageTtl = this.eventMap.getTtl(functionName);
        }
        return args;
    }

    private async pServe(ch: amqplib.Channel) {
        for (let key in this.handlers) {
            const functionName = key;
            // declare dlx
            const args = await this.declareDlxAndBuildArgs(ch, functionName);
            // declare queue
            const exchangeName = this.eventMap.getExchangeName(functionName);
            const queueName = this.eventMap.getExchangeName(functionName);
            const handler = this.handlers[functionName];
            await rmqDeclareAndBindQueue(ch, exchangeName, queueName, args);
            this.logger.log(`[INFO RmqRPCServer] Serve ${functionName}`);
            // consume
            const autoAck = this.eventMap.getAutoAck(functionName);
            rmqConsume(ch, queueName, args, autoAck, async (rmqMessageOrNull) => {
                try {
                    const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                    if (!autoAck) {
                        ch.ack(rmqMessage);
                    }
                    const { replyTo } = rmqMessage.properties;
                    const jsonEnvelopedInput = rmqMessage.content.toString();
                    const envelopedInput = new EnvelopedMessage(jsonEnvelopedInput);
                    const inputs = envelopedInput.getMessage().inputs;
                    try {
                        const output = await handler(...inputs);
                        this.logger.log(`[INFO RmqRPCServer] Reply ${functionName} by Sending to Queue ${replyTo}: `, JSON.stringify(output));
                        try {
                            await rmqRpcReplyOutput(ch, replyTo, envelopedInput, output);
                        } catch (err) {
                            this.logger.log(`[INFO RmqRPCServer] Failed to Reply ${functionName} by Sending to Queue ${replyTo}: `, err);
                        }
                    } catch (err) {
                        this.logger.error(`[ERROR RmqRPCServer] Reply Error ${functionName} by Sending to queue ${replyTo}: `, err);
                        await rmqRpcReplyError(ch, replyTo, envelopedInput, err);
                    }
                } catch (err) {
                    this.logger.error(`[ERROR RmqRPCServer] Error Replying ${functionName}: `, err);
                }
            });
        }
    }

}