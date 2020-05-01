import { Config } from "./config";
import { MainApp, createSetup } from "./core";
import * as defaultComponent from "./components/defaultcomponent";
import * as monitoring from "./components/monitoring";
import * as example from "./components/example";

function main() {

    // create config and app
    const config = new Config();
    console.log("CONFIG:", JSON.stringify(config));
    const app = new MainApp(
        config.httpPort,
        config.globalRmqConnectionString,
        config.localRmqConnectionString,
    );

    // setup components
    app.setup([
        defaultComponent.createSetup(app, config),          // setup default
        monitoring.createSetup(app, config),                // setup monitoring
        createSetup(new example.Component(app, config)),    // setup example
    ]);

    // run
    app.run();

}

main();