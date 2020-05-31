import { SetupComponent } from "./interfaces";

export function createSetup(obj: { setup: SetupComponent }): SetupComponent {
    return () => {
        obj.setup();
    }
}