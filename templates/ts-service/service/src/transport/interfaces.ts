export type Message = { [key: string]: any };
export type RPCHandler = (...inputs: any[]) => any;
export type EventHandler = (input: Message) => void;

export interface RPCClient {
    call: (functionName: string, ...inputs: any[]) => any;
    setLogger: (logger: Console) => RPCClient;
}

export interface RPCServer {
    registerHandler: (functionName: string, handler: RPCHandler) => RPCServer;
    setLogger: (logger: Console) => RPCServer;
    Serve: () => void;
}

export interface Publisher {
    publish: (functionName: string, msg: Message) => void;
    setLogger: (logger: Console) => Publisher
}

export interface Subscriber {
    registerHandler: (functionName: string, handler: EventHandler) => Subscriber;
    setLogger: (logger: Console) => Subscriber;
    Subscribe: () => void;
}