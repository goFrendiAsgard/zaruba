import { Express, RequestHandler } from "express";
import express from "express";
import bodyParser from "body-parser";
import { App, SetupComponent } from "./interfaces";
import { Publisher, Subscriber, RPCServer, RPCClient, RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient } from "../transport";

export class Application implements App {
    private static _instance: Application;

    private _readiness: boolean;
    private _liveness: boolean;
    private _logger: Console;
    private _router: Express;
    private _globalPublisher: Publisher;
    private _localPublisher: Publisher;
    private _globalSubscriber: Subscriber;
    private _localSubscriber: Subscriber;
    private _globalRPCServer: RPCServer;
    private _localRPCServer: RPCServer;
    private _globalRPCClient: RPCClient;
    private _localRPCClient: RPCClient;
    private _httpPort: number;

    constructor(httpPort: number, globalRmqConnectionString: string, localRmqConnectionString: string) {
        if (Application._instance) {
            throw new Error("Application initialized, use Application.getInstance() instead");
        }
        this._httpPort = httpPort;
        this._readiness = false;
        this._liveness = false;
        this._logger = console;
        this._router = express();
        this._router.use(bodyParser.urlencoded({ extended: false }));
        this._router.use(bodyParser.json());
        this._router.use(createHttpLogger(this._logger));
        this._globalPublisher = new RmqPublisher(globalRmqConnectionString).setLogger(this._logger);
        this._localPublisher = new RmqPublisher(localRmqConnectionString).setLogger(this._logger);
        this._globalSubscriber = new RmqSubscriber(globalRmqConnectionString).setLogger(this._logger);
        this._localSubscriber = new RmqSubscriber(localRmqConnectionString).setLogger(this._logger);
        this._globalRPCServer = new RmqRPCServer(globalRmqConnectionString).setLogger(this._logger);
        this._localRPCServer = new RmqRPCServer(localRmqConnectionString).setLogger(this._logger);
        this._globalRPCClient = new RmqRPCClient(globalRmqConnectionString).setLogger(this._logger);
        this._localRPCClient = new RmqRPCClient(localRmqConnectionString).setLogger(this._logger);
    }

    getInstance() {
        return Application._instance;
    }

    setup(setupComponents: SetupComponent[]) {
        for (let setup of setupComponents) {
            setup();
        }
    }

    run() {
        Promise.all([
            this._globalRPCServer.serve(),
            this._localRPCServer.serve(),
            this._globalSubscriber.subscribe(),
            this._localSubscriber.subscribe(),
            this._router.listen(this._httpPort, () => {
                logExpressRoutes(this._router, this._logger);
                this._logger.log(`Run at port ${this._httpPort}`);
                this._liveness = true;
                this._readiness = true;
            }),
        ]).catch((err) => {
            this._liveness = false;
            this._liveness = true;
            this._logger.error(err);
        });

    }

    logger() {
        return this._logger;
    }

    router() {
        return this._router;
    }

    globalPublisher() {
        return this._globalPublisher;
    }

    localPublisher() {
        return this._localPublisher;
    }

    globalSubscriber() {
        return this._globalSubscriber;
    }

    localSubscriber() {
        return this._localSubscriber;
    }


    globalRPCServer() {
        return this._globalRPCServer;
    }

    localRPCServer() {
        return this._localRPCServer;
    }

    globalRPCClient() {
        return this._globalRPCClient;
    }

    localRPCClient() {
        return this._localRPCClient;
    }

    liveness() {
        return this._liveness;
    }

    setLiveness(liveness: boolean) {
        this._liveness = liveness;
    }

    readiness() {
        return this._readiness;
    }

    setReadiness(readiness: boolean) {
        this._readiness = readiness;
    }

}


export function createHttpLogger(logger: Console): RequestHandler {
    return async (req, res, next) => {
        const label = `HTTP REQUEST: ${req.method} ${req.url}`;
        logger.time(label);
        try {
            next();
        } catch (err) {
            logger.error(err);
            res.sendStatus(500);
        }
        logger.timeEnd(label);
    }
}

export function logExpressRoutes(router: Express, logger: Console) {
    try {
        for (let layer of router._router.stack) {
            if (layer.route && layer.route.path && layer.route.methods) {
                const path = layer.route.path;
                for (let method in layer.route.methods) {
                    const shownMethod = method.toUpperCase().padEnd(15, " ");
                    logger.log(`${shownMethod}\t ${path}`);
                }
            }
        }
    } catch (err) {
        logger.error(err);
    }
}
