import { v4 } from "uuid";
import amqplib from "amqplib";
import { EnvelopedMessage } from "./envelopedMessage";

const uuid = v4;

export function rpcCreateEnvelopedInputMessage(inputs: any[]): EnvelopedMessage {
    const envelopedMessage = new EnvelopedMessage();
    envelopedMessage.message = { "inputs": inputs };
    return envelopedMessage;
}

export function rpcInputsToJSON(inputs: any[]): string {
    return rpcCreateEnvelopedInputMessage(inputs).toJson();
}

export function rpcCreateEnvelopedErrorMessage(envelopedInput: EnvelopedMessage, err: Error | string): EnvelopedMessage {
    const envelopedError = new EnvelopedMessage().setCorrelationId(envelopedInput.correlationId);
    const errorMessage: string = typeof err === "string" ? err : err.message;
    envelopedError.message = { "output": "", "error": errorMessage };
    envelopedError.errorMessage = errorMessage;
    return envelopedError;
}

export function rpcCreateEnvelopedOutputMessage(envelopedInput: EnvelopedMessage, output: any): EnvelopedMessage {
    const envelopedOutput = new EnvelopedMessage().setCorrelationId(envelopedInput.correlationId);
    envelopedOutput.message = { "output": output, "error": "" };
    return envelopedOutput;
}


export function rmqRpcGenerateReplyQueueName(functionName: string): string {
    const randomId = uuid().split("-").join("");
    return `${functionName}.reply.${randomId}`
}

export async function rmqRpcCall(ch: amqplib.Channel, functionName: string, replyTo: string, inputs: any[]): Promise<boolean> {
    const envelopedInput = rpcCreateEnvelopedInputMessage(inputs);
    const jsonMessage = envelopedInput.toJson();
    return rmqPublish(ch, functionName, "", new Buffer(jsonMessage), {
        contentType: "text/json",
        correlationId: envelopedInput.correlationId,
        replyTo,
    });
}

export async function rmqRpcReply(ch: amqplib.Channel, replyTo: string, envelopedInput: EnvelopedMessage, output: any): Promise<boolean> {
    const jsonMessage = rpcCreateEnvelopedOutputMessage(envelopedInput, output).toJson();
    return rmqPublish(ch, "", replyTo, new Buffer(jsonMessage), {
        contentType: "text/json",
        correlationId: envelopedInput.correlationId,
    });
}

export async function rmqRpcReplyError(ch: amqplib.Channel, replyTo: string, envelopedInput: EnvelopedMessage, err: Error | string): Promise<boolean> {
    const jsonMessage = rpcCreateEnvelopedErrorMessage(envelopedInput, err).toJson();
    return rmqPublish(ch, "", replyTo, new Buffer(jsonMessage), {
        contentType: "text/json",
        correlationId: envelopedInput.correlationId,
    });
}

export async function rmqCreateConnectionAndChannel(connectionString: string): Promise<{ conn: amqplib.Connection, ch: amqplib.Channel }> {
    const conn = await amqplib.connect(connectionString);
    const ch = await conn.createChannel();
    return { conn, ch };
}

export async function rmqDeclareQueueAndBindToDefaultExchange(ch: amqplib.Channel, queueName: string): Promise<amqplib.Replies.AssertQueue> {
    const exchangeName = queueName
    await rmqDeclareFanoutExchange(ch, exchangeName)
    const q = await rmqDeclareQueue(ch, queueName)
    await ch.bindQueue(queueName, exchangeName, "");
    return q
}

export async function rmqDeclareQueue(ch: amqplib.Channel, queueName: string): Promise<amqplib.Replies.AssertQueue> {
    return await ch.assertQueue(queueName, { durable: false });
}

export async function rmqDeclareFanoutExchange(ch: amqplib.Channel, exchangeName: string): Promise<amqplib.Replies.AssertExchange> {
    return await ch.assertExchange(exchangeName, "fanout", {});
}

export async function rmqConsume(ch: amqplib.Channel, queueName: string, handler: (msq: amqplib.ConsumeMessage | null) => any): Promise<amqplib.Replies.Consume> {
    return await ch.consume(
        queueName,
        handler,
        {
            noAck: true,
        },
    );
}

export async function rmqPublish(ch: amqplib.Channel, exchangeName: string, routingKey: string, data: Buffer, options: amqplib.Options.Publish = {}): Promise<boolean> {
    return ch.publish(
        exchangeName, routingKey,
        data,
        options
    );
}

export async function rmqCloseConnectionAndChannel(conn: amqplib.Connection, ch: amqplib.Channel): Promise<void> {
    try {
        await ch.close();
    } catch (err) { }
    try {
        await conn.close();
    } catch (err) { }
}