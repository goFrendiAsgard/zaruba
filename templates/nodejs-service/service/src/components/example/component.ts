import { Express } from "express";
import { Publisher, Subscriber, RPCServer, RPCClient } from "../../transport";
import { App, Comp } from "../../core";
import { Config } from "../../config";
import { Controller } from "./controller";

export class Component implements Comp {

    private controller: Controller = new Controller(this.publisher, this.rpcClient);

    constructor(private config: Config, private app: App, private router: Express, private publisher: Publisher, private subscriber: Subscriber, private rpcServer: RPCServer, private rpcClient: RPCClient) {
    }

    public setup() {
        this.route();
        this.registerRPCHandler();
        this.registerMessageHandler();
    }

    private route() {
        const controller = this.controller;
        // Use the same HTTP Handler for multiple URLS
        this.router.get("/hello", controller.handleHTTPHello.bind(controller));
        this.router.get("/hello/:name", controller.handleHTTPHello.bind(controller));
        this.router.post("/hello", controller.handleHTTPHello.bind(controller));
        // Use HTTP Handler that take state from component
        this.router.get("/hello-all", controller.handleHTTPHelloAll.bind(controller));
        // Trigger RPC Call
        this.router.get("/hello-rpc", controller.handleHTTPHelloRPC.bind(controller));
        this.router.get("/hello-rpc/:name", controller.handleHTTPHelloRPC.bind(controller));
        this.router.post("/hello-rpc", controller.handleHTTPHelloRPC.bind(controller));
        // Trigger Publisher
        this.router.get("/hello-pub", controller.handleHTTPHelloPub.bind(controller));
        this.router.get("/hello-pub/:name", controller.handleHTTPHelloPub.bind(controller));
        this.router.post("/hello-pub", controller.handleHTTPHelloPub.bind(controller));
    }

    private registerRPCHandler() {
        const controller = this.controller;
        this.rpcServer.registerHandler("helloRPC", controller.handleRPCHello.bind(controller));
    }

    private registerMessageHandler() {
        const controller = this.controller;
        this.subscriber.registerHandler("hello", controller.handleEventHello.bind(controller));
    }

}