import { Express } from "express";
import { Publisher, Subscriber, RPCClient, RPCServer } from "../transport";

export type SetupComponent = () => void;

export interface App {
    liveness: () => boolean;
    readiness: () => boolean;
    setLiveness: (liveness: boolean) => void;
    setReadiness: (readiness: boolean) => void;
    setup: (setupComponents: SetupComponent[]) => void;
    run: () => void;
}