import amqplib from "amqplib";
import { v4 } from "uuid";
import { EnvelopedMessage } from "./envelopedMessage";
import { rpcCreateEnvelopedInput, rpcCreateEnvelopedOutput, rpcCreateEnvelopedError } from "./rpcHelper";
import { rmqPublish } from "./rmqHelper";
const uuid = v4;

export function rmqRpcGenerateReplyQueueName(queueName: string): string {
    const randomId = uuid().split("-").join("");
    return `${queueName}.reply.${randomId}`
}

export async function rmqRpcCall(ch: amqplib.Channel, exchangeName: string, replyTo: string, inputs: any[]): Promise<boolean> {
    const envelopedInput = rpcCreateEnvelopedInput(inputs);
    const jsonInput = envelopedInput.toJson();
    return rmqPublish(ch, exchangeName, "", new Buffer(jsonInput), {
        contentType: "text/json",
        correlationId: envelopedInput.getCorrelationId(),
        replyTo,
    });
}

export async function rmqRpcReplyOutput(ch: amqplib.Channel, replyTo: string, envelopedInput: EnvelopedMessage, output: any): Promise<boolean> {
    const jsonOutput = rpcCreateEnvelopedOutput(envelopedInput, output).toJson();
    return rmqPublish(ch, "", replyTo, new Buffer(jsonOutput), {
        contentType: "text/json",
        correlationId: envelopedInput.getCorrelationId(),
    });
}

export async function rmqRpcReplyError(ch: amqplib.Channel, replyTo: string, envelopedInput: EnvelopedMessage, err: Error | string): Promise<boolean> {
    const jsonError = rpcCreateEnvelopedError(envelopedInput, err).toJson();
    return rmqPublish(ch, "", replyTo, new Buffer(jsonError), {
        contentType: "text/json",
        correlationId: envelopedInput.getCorrelationId(),
    });
}