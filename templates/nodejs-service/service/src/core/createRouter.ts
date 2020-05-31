import express from "express";
import { Express } from "express";
import bodyParser from "body-parser";
import { createHttpLogger } from "./expressMiddlewares";

export function createRouter(logger: Console): Express {
    const router = express();
    router.use(bodyParser.urlencoded({ extended: false }));
    router.use(bodyParser.json());
    router.use(createHttpLogger(logger));
    return router;
}