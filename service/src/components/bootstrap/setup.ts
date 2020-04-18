import { Setting } from "../setting";
import { setup as monitoringSetup } from "../monitoring";
import { setup as exampleSetup } from "../example";

export function setup(s: Setting) {
    monitoringSetup(s);
    exampleSetup(s);
}
