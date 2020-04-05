import { Context } from "../../context";
import { EventHandler, Message } from "../../transport";

export function createRegisterPersonHandler(ctx: Context): EventHandler {
    return (msg: Message) => {
        const { name } = msg;
        ctx.initLocalCache("names", []);
        if (name != "") {
            const { names } = ctx.localCache;
            names.push(name);
            ctx.localCache["names"] = names;
        }
    }
}