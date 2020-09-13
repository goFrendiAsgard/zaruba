package transport

import (
	"github.com/streadway/amqp"
)

func rmqDeclareAndBindQueue(ch *amqp.Channel, exchangeName, queueName string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (q amqp.Queue, err error) {
	// declare exchange
	if err = rmqDeclareFanoutExchange(ch, exchangeName); err != nil {
		return q, err
	}
	// declare queue
	q, err = rmqDeclareQueue(ch, queueName, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return q, err
	}
	// bind queue to exchange
	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	return q, err
}

func rmqDeclareQueue(ch *amqp.Channel, queueName string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (q amqp.Queue, err error) {
	return ch.QueueDeclare(queueName, durable, autoDelete, exclusive, noWait, args)
}

func rmqDeclareFanoutExchange(ch *amqp.Channel, exchangeName string) (err error) {
	return ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func rmqConsume(ch *amqp.Channel, queueName string, durable, autoDelete, exclusive, noWait, autoAck bool, args amqp.Table) (rmqMessages <-chan amqp.Delivery, err error) {
	if _, err = rmqDeclareQueue(ch, queueName, durable, autoDelete, exclusive, noWait, args); err != nil {
		return rmqMessages, err
	}
	// start consume
	return ch.Consume(
		queueName, // queue
		"",        // consumer
		autoAck,   // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}

func rmqPublish(ch *amqp.Channel, exchangeName, routingKey string, data amqp.Publishing) (err error) {
	if exchangeName != "" {
		if err = rmqDeclareFanoutExchange(ch, exchangeName); err != nil {
			return err
		}
	}
	return ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		data,         // data
	)
}
