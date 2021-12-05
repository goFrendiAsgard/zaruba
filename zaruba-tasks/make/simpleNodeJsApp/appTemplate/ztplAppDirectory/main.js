const process = require('process');

// get httpPort
const httpPort = process.env.APP_HTTP_PORT || 3000

const server = require("http").createServer((req, res) => {
  res.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
  res.write("Hello world 🐸");
  res.end();
});
 
// serve
server.listen(httpPort);
console.log(`Serve HTTP on port ${httpPort}`);