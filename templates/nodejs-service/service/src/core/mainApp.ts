import { Express } from "express";
import express from "express";
import bodyParser from "body-parser";
import { App, SetupComponent } from "./interfaces";
import { Publisher, Subscriber, RPCServer, RPCClient, RmqPublisher, RmqSubscriber, RmqRPCServer, RmqRPCClient } from "../transport";
import { createHttpLogger, logExpressRoutes } from "./expressMiddlewares";

export class MainApp implements App {

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
            this._readiness = false;
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
