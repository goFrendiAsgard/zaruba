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
        return new Promise(async (resolve, reject) => {
            try {
                const replyTo = rmqRpcGenerateReplyQueueName(queueName);
                const ch = await self.connection.createChannel();
                // consume
                await rmqDeclareQueue(ch, replyTo, {durable: true, autoDelete: true});
                rmqConsume(ch, replyTo, {durable: true, autoDelete: true}, true, async (rmqMessageOrNull) => {
                    try {
                        const rmqMessage = rmqMessageOrNull as amqplib.ConsumeMessage;
                        const jsonEnvelopedOutput = rmqMessage.content.toString();
                        const envelopedOutput = new EnvelopedMessage(jsonEnvelopedOutput);
                        if (envelopedOutput.getErrorMessage()) {
                            self.logger.log(`[ERROR RmqRPCClient] Get Error Reply ${functionName}`, JSON.stringify(inputs), ":", envelopedOutput.getErrorMessage());
                            return reject(new Error(envelopedOutput.getErrorMessage()));
                        }
                        const message = envelopedOutput.getMessage();
                        self.logger.log(`[INFO RmqRPCClient] Get Reply ${functionName}`, JSON.stringify(inputs), ":", JSON.stringify(message.output));
                        resolve(message.output);
                    } catch (err) {
                        self.logger.error(`[ERROR RmqRPCClient] Error While Processing Reply ${functionName}`, JSON.stringify(inputs), ":", err);
                        reject(err);
                    }
                });
                // timeout
                const timeout = self.eventMap.getRpcTimeout(functionName);
                setTimeout(() => {
                    const err = new Error(`Timeout ${timeout}`);
                    self.logger.error(`[ERROR RmqRPCClient] Get timeout ${functionName} ${inputs}: ${err}`)
                    reject(err);
                }, timeout);
                // send message
                self.logger.log(`[INFO RmqRPCClient] Call ${functionName}`, JSON.stringify(inputs));
                await rmqRpcCall(ch, exchangeName, replyTo, inputs);
            } catch (err) {
                self.logger.error(err);
                reject(err);
            }
        });
    }

}