export type Message = { [key: string]: any };
export type RPCHandler = (...inputs: any[]) => any;
export type EventHandler = (input: Message) => void;

export interface RPCClient {
    call: (functionName: string, ...inputs: any[]) => Promise<any>;
    setLogger: (logger: Console) => RPCClient;
}

export interface RPCServer {
    registerHandler: (functionName: string, handler: RPCHandler) => RPCServer;
    setLogger: (logger: Console) => RPCServer;
    serve: () => void;
}

export interface Publisher {
    publish: (eventName: string, msg: Message) => void;
    setLogger: (logger: Console) => Publisher
}

export interface Subscriber {
    registerHandler: (eventName: string, handler: EventHandler) => Subscriber;
    setLogger: (logger: Console) => Subscriber;
    subscribe: () => void;
}