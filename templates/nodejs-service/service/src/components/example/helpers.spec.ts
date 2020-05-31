import { getName } from "./helpers";

test("getName empty", () => {
    const name = getName({ params: {}, query: {}, body: {} });
    expect(name).toBe("");
});

test("getName params", () => {
    const name = getName({ params: { name: "Kouga" }, query: { name: "query" }, body: { name: "body" } });
    expect(name).toBe("Kouga");
});

test("getName query", () => {
    const name = getName({ params: {}, query: { name: "Kouga" }, body: { name: "body" } });
    expect(name).toBe("Kouga");
});


test("getName body", () => {
    const name = getName({ params: {}, query: {}, body: { name: "Kouga" } });
    expect(name).toBe("Kouga");
});
