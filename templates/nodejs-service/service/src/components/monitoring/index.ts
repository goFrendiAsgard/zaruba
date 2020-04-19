import { App } from "../../core";
import { Config } from "../../config";

export function createSetup(app: App, config: Config): () => void {
    return () => {
        const serviceName = config.serviceName;
        const r = app.router();

        r.get("/liveness", (_, res) => {
            const liveness = app.liveness();
            const httpCode = liveness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_alive: liveness,
            });
        });

        r.get("/readiness", (_, res) => {
            const readiness = app.readiness();
            const httpCode = readiness ? 200 : 500;
            res.status(httpCode).send({
                service_name: serviceName,
                is_alive: readiness,
            });
        });

    }

}