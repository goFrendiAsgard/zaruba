import { Message, Publisher, RPCClient } from "../../transport";
import { greet, greetEveryone } from "./services";

export class Controller {

    private names: string[] = [];
    constructor(private publisher: Publisher, private rpcClient: RPCClient) {
    }

    private getName(req: any): string {
        return req.params.name || req.query.name || req.body.name || "";
    }

    public async handleHTTPHello(req: any, res: any) {
        const name = this.getName(req);
        res.send(greet(name));
    }

    public async handleHTTPHelloAll(req: any, res: any) {
        res.send(greetEveryone(this.names));
    }

    public async handleHTTPHelloRPC(req: any, res: any) {
        const name = this.getName(req);
        try {
            const greeting = await this.rpcClient.call("helloRPC", name);
            res.send(greeting);
        } catch (err) {
            res.status(500).send(err);
        }
    }

    public async handleHTTPHelloPub(req: any, res: any) {
        const name = this.getName(req);
        try {
            await this.publisher.publish("hello", { name });
            res.send("Message sent");
        } catch (err) {
            res.status(500).send(err);
        }
    }

    public async handleRPCHello(...inputs: any[]): Promise<any> {
        if (inputs.length === 0) {
            throw new Error("Message accepted but input is invalid");
        }
        const name = inputs[0] as string
        return greet(name);
    }

    public async handleEventHello(msg: Message) {
        const { name } = msg;
        this.names.push(name);
    }

}