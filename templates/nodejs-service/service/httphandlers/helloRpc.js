function createHelloRPCHandler(rpc) {
    return async function hello(req, res) {
        let name = req.params.name;
        if (!name) {
            name = req.query.name;
        }
        if (!name) {
            name = req.body.name;
        }
        if (!name) {
            name = "world";
        }
        try {
            const { greeting } = await rpc.call("servicename", "helloRpc", { name });
            res.send(greeting);
        } catch (err) {
            return res.status(500).send(err);
        }
    }
}

module.exports = { createHelloRPCHandler };