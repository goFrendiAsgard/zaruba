import { RequestHandler } from "express";

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