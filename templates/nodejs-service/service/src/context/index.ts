import { Config } from "../config";
import { Message } from "../transport";

export class Status {
    isAlive: boolean;
    isReady: boolean;
    constructor() {
        this.isAlive = true;
        this.isReady = true;
    }
}

export class Context {
    static globalContext: Context

    config: Config;
    status: Status;
    localCache: Message;

    constructor() {
        if (Context.globalContext) {
            throw new Error("context initialized, use Context.getInstance() instead");
        }
        this.config = new Config();
        this.status = new Status();
        this.localCache = {};
        Context.globalContext = this;
    }

    initLocalCache(key: string, val: any) {
        if (key in this.localCache) {
            return
        }
        this.localCache[key] = val;
    }

    getInstance() {
        return Context.globalContext;
    }

}
