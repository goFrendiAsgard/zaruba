import * as process from 'process';
import * as http from 'http';

// get httpPort
const httpPort = process.env.APP_HTTP_PORT || 3000

const server = http.createServer((req: http.IncomingMessage, res: http.ServerResponse) => {
  res.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
  res.write("Hello world 🐸");
  res.end();
});
 
// serve
server.listen(httpPort);
console.log(`Serve HTTP on port ${httpPort}`);