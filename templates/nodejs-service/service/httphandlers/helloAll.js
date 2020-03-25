
function createHelloAllHandler(context) {
    return function (req, res) {
        context.initLocalCache("names", []);
        names = context.getLocalCache("names");
        const greeting = names.length == 0 ? "Hello, everyone..." : "Hello, " + names.join(", ") + " and everyone";
        res.send(greeting);
    }
}

module.exports = { createHelloAllHandler };