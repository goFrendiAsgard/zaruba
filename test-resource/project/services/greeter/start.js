const NATS = require("nats");
const greeting = require("./greeting-lib/greeting")

const nats = NATS.connect();

const subscribeEvent = process.env.getMessageEvent || "foo";
const publishEvent = process.env.sendMessageEvent || "bar";

console.log(`Listen to ${subscribeEvent} event`);
nats.subscribe(subscribeEvent, (receivedMessage) => {
    console.log(`Received a message: ${receivedMessage}`);
    const publishedMessage = greeting.greet(receivedMessage);
    console.log(`Publish to ${publishEvent} event: ${publishedMessage}`);
    nats.publish(publishEvent, publishedMessage);
});