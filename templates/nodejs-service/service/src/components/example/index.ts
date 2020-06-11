import { Express } from "express";
import { Message, Publisher, Subscriber, RPCServer, RPCClient } from "../../transport";
import { Comp } from "../../core";
import { Config } from "../../config";
import { getName } from "./helpers";
import { greet, greetEveryone } from "./services";

export class Component implements Comp {

    private names: string[];

    constructor(private config: Config, private router: Express, private publisher: Publisher, private subscriber: Subscriber, private rpcServer: RPCServer, private rpcClient: RPCClient) {
        this.names = [];
    }

    setup() {

        // Use the same HTTP Handler for multiple URLS
        this.router.get("/hello", this.handleHTTPHello.bind(this));
        this.router.get("/hello/:name", this.handleHTTPHello.bind(this));
        this.router.post("/hello", this.handleHTTPHello.bind(this));

        // Use HTTP Handler that take state from component
        this.router.get("/hello-all", this.handleHTTPHelloAll.bind(this));

        // Trigger RPC Call
        this.router.get("/hello-rpc", this.handleHTTPHelloRPC.bind(this));
        this.router.get("/hello-rpc/:name", this.handleHTTPHelloRPC.bind(this));
        this.router.post("/hello-rpc", this.handleHTTPHelloRPC.bind(this));

        // Trigger Publisher
        this.router.get("/hello-pub", this.handleHTTPHelloPub.bind(this));
        this.router.get("/hello-pub/:name", this.handleHTTPHelloPub.bind(this));
        this.router.post("/hello-pub", this.handleHTTPHelloPub.bind(this));

        // Serve RPC
        this.rpcServer.registerHandler("servicename.helloRPC", this.handleRPCHello.bind(this));

        // Event
        this.subscriber.registerHandler("servicename.helloEvent", this.handleEventHello.bind(this));

    }

    async handleHTTPHello(req: any, res: any) {
        const name = getName(req);
        res.send(greet(name));
    }

    async handleHTTPHelloAll(req: any, res: any) {
        res.send(greetEveryone(this.names));
    }

    async handleHTTPHelloRPC(req: any, res: any) {
        const name = getName(req);
        try {
            const greeting = await this.rpcClient.call("servicename.helloRPC", name);
            res.send(greeting);
        } catch (err) {
            res.status(500).send(err);
        }
    }

    async handleHTTPHelloPub(req: any, res: any) {
        const name = getName(req);
        try {
            await this.publisher.publish("servicename.helloEvent", { name });
            res.send("Message sent");
        } catch (err) {
            res.status(500).send(err);
        }
    }

    async handleRPCHello(...inputs: any[]): Promise<any> {
        if (inputs.length === 0) {
            throw new Error("Message accepted but input is invalid");
        }
        const name = inputs[0] as string
        return greet(name);
    }

    async handleEventHello(msg: Message) {
        const { name } = msg;
        this.names.push(name);
    }

}