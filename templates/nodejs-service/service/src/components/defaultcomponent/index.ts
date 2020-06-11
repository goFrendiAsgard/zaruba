import { Express } from "express";
import { Config } from "../../config";
import { Comp } from "../../core";

export class Component implements Comp {
    constructor(private config: Config, private router: Express) { }

    setup() {
        const serviceName = this.config.serviceName;

        this.router.all("/", (_, res) => {
            res.send({
                service_name: serviceName,
            });
        });
    }
}