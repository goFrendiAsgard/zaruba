import { Express } from "express";
import { Publisher, Subscriber, RPCClient, RPCServer } from "../transport";

export type SetupComponent = () => void;

export interface App {
    logger: () => Console;
    router: () => Express;
    globalPublisher: () => Publisher;
    localPublisher: () => Publisher;
    globalSubscriber: () => Subscriber;
    localSubscriber: () => Subscriber;
    globalRPCServer: () => RPCServer;
    localRPCServer: () => RPCServer;
    globalRPCClient: () => RPCClient;
    localRPCClient: () => RPCClient;
    liveness: () => boolean;
    readiness: () => boolean;
    setLiveness: (liveness: boolean) => void;
    setReadiness: (readiness: boolean) => void;
    setup: (setupComponents: SetupComponent[]) => void;
    run: () => void;
}