import { Express } from "express";
import { App, Comp } from "./interfaces";
import { logExpressRoutes } from "./logExpressRoutes";
import { Subscriber, RPCServer } from "../transport";

export interface MainAppConfig {
    logger: Console,
    router: Express,
    subscribers: Subscriber[],
    rpcServers: RPCServer[],
    httpPort: number,
}

export class MainApp implements App {

    private _readiness: boolean;
    private _liveness: boolean;
    private _httpPort: number;
    private _router: Express;
    private _rpcServers: RPCServer[];
    private _subscribers: Subscriber[];
    private _logger: Console;


    constructor(config: MainAppConfig) {
        this._httpPort = config.httpPort;
        this._readiness = false;
        this._liveness = false;
        this._logger = config.logger;
        this._router = config.router;
        this._subscribers = config.subscribers;
        this._rpcServers = config.rpcServers;
    }

    setup(components: Comp[]) {
        for (let component of components) {
            component.setup();
        }
    }

    run() {
        const pRouter = new Promise((resolve, reject) => {
            this._router.listen(this._httpPort, () => {
                this._logger.info(`Run at port ${this._httpPort}`);
                logExpressRoutes(this._router, this._logger);
                this._liveness = true;
                this._readiness = true;
                resolve();
            }).on("error", (err) => reject(err));
        });
        const promises: Promise<any>[] = [pRouter];
        for (let rpcServer of this._rpcServers) {
            promises.push(rpcServer.serve());
        }
        for (let subscriber of this._subscribers) {
            promises.push(subscriber.subscribe())
        }
        Promise.all(promises).catch((err) => {
            this._liveness = false;
            this._readiness = false;
            this._logger.error(err);
        });
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
