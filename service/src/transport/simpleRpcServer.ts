import { RPCServer, RPCHandler } from "./interfaces";
import { Express } from "express";
import { EnvelopedMessage } from "./envelopedMessage";
import { rpcCreateEnvelopedErrorMessage, rpcCreateEnvelopedOutputMessage } from "./helpers";

export class SimpleRPCServer implements RPCServer {
    logger: Console;
    engine: Express;
    handlers: { [functionName: string]: RPCHandler };

    constructor(engine: Express) {
        this.logger = console;
        this.engine = engine;
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

    serve(): void {
        for (let key in this.handlers) {
            const functionName = key;
            const handler = this.handlers[functionName];
            this.engine.post(`/api/${functionName}`, function (req, res) {
                const envelopedInput = new EnvelopedMessage(req.body);
                try {
                    const inputs = envelopedInput.message.inputs as any[];
                    const output = handler(...inputs);
                    res.json(rpcCreateEnvelopedOutputMessage(envelopedInput, output));
                } catch (err) {
                    res.json(rpcCreateEnvelopedErrorMessage(envelopedInput, err))
                }
            });
        }
    }

}