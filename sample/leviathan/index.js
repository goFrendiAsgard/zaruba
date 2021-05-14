var express = require('express');
var http = require('http')
var app = express();

port = process.env.PORT || '3000';
masterHostName = process.env.MASTER_HOSTNAME || 'localhost';
masterPort = process.env.MASTER_PORT || '3000';

app.get('/', function(req, res){
   res.send(`bahemot ${mode} on port ${port}`);
});

var options = {
    hostname: masterHostName,
    port: masterPort,
    path: '/',
    method: 'GET'
}
var req = http.request(options, res => {
    if (res.statusCode == 200) {
        app.listen(port);
        console.log(`leviathan on port ${port}`);
    } else {
        console.error('bad response');
    }
});
req.on('error', error => {
    console.error(error)
});
req.end();