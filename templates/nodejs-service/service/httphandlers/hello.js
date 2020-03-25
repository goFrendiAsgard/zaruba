function hello(req, res) {
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
    res.send(`Hello ${name} !!!`);
}

module.exports = { hello };