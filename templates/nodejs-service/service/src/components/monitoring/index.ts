import { Express } from "express";
import { App } from "../../core";
import { Config } from "../../config";

export function createSetup(config: Config, app: App, router: Express): () => void {
    return () => {
        const serviceName = config.serviceName;

        router.get("/liveness", (_, res) => {
            const liveness = app.liveness();
            const httpCode = liveness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_alive: liveness,
            });
        });

        router.get("/readiness", (_, res) => {
            const readiness = app.readiness();
            const httpCode = readiness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_ready: readiness,
            });
        });

    }
}