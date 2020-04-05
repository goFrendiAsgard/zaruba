import { greet, greetEveryone } from "./serviceGreeting";
import { Context } from "../../context";
import { RPCClient, Publisher } from "../../transport";

function getName(req: any) {
    return req.params.name || req.query.name || req.body.name || "";
}

export function greetHttpController(req: any, res: any) {
    const name = getName(req);
    res.send(greet(name));
}

export function createGreetEveryoneHttpController(ctx: Context) {
    return (req: any, res: any) => {
        ctx.initLocalCache("names", []);
        const names = ctx.localCache["names"];
        res.send(greetEveryone(names));
    }
}

export function createGreetRPCHttpController(rpcClient: RPCClient, functionName: string) {
    return async (req: any, res: any) => {
        const name = getName(req);
        try {
            const greeting = await rpcClient.call(functionName, name);
            res.send(greeting);
        } catch (err) {
            res.status(500).send(err);
        }
    }
}

export function createGreetPublishHttpController(publisher: Publisher, eventName: string) {
    return (req: any, res: any) => {
        const name = getName(req);
        try {
            publisher.publish(eventName, { name });
            res.send("Message sent");
        } catch (err) {
            res.status(500).send(err);
        }
    }
}