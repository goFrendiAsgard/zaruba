export interface Comp {
    setup: () => void | Promise<void>;
}

export interface App {
    liveness: () => boolean;
    readiness: () => boolean;
    setLiveness: (liveness: boolean) => void;
    setReadiness: (readiness: boolean) => void;
    setup: (components: Comp[]) => void;
    run: () => void;
}