function createHelloPubSubHandler(context, pubSub) {
    return function hello(req, res) {
        const config = context.getConfig();
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
            pubSub.publish(config.defaultRmqEvent, { name });
        } catch (err) {
            config.logger.error(err);
            return res.status(500).send("Sending error");
        }
        res.send("Message sent");
    }
}

module.exports = { createHelloPubSubHandler };