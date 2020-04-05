import { RPCClient } from "./interfaces";
import { rpcCreateEnvelopedInputMessage } from "./helpers";

import axios from "axios";
import { EnvelopedMessage } from "./envelopedMessage";


export class SimpleRPCClient implements RPCClient {
    serverAddress: string
    logger: Console;

    constructor(serverAddress: string) {
        this.serverAddress = serverAddress;
        this.logger = console;
    }

    setLogger(logger: Console): RPCClient {
        this.logger = logger;
        return this;
    }

    async call(functionName: string, ...inputs: any[]): Promise<any> {
        const envelopedInput = rpcCreateEnvelopedInputMessage(inputs);
        const remoteAddr = `${this.serverAddress}/api/${functionName}`
        const res = await axios.post(remoteAddr, envelopedInput);
        const envelopedOutput = new EnvelopedMessage(res.data);
        return envelopedOutput.message.output;
    }

}