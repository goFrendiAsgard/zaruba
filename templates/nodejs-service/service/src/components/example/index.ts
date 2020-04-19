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

        // Simple HTTP Handler
        r.all("/", (_, res) => res.send("servicename"));

        // More complex HTTP Handler, with side-effect
        r.get("/toggle-readiness", (_, res) => {
            this.app.setReadiness(!this.app.readiness());
            const readiness = this.app.readiness();
            res.send(`Readiness: ${readiness}`);
        });

        // Use the same HTTP Handler for multiple URLS
        const handleHTTPHello = (req: any, res: any) => this.handleHTTPHello.call(this, req, res);
        r.get("/hello", handleHTTPHello);
        r.get("/hello/:name", handleHTTPHello);
        r.post("/hello", handleHTTPHello);

        // Use HTTP Handler that take state from component
        const handleHTTPHelloAll = (req: any, res: any) => this.handleHTTPHelloAll.call(this, req, res);
        r.get("/hello-all", handleHTTPHelloAll);

        // Trigger RPC Call
        const handleHTTPHelloRPC = (req: any, res: any) => this.handleHTTPHelloRPC.call(this, req, res);
        r.get("/hello-rpc", handleHTTPHelloRPC);
        r.get("/hello-rpc/:name", handleHTTPHelloRPC);
        r.post("/hello-rpc", handleHTTPHelloRPC);

        // Trigger RPC Call
        const handleHTTPHelloPub = (req: any, res: any) => this.handleHTTPHelloPub.call(this, req, res);
        r.get("/hello-pub", handleHTTPHelloPub);
        r.get("/hello-pub/:name", handleHTTPHelloPub);
        r.post("/hello-pub", this.handleHTTPHelloPub);

        // Serve RPC
        const handleRPCHello = (...inputs: any[]) => this.handleRPCHello.call(this, ...inputs);
        rpcServer.registerHandler("servicename.helloRPC", handleRPCHello);

        // Event
        const handleEventHello = (msg: Message) => this.handleEventHello.call(this, msg);
        subscriber.registerHandler("servicename.helloEvent", handleEventHello);

    }

    handleHTTPHello(req: any, res: any) {
        const name = getName(req);
        res.send(greet(name));
    }

    handleHTTPHelloAll(req: any, res: any) {
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

    handleHTTPHelloPub(req: any, res: any) {
        const publisher = this.app.globalPublisher();
        const name = getName(req);
        try {
            publisher.publish("servicename.helloEvent", { name });
            res.send("Message sent");
        } catch (err) {
            res.status(500).send(err);
        }
    }

    handleRPCHello(...inputs: any[]): any {
        if (inputs.length === 0) {
            throw new Error("Message accepted but input is invalid");
        }
        const name = inputs[0] as string
        return greet(name);
    }

    handleEventHello(msg: Message) {
        const { name } = msg;
        this.names.push(name);
    }

}