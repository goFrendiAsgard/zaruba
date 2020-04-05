import { RequestHandler } from "express";

export function createHttpLogger(logger: Console): RequestHandler {
    return async (req, res, next) => {
        const processName = `HTTP REQUEST: ${req.method} ${req.url}`;
        logger.time(processName);
        try {
            next();
        } catch (err) {
            logger.error(err);
            res.sendStatus(500);
        }
        logger.timeEnd(processName);
    }
}