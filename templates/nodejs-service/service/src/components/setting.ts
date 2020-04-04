import express from "express";
import { Express } from "express";
import { Context } from "../context";
import { Publisher, Subscriber, RPCServer, RPCClient, RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient } from "../transport";

export class Setting {
    static globalSetting: Setting;

    ctx: Context;
    router: Express;
    publishers: Publishers;
    subscribers: Subscribers;
    rpcServers: RPCServers;
    rpcClients: RPCClients;

    constructor(ctx: Context, router: Express, publishers: Publishers, subscribers: Subscribers, rpcServers: RPCServers, rpcClients: RPCClients) {
        if (Setting.globalSetting) {
            throw new Error("setting initialized, user Setting.getInstance() instead");
        }
        this.ctx = ctx;
        this.router = router;
        this.publishers = publishers;
        this.subscribers = subscribers;
        this.rpcServers = rpcServers;
        this.rpcClients = rpcClients;
        Setting.globalSetting = this;
    }

    getInstance() {
        return Setting.globalSetting;
    }
}

export class Publishers {
    main: Publisher;
    constructor(main: Publisher) {
        this.main = main;
    }
}

export class Subscribers {
    main: Subscriber;
    constructor(main: Subscriber) {
        this.main = main;
    }
}

export class RPCServers {
    main: RPCServer;
    secondary: RPCServer;
    constructor(main: RPCServer, secondary: RPCServer) {
        this.main = main;
        this.secondary = secondary;
    }
}

export class RPCClients {
    mainLoopBack: RPCClient;
    secondaryLoopBack: RPCClient;
    constructor(mainLoopBack: RPCClient, secondaryLoopBack: RPCClient) {
        this.mainLoopBack = mainLoopBack;
        this.secondaryLoopBack = secondaryLoopBack;
    }
}
