import { Message } from "../../transport";
import { App } from "../../core";
import { Config } from "../../config";
import { getName } from "./helpers";
import { greet, greetEveryone } from "./service";

export class Component {
    private names: string[];
    private app: App;
    private config: Config;

    constructor(app: App, config: Config) {
        this.names = [];
        this.app = app;
        this.config = config;
    }

    setup() {
        const r = this.app.router();
        const rpcServer = this.app.globalRPCServer();
        const subscriber = this.app.globalSubscriber();

        // Use the same HTTP Handler for multiple URLS
        r.get("/hello", this.handleHTTPHello.bind(this));
        r.get("/hello/:name", this.handleHTTPHello.bind(this));
        r.post("/hello", this.handleHTTPHello.bind(this));

        // Use HTTP Handler that take state from component
        r.get("/hello-all", this.handleHTTPHelloAll.bind(this));

        // Trigger RPC Call
        r.get("/hello-rpc", this.handleHTTPHelloRPC.bind(this));
        r.get("/hello-rpc/:name", this.handleHTTPHelloRPC.bind(this));
        r.post("/hello-rpc", this.handleHTTPHelloRPC.bind(this));

        // Trigger RPC Call
        r.get("/hello-pub", this.handleHTTPHelloPub.bind(this));
        r.get("/hello-pub/:name", this.handleHTTPHelloPub.bind(this));
        r.post("/hello-pub", this.handleHTTPHelloPub.bind(this));

        // Serve RPC
        rpcServer.registerHandler("servicename.helloRPC", this.handleRPCHello.bind(this));

        // Event
        subscriber.registerHandler("servicename.helloEvent", this.handleEventHello.bind(this));

    }

    async handleHTTPHello(req: any, res: any) {
        const name = getName(req);
        res.send(greet(name));
    }

    async handleHTTPHelloAll(req: any, res: any) {
        res.send(greetEveryone(this.names));
    }

    async handleHTTPHelloRPC(req: any, res: any) {
        const rpcClient = this.app.globalRPCClient();
        const name = getName(req);
        try {
            const greeting = await rpcClient.call("servicename.helloRPC", name);
            res.send(greeting);
        } catch (err) {
            res.status(500).send(err);
        }
    }

    async handleHTTPHelloPub(req: any, res: any) {
        const publisher = this.app.globalPublisher();
        const name = getName(req);
        try {
            await publisher.publish("servicename.helloEvent", { name });
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