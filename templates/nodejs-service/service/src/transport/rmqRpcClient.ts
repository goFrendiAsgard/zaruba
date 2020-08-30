import amqplib from "amqplib";
import { RPCClient } from "./interfaces";
import { rmqConsume, rmqDeclareQueue } from "./rmqHelper";
import { rmqRpcGenerateReplyQueueName, rmqRpcCall } from "./rmqRpcHelper";
import { EnvelopedMessage } from "./envelopedMessage";
import { RmqEventMap } from "./rmqEventMap";


export class RmqRPCClient implements RPCClient {

    constructor(private logger: Console, private connection: amqplib.Connection, private eventMap: RmqEventMap) {}

    public async call(functionName: string, ...inputs: any[]): Promise<any> {
        const self = this;
        const exchangeName = self.eventMap.getExchangeName(functionName);
        const queueName = self.eventMap.getQueueName(functionName);
        return new Promise(async function (resolve, reject) {
            try {
                const replyTo = rmqRpcGenerateReplyQueueName(queueName);
                const ch = await self.connection.createChannel();
                // consume
                await rmqDeclareQueue(ch, replyTo);
                rmqConsume(ch, replyTo, async (rmqMessageOrNull) => {
                    try {
                        const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                        const jsonEnvelopedOutput = rmqMessage.content.toString();
                        const envelopedOutput = new EnvelopedMessage(jsonEnvelopedOutput);
                        if (envelopedOutput.getErrorMessage()) {
                            self.logger.log(`[ERROR RmqRPCClient] Get Error Reply ${functionName}`, JSON.stringify(inputs), ":", envelopedOutput.getErrorMessage());
                            await self.deleteChannelAndQueue(ch, replyTo);
                            return reject(new Error(envelopedOutput.getErrorMessage()));
                        }
                        const message = envelopedOutput.getMessage();
                        await self.deleteChannelAndQueue(ch, replyTo);
                        self.logger.log(`[INFO RmqRPCClient] Get Reply ${functionName}`, JSON.stringify(inputs), ":", JSON.stringify(message.output));
                        resolve(message.output);
                    } catch (err) {
                        self.logger.error(`[ERROR RmqRPCClient] Error While Processing Reply ${functionName}`, JSON.stringify(inputs), ":", err);
                        await self.deleteChannelAndQueue(ch, replyTo);
                        reject(err);
                    }
                });
                // send message
                self.logger.log(`[INFO RmqRPCClient] Call ${functionName}`, JSON.stringify(inputs));
                await rmqRpcCall(ch, exchangeName, replyTo, inputs);
            } catch (err) {
                self.logger.error(err);
                reject(err);
            }
        });
    }

    async deleteChannelAndQueue(ch: amqplib.Channel, queueName: string) {
        try {
            await ch.deleteQueue(queueName, { ifUnused: false, ifEmpty: false });
            await ch.close();
        } catch (err) {
        }
    }


}