export function greet(name: string): string {
    if (name === "") {
        return "Hello world !!!";
    }
    return `Hello ${name}`;
}

export function greetEveryone(names: string[]): string {
    if (names.length === 0) {
        return "Hello everyone !!!";
    }
    const joinedNames = names.join(", ");
    return `Hello ${joinedNames}, and everyone`;
}