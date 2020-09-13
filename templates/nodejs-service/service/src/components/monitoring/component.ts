import { Express } from "express";
import { App, Comp } from "../../core";
import { Config } from "../../config";

export class Component implements Comp {
    constructor(private config: Config, private app: App, private router: Express) { }

    setup() {
        const serviceName = this.config.serviceName;

        this.router.get("/liveness", (_, res) => {
            const liveness = this.app.liveness();
            const httpCode = liveness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_alive: liveness,
            });
        });

        this.router.get("/readiness", (_, res) => {
            const readiness = this.app.readiness();
            const httpCode = readiness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_ready: readiness,
            });
        });
    }
}