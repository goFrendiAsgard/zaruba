package transport

import (
	"encoding/json"

	"github.com/google/uuid"
)

// CreateEnvelopedMessageFromJSON create new EnvelopedMessage
func CreateEnvelopedMessageFromJSON(jsonMessage []byte) (em *EnvelopedMessage, err error) {
	em = &EnvelopedMessage{}
	if err = json.Unmarshal(jsonMessage, em); err != nil {
		return em, err
	}
	return em, err
}

// CreateEnvelopedMessage create enveloped message
func CreateEnvelopedMessage() (em *EnvelopedMessage) {
	return &EnvelopedMessage{}
}

// EnvelopedMessage Message structure while transporting
type EnvelopedMessage struct {
	CorrelationID string  `json:"correlation_id"`
	Message       Message `json:"message"`
	ErrorMessage  string  `json:"error"`
}

// SetNewCorrelationID set new correlationID
func (em *EnvelopedMessage) SetNewCorrelationID() (err error) {
	correlationID, err := uuid.NewUUID()
	if err == nil {
		em.CorrelationID = correlationID.String()
	}
	return err
}

// ToJSON change envelopedMessage to JSON
func (em *EnvelopedMessage) ToJSON() (jsonMessage []byte, err error) {
	return json.Marshal(em)
}
