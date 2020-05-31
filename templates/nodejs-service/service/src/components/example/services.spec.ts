import { greet, greetEveryone } from "./services";

test("greet empty parameter", () => {
    const greetings = greet("");
    expect(greetings).toBe("Hello world !!!");
});

test("greet non empty parameter", () => {
    const greetings = greet("Kouga");
    expect(greetings).toBe("Hello Kouga");
});

test("greetEveryone empty parameter", () => {
    const greetings = greetEveryone([]);
    expect(greetings).toBe("Hello everyone !!!");
});

test("greetEveryone non empty parameter", () => {
    const greetings = greetEveryone(["Kouga", "Kaoru"]);
    expect(greetings).toBe("Hello Kouga, Kaoru, and everyone");
});
