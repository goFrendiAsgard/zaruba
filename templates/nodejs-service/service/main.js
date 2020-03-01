const express = require('express');
const app = express();
const port = process.env.port || 3000;

// liveness, typically used if you want to deploy the service into kubernetes cluster
app.get('/liveness', (req, res) => {
    res.send('I am alive');
});

// readiness, typically used if you want to deploy the service into kubernetes cluster
app.get('/readiness', (req, res) => {
    res.send('I am ready');
});

// legendary hello world :)
app.get('/', (req, res) => {
    res.send('Hello world');
})

app.listen(port, () => console.log(`Example app listening on port ${port}!`))