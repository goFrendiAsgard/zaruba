import { Setting } from "../setting";

export function setup(s: Setting) {

    s.router.get("/liveness", function (req, res) {
        const httpCode = s.ctx.status.isAlive ? 200 : 500;
        const { serviceName } = s.ctx.config;
        res.status(httpCode).send({
            service_name: serviceName,
            is_alive: s.ctx.status.isAlive
        })
    });

    s.router.get("/readiness", function (req, res) {
        const httpCode = s.ctx.status.isReady ? 200 : 500;
        const { serviceName } = s.ctx.config;
        res.status(httpCode).send({
            service_name: serviceName,
            is_alive: s.ctx.status.isReady
        })
    });

}