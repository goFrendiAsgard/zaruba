export type Message = { [key: string]: any };
export type RPCHandler = (...inputs: any[]) => any;
export type EventHandler = (input: Message) => void;

export interface RPCClient {
    call: (functionName: string, ...inputs: any[]) => Promise<any>;
}

export interface RPCServer {
    registerHandler: (functionName: string, handler: RPCHandler) => RPCServer;
    serve: () => Promise<void>;
}

export interface Publisher {
    publish: (eventName: string, msg: Message) => Promise<void>;
}

export interface Subscriber {
    registerHandler: (eventName: string, handler: EventHandler) => Subscriber;
    subscribe: () => Promise<void>;
}

export interface RmqEventConfig {
    queueName: string;
    exchangeName: string;
    rpcTimeout?: number;
    deadLetterExchange: string;
    deadLetterQueue: string;
    ttl: number;
    autoAck: boolean;
}