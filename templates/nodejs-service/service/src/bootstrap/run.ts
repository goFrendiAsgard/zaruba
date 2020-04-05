import { logExpressRoutes } from "../transport";
import { Setting } from "../components/setting";

export function run(s: Setting) {

    // create promises
    const promises = [];
    promises.push(s.rpcServers.main.serve());
    promises.push(s.rpcServers.secondary.serve());
    promises.push(s.subscribers.main.subscribe());
    promises.push(s.router.listen(s.ctx.config.httpPort, () => {
        logExpressRoutes(s.router, s.ctx.config.logger);
        s.ctx.config.logger.log(`Run at port ${s.ctx.config.httpPort}`);
    }));

    // handle promises
    Promise.all(promises).catch((reason) => {
        s.ctx.status.isAlive = false;
        s.ctx.status.isReady = false;
        s.ctx.config.logger.error(reason);
    });

}
