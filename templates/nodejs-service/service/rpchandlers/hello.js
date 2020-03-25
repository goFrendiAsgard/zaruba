function hello(input) {
    const name = input["name"];
    const output = { greeting: `Hello ${name}` };
    return output;
}

module.exports = { hello };