import { Config } from "./config";
import { Application } from "./core";
import * as monitoring from "./components/monitoring";
import * as example from "./components/example";

function main() {
    // create config and app
    const config = new Config();
    console.log("CONFIG:", config);
    const app = new Application(
        config.httpPort,
        config.globalRmqConnectionString,
        config.localRmqConnectionString,
    );
    // setup components
    app.setup([
        monitoring.createSetup(app, config),                // setup monitoring
        () => (new example.Component(app, config)).setup(), // setup example
    ]);
    // run
    app.run();
}

main();