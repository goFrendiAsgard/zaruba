from typing import List, Any
from .envelopedMessage import EnvelopedMessage

def rpc_create_enveloped_input(inputs: List[Any]) -> EnvelopedMessage:
    enveloped_message = EnvelopedMessage()
    enveloped_message.message = {"inputs": inputs}
    return enveloped_message


def rpc_inputs_to_json(inputs: List[Any]) -> str:
    return rpc_create_enveloped_input(inputs).to_json()


def rpc_create_enveloped_error(enveloped_input: EnvelopedMessage, err: Exception) -> EnvelopedMessage:
    enveloped_error = EnvelopedMessage().set_correlation_id(
        enveloped_input.correlation_id)
    error_message: str = str(Exception)
    enveloped_error.message = {"output": "", "error": error_message}
    enveloped_error.error_message = error_message
    return enveloped_error


def rpc_create_enveloped_output(enveloped_input: EnvelopedMessage, output: Any) -> EnvelopedMessage:
    enveloped_output = EnvelopedMessage().set_correlation_id(
        enveloped_input.correlation_id)
    enveloped_output.message = {"output": output, "error": ""}
    return enveloped_output

