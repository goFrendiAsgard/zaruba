package transport

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func rpcCreateEnvelopedInput(inputs []interface{}) (envelopedInput *EnvelopedMessage, err error) {
	msg := Message{"inputs": inputs}
	envelopedInput = CreateEnvelopedMessage()
	if err = envelopedInput.SetNewCorrelationID(); err != nil {
		return envelopedInput, err
	}
	envelopedInput.Message = msg
	return envelopedInput, err
}

func rpcInputsToJSON(inputs []interface{}) (jsonMessage []byte, err error) {
	envelopedInput, err := rpcCreateEnvelopedInput(inputs)
	if err != nil {
		return jsonMessage, err
	}
	return envelopedInput.ToJSON()
}

func rpcCreateEnvelopedError(envelopedInput *EnvelopedMessage, err error) (envelopedError *EnvelopedMessage) {
	errorMessage := fmt.Sprintf("%s", err)
	envelopedError = CreateEnvelopedMessage()
	envelopedError.CorrelationID = envelopedInput.CorrelationID
	envelopedError.Message = Message{"output": "", "error": errorMessage}
	envelopedError.ErrorMessage = errorMessage
	return envelopedError
}

func rpcCreateEnvelopedOutput(envelopedInput *EnvelopedMessage, output interface{}) (envelopedOutput *EnvelopedMessage) {
	envelopedOutput = CreateEnvelopedMessage()
	envelopedOutput.CorrelationID = envelopedInput.CorrelationID
	envelopedOutput.Message = Message{"output": output, "error": ""}
	return envelopedOutput
}

func rmqRPCGenerateReplyQueueName(functionName string) (queueName string, err error) {
	randomID, err := uuid.NewUUID()
	if err != nil {
		return queueName, err
	}
	queueName = fmt.Sprintf("%s.reply.%s", functionName, strings.ReplaceAll(randomID.String(), "-", ""))
	return queueName, err
}

func rmqRPCCall(ch *amqp.Channel, functionName, replyTo string, inputs []interface{}) (err error) {
	envelopedInput, err := rpcCreateEnvelopedInput(inputs)
	if err != nil {
		return err
	}
	jsonMessage, err := envelopedInput.ToJSON()
	if err != nil {
		return err
	}
	// publish
	return rmqPublish(ch, functionName, "",
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: envelopedInput.CorrelationID,
			ReplyTo:       replyTo,
			Body:          jsonMessage,
		})
}

func rmqRPCReplyOutput(ch *amqp.Channel, replyTo string, envelopedInput *EnvelopedMessage, output interface{}) (err error) {
	envelopedOutput := rpcCreateEnvelopedOutput(envelopedInput, output)
	jsonMessage, err := envelopedOutput.ToJSON()
	if err != nil {
		return err
	}
	// reply
	return rmqPublish(ch, "", replyTo,
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: envelopedInput.CorrelationID,
			Body:          jsonMessage,
		})
}

func rmqRPCReplyError(ch *amqp.Channel, replyTo string, envelopedInput *EnvelopedMessage, errReply error) (err error) {
	envelopedErr := rpcCreateEnvelopedError(envelopedInput, errReply)
	jsonMessage, err := envelopedErr.ToJSON()
	if err != nil {
		return err
	}
	// reply
	return rmqPublish(ch, "", replyTo,
		amqp.Publishing{
			ContentType:   "text/json",
			CorrelationId: envelopedInput.CorrelationID,
			Body:          jsonMessage,
		})
}

func rmqDeclareQueueAndBindToDefaultExchange(ch *amqp.Channel, queueName string) (q amqp.Queue, err error) {
	// declare exchange
	err = rmqDeclareFanoutExchange(ch, queueName)
	if err != nil {
		return q, err
	}
	// declare queue
	q, err = rmqDeclareQueue(ch, queueName)
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

func rmqDeclareQueue(ch *amqp.Channel, queueName string) (q amqp.Queue, err error) {
	return ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
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

func rmqConsume(ch *amqp.Channel, queueName string) (rmqMessages <-chan amqp.Delivery, err error) {
	_, err = rmqDeclareQueue(ch, queueName)
	if err != nil {
		return rmqMessages, err
	}
	// start consume
	return ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
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
