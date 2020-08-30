import { EnvelopedMessage } from "./envelopedMessage";

export function rpcCreateEnvelopedInput(inputs: any[]): EnvelopedMessage {
    const envelopedInput = new EnvelopedMessage();
    envelopedInput.setMessage({ "inputs": inputs });
    return envelopedInput;
}

export function rpcInputsToJSON(inputs: any[]): string {
    return rpcCreateEnvelopedInput(inputs).toJson();
}

export function rpcCreateEnvelopedError(envelopedInput: EnvelopedMessage, err: Error | string): EnvelopedMessage {
    const envelopedError = new EnvelopedMessage().setCorrelationId(envelopedInput.getCorrelationId());
    const errorMessage: string = typeof err === "string" ? err : err.message;
    envelopedError.setMessage({ "output": "", "error": errorMessage });
    envelopedError.setErrorMessage(errorMessage);
    return envelopedError;
}

export function rpcCreateEnvelopedOutput(envelopedInput: EnvelopedMessage, output: any): EnvelopedMessage {
    const envelopedOutput = new EnvelopedMessage().setCorrelationId(envelopedInput.getCorrelationId());
    envelopedOutput.setMessage({ "output": output, "error": "" });
    return envelopedOutput;
}

