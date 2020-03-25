
function createHelloHandler(context) {
	return function (input) {
		const { logger } = context.getConfig();

		name = input.name
		logger.log(`[RMQ PUBSUB] Hello ${name}`);

		// add name to localCache
		context.initLocalCache("names", []);
		names = context.getLocalCache("names");
		names.push(name);
		context.setLocalCache("names", names);
	}
}

module.exports = { createHelloHandler };