package communication

import "github.com/streadway/amqp"

func createRmqConnectionAndChannel(connectionString string) (conn *amqp.Connection, ch *amqp.Channel, err error) {
	conn, err = amqp.Dial(connectionString)
	if err != nil {
		return conn, ch, err
	}
	ch, err = conn.Channel()
	return conn, ch, err
}

func declareAndBindRmqQueueToExchange(ch *amqp.Channel, queueName string) (q amqp.Queue, err error) {
	// declare exchange
	err = declareRmqFanoutExchange(ch, queueName)
	if err != nil {
		return q, err
	}
	// declare queue
	q, err = declareRmqQueue(ch, queueName)
	if err != nil {
		return q, err
	}
	// bind queue to exchange
	err = ch.QueueBind(
		q.Name,    // queue name
		"",        // routing key
		queueName, // exchange
		false,
		nil,
	)
	return q, err
}

func declareRmqQueue(ch *amqp.Channel, queueName string) (q amqp.Queue, err error) {
	return ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func declareRmqFanoutExchange(ch *amqp.Channel, exchangeName string) (err error) {
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

func rmqConsume(ch *amqp.Channel, q amqp.Queue) (rmqMessages <-chan amqp.Delivery, err error) {
	// start consume
	return ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}

func rmqPublish(ch *amqp.Channel, exchangeName, routingKey string, data amqp.Publishing) (err error) {
	return ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		data,         // data
	)

}
