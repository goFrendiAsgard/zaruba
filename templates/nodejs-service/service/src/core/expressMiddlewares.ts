import { Express, RequestHandler } from "express";

export function createHttpLogger(logger: Console): RequestHandler {
    return async (req, res, next) => {
        const startHr = process.hrtime();
        try {
            await next();
        } catch (err) {
            logger.error(err);
            res.sendStatus(500);
        }
        const durationHr = process.hrtime(startHr);
        const duration = durationHr[0] * 1000000 + durationHr[1] / 1000;
        logger.log(`HTTP Request ${req.method} ${req.url} ${res.statusCode} ${duration} ms`);
    }
}

export function logExpressRoutes(router: Express, logger: Console) {
    try {
        for (let layer of router._router.stack) {
            if (layer.route && layer.route.path && layer.route.methods) {
                const path = layer.route.path;
                for (let method in layer.route.methods) {
                    const shownMethod = method.toUpperCase().padEnd(15, " ");
                    logger.log(`${shownMethod}\t ${path}`);
                }
            }
        }
    } catch (err) {
        logger.error(err);
    }
}
