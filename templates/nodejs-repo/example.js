import { format } from 'util';

function greet(name) {
    return format("Hi %s, greeting from NodeJs", name);
}

module.exports = { greet }