const http = require("http");
const NATS = require("nats");
const greeting = require("./greeting-lib/greeting")

const nats = NATS.connect();

const port = process.env.port || 80;
const subscribeEvent = process.env.sendMessageEvent || "bar";
const publishEvent = process.env.getMessageEvent || "foo";

//create a server object:
console.log(`Serve HTTP on port ${port}`);
http.createServer(function (req, res) {
    nats.subscribe(subscribeEvent,
        (receivedMessage) => {
            console.log(`Receive request from url: ${req.url}`);
            receivedMessage = receivedMessage || greeting.defaultGreeting
            res.end(receivedMessage);
        },
        (error) => {
            console.error(error);
            res.end("Internal server error");
        }
    );
    nats.publish(publishEvent, req.url);
}).listen(port);