import { App } from "../../core";
import { Config } from "../../config";

export function createSetup(app: App, config: Config): () => void {
    return () => {
        const serviceName = config.serviceName;
        const r = app.router();

        r.all("/", (_, res) => {
            res.send({
                service_name: serviceName,
            });
        });

    }
}