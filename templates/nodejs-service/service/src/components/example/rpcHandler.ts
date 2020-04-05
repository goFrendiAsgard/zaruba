import { greet } from "./serviceGreeting";

export function greetRpcController(...inputs: any[]): any {
    if (inputs.length === 0) {
        throw new Error("Message accepted but input is invalid");
    }
    const name = inputs[0] as string
    return greet(name);
}