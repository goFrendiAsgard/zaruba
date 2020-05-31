from __future__ import annotations
from .interfaces import Message
from typing import Dict, Type, TypeVar
import uuid
import json


class EnvelopedMessage:

    def __init__(self, json_enveloped_message: str = None):
        self.correlation_id: str = ""
        self.message: Message = {}
        self.error_message: str = ""
        if isinstance(json_enveloped_message, bytes) or isinstance(json_enveloped_message, str):
            obj = json.loads(json_enveloped_message)
            self.correlation_id = obj["correlation_id"]
            self.message = obj["message"]
            self.error_message = obj["error_message"]

    def set_correlation_id(self, correlation_id: str = None) -> EnvelopedMessage:
        self.correlation_id = correlation_id if isinstance(
            correlation_id, str) else str(uuid.uuid4())
        return self

    def set_message(self, message: Message) -> EnvelopedMessage:
        self.message = message
        return self

    def to_json(self) -> str:
        data = {
            "correlation_id": self.correlation_id,
            "message": self.message,
            "error_message": self.error_message
        }
        return json.dumps(data)
