import { logExpressRoutes } from "../transport";
import { Setting } from "../components/setting";

export function run(s: Setting) {
    s.rpcServers.main.serve();
    s.rpcServers.secondary.serve();
    s.subscribers.main.subscribe();
    s.router.listen(s.ctx.config.httpPort, () => {
        logExpressRoutes(s.router, s.ctx.config.logger);
        s.ctx.config.logger.log(`Run at port ${s.ctx.config.httpPort}`);
    });
}
