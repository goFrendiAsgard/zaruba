import { Setting } from "../components/setting";
import { setup as monitoringSetup } from "../components/monitoring";
import { setup as exampleSetup } from "../components/example";

export function setup(s: Setting) {
    monitoringSetup(s);
    exampleSetup(s);
}