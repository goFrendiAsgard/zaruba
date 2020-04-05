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

export interface Publishers {
    main: Publisher;
}

export interface Subscribers {
    main: Subscriber;
}

export interface RPCServers {
    main: RPCServer;
    secondary: RPCServer;
}

export interface RPCClients {
    mainLoopBack: RPCClient;
    secondaryLoopBack: RPCClient;
}