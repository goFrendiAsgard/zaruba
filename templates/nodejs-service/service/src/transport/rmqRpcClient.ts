import amqplib from "amqplib";
import { RPCClient } from "./interfaces";
import { rmqRpcGenerateReplyQueueName, rmqConsume, rmqCreateConnectionAndChannel, rmqDeclareQueue, rmqCloseConnectionAndChannel, rmqRpcCall } from "./helpers";
import { EnvelopedMessage } from "./envelopedMessage";


export class RmqRPCClient implements RPCClient {
    connectionString: string
    logger: Console;

    constructor(connectionString: string) {
        this.connectionString = connectionString;
        this.logger = console;
    }

    setLogger(logger: Console): RPCClient {
        this.logger = logger;
        return this;
    }

    async call(functionName: string, ...inputs: any[]): Promise<any> {
        const self = this;
        return new Promise(async function (resolve, reject) {
            try {
                const replyTo = rmqRpcGenerateReplyQueueName(functionName);
                const { conn, ch } = await rmqCreateConnectionAndChannel(self.connectionString);
                // consume
                await rmqDeclareQueue(ch, replyTo);
                let replyAccepted = false;
                rmqConsume(ch, replyTo, async (rmqMessageOrNull) => {
                    if (replyAccepted) {
                        return false;
                    }
                    replyAccepted = true;
                    try {
                        const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                        const jsonEnvelopedOutput = rmqMessage.content.toString();
                        const envelopedOutput = new EnvelopedMessage(jsonEnvelopedOutput);
                        if (envelopedOutput.errorMessage) {
                            self.logger.log(`[ERROR RmqRPCClient] Get Error Reply ${functionName}`, JSON.stringify(inputs), ":", envelopedOutput.errorMessage);
                            await self.deleteQueueAndCloseConnection(conn, ch, replyTo);
                            return reject(new Error(envelopedOutput.errorMessage));
                        }
                        const message = envelopedOutput.message;
                        await self.deleteQueueAndCloseConnection(conn, ch, replyTo);
                        self.logger.log(`[INFO RmqRPCClient] Get Reply ${functionName}`, JSON.stringify(inputs), ":", JSON.stringify(message.output));
                        resolve(message.output);
                    } catch (err) {
                        self.logger.error(`[ERROR RmqRPCClient] Error While Processing Reply ${functionName}`, JSON.stringify(inputs), ":", err);
                        await self.deleteQueueAndCloseConnection(conn, ch, replyTo);
                        reject(err);
                    }
                });
                // send message
                self.logger.log(`[INFO RmqRPCClient] Call ${functionName}`, JSON.stringify(inputs));
                await rmqRpcCall(ch, functionName, replyTo, inputs);
            } catch (err) {
                self.logger.error(err);
                reject(err);
            }
        });
    }

    async deleteQueueAndCloseConnection(conn: amqplib.Connection, ch: amqplib.Channel, queueName: string) {
        try {
            await ch.deleteQueue(queueName, { ifUnused: false, ifEmpty: false });
            await rmqCloseConnectionAndChannel(conn, ch);
        } catch (err) {
        }
    }


}