import { Setting } from "../setting";
import { greetHttpController, createGreetEveryoneHttpController, createGreetRPCHttpController, createGreetPublishHttpController } from "./httpController";
import { greetRpcController } from "./rpcHandler";
import { createRegisterPersonHandler } from "./eventHandler";

export function setup(s: Setting) {

    // HTTP EXAMPLE =======================================================================================

    // Example: simple HTTP Handler
    s.router.all("/", (req, res) => res.send("servicename"));

    // Example: More complex HTTP handler, with side-effect
    s.router.get("/toggle-readiness", (req, res) => {
        s.ctx.status.isReady = !s.ctx.status.isReady;
        res.send(`Readiness: ${s.ctx.status.isReady}`);
    });

    // Example: Use HTTP Handler from greeting component
    s.router.get("/hello", greetHttpController);
    s.router.post("/hello", greetHttpController);
    s.router.get("/hello/:name", greetHttpController);

    const greetEveryoneHttpController = createGreetEveryoneHttpController(s.ctx);
    s.router.get("/hello-all", greetEveryoneHttpController);
    s.router.post("/hello-all", greetEveryoneHttpController);
    s.router.get("/hello-all/:name", greetEveryoneHttpController);

    // RPC EXAMPLE ========================================================================================

    // Example: RPC Handler  (Main)
    s.rpcServers.main.registerHandler("greetRPC", greetRpcController);

    // Example: HTTP handler to trigger RPC
    const greetRpcHttpController = createGreetRPCHttpController(s.rpcClients.mainLoopBack, "greetRPC")
    s.router.get("/hello-rpc", greetRpcHttpController)
    s.router.post("/hello-rpc", greetRpcHttpController)
    s.router.get("/hello-rpc/:name", greetRpcHttpController)

    // RPC EXAMPLE ========================================================================================

    // Example: RPC Handler  (Main)
    s.rpcServers.secondary.registerHandler("greetRPC", greetRpcController);

    // Example: HTTP handler to trigger RPC
    const secondaryGreetRpcHttpController = createGreetRPCHttpController(s.rpcClients.secondaryLoopBack, "greetRPC")
    s.router.get("/hello-secondary-rpc", secondaryGreetRpcHttpController)
    s.router.post("/hello-secondary-rpc", secondaryGreetRpcHttpController)
    s.router.get("/hello-secondary-rpc/:name", secondaryGreetRpcHttpController)

    // PUB SUB EXAMPLE =====================================================================================

    // Example: Event Handler
    const registerPersonEvenHandler = createRegisterPersonHandler(s.ctx)
    s.subscribers.main.registerHandler("personRegistered", registerPersonEvenHandler)

    // Example: HTTP handler to publish event
    const greetPublishHTTPController = createGreetPublishHttpController(s.publishers.main, "personRegistered")
    s.router.get("/hello-pub", greetPublishHTTPController)
    s.router.post("/hello-pub", greetPublishHTTPController)
    s.router.get("/hello-pub/:name", greetPublishHTTPController)
}

export default { setup };
