import { Express } from "express";
import { Config } from "../../config";

export function createSetup(config: Config, router: Express): () => void {
    return () => {
        const serviceName = config.serviceName;

        router.all("/", (_, res) => {
            res.send({
                service_name: serviceName,
            });
        });

    }
}