const { Config } = require("../config")

class Context {

    constructor() {
        if (!("status" in Context)) {
            Context.status = {
                isAlive: true,
                isReady: true
            };
        }
        if (!("localCache" in Context)) {
            Context.localCache = {};
        }
        if (!("config" in Context)) {
            Context.config = new Config();
        }
    }

    initLocalCache(key, val) {
        if (!(key in Context.localCache)) {
            this.setLocalCache(key, val);
        }
        return this;
    }

    setLocalCache(key, val) {
        Context.localCache[key] = val;
        return this;
    }

    getLocalCache(key) {
        return Context.localCache[key];
    }

    getConfig() {
        return Context.config;
    }

    getLiveness() {
        return Context.status.isAlive;
    }

    setLiveness(liveness) {
        Context.status.isAlive = liveness;
    }

    getReadiness() {
        return Context.status.isReady;
    }

    setReadiness(readiness) {
        Context.status.isReady = readiness;
    }

}

module.exports = {
    Context
};
