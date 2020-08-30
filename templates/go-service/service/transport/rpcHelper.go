package transport

import (
	"fmt"
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
