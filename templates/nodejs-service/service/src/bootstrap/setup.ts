import { Setting } from "../components/setting";
import { setup as monitoringSetup } from "../components/monitoring";

export function setup(s: Setting) {
    monitoringSetup(s);

    // TODO: Add your custom handlers here
}