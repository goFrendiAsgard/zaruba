package transport

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func rmqRPCGenerateReplyQueueName(queueName string) (replyQueueName string, err error) {
	randomID, err := uuid.NewUUID()
	if err != nil {
		return queueName, err
	}
	replyQueueName = fmt.Sprintf("%s.reply.%s", queueName, strings.ReplaceAll(randomID.String(), "-", ""))
	return replyQueueName, err
}

func rmqRPCCall(ch *amqp.Channel, exchangeName, replyTo string, inputs []interface{}) (err error) {
	envelopedInput, err := rpcCreateEnvelopedInput(inputs)
	if err != nil {
		return err
	}
	jsonMessage, err := envelopedInput.ToJSON()
	if err != nil {
		return err
	}
	// publish
	return rmqPublish(ch, exchangeName, "",
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
