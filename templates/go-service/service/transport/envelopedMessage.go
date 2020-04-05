package transport

import (
	"encoding/json"

	"github.com/google/uuid"
)

// NewEnvelopedMessage create new EnvelopedMessage
func NewEnvelopedMessage(message Message) (em *EnvelopedMessage, err error) {
	correlationID, err := uuid.NewUUID()
	if err != nil {
		return em, err
	}
	em = NewEnvelopedMessageWithCorrelationID(correlationID.String(), message)
	return em, err
}

// NewEnvelopedMessageWithCorrelationID create new EnvelopedMessage
func NewEnvelopedMessageWithCorrelationID(correlationID string, message Message) (em *EnvelopedMessage) {
	em = &EnvelopedMessage{
		CorrelationID: correlationID,
		Message:       message,
	}
	return em
}

// NewEnvelopedMessageFromJSON create new EnvelopedMessage
func NewEnvelopedMessageFromJSON(jsonMessage []byte) (em *EnvelopedMessage, err error) {
	em = &EnvelopedMessage{}
	if err = json.Unmarshal(jsonMessage, em); err != nil {
		return em, err
	}
	return em, err
}

// EnvelopedMessage Message structure while transporting
type EnvelopedMessage struct {
	CorrelationID string  `json:"correlation_id"`
	Message       Message `json:"message"`
	ErrorMessage  string  `json:"error"`
}

// ToJSON change envelopedMessage to JSON
func (em *EnvelopedMessage) ToJSON() (jsonMessage []byte, err error) {
	return json.Marshal(em)
}
