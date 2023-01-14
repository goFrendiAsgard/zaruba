from typing import Any, Callable, Mapping, TypedDict
from helper.transport.serial import (
    create_json_string_encoder, create_json_decoder
)
import datetime


class KafkaEventConfig(TypedDict):
    topic: str
    group_id: str
    encoder: Callable[[Any], bytes]
    decoder: Callable[[bytes], Any]


class KafkaEventMap:

    def __init__(self, mapping: Mapping[str, KafkaEventConfig]):
        self.mapping = mapping
        self.default_encoder = create_json_string_encoder()
        self.default_decoder = create_json_decoder()

    def get_topic(self, event_name: str) -> str:
        if self._is_key_set_in_event_map(event_name, 'topic'):
            return self.mapping[event_name]['topic']
        return event_name

    def get_group_id(self, event_name: str) -> str:
        if self._is_key_set_in_event_map(event_name, 'group_id'):
            return self.mapping[event_name]['group_id']
        return 'default'

    def get_key_maker(self, event_name: str) -> Callable[[Any], Any]:
        if self._is_key_callable_in_event_map(event_name, 'key_maker'):
            return self.mapping[event_name]['key_maker']
        return lambda msg: '{}'.format(datetime.datetime.now().timestamp())

    def get_encoder(self, event_name: str) -> Callable:
        if self._is_key_callable_in_event_map(event_name, 'encoder'):
            return self.mapping[event_name]['encoder']
        return self.default_encoder

    def get_decoder(self, event_name: str) -> Callable:
        if self._is_key_callable_in_event_map(event_name, 'decoder'):
            return self.mapping[event_name]['decoder']
        return self.default_decoder

    def _is_key_callable_in_event_map(self, event_name: str, key: str) -> bool:
        if not self._is_key_set_in_event_map(event_name, key):
            return False
        return callable(self.mapping[event_name][key])

    def _is_key_set_in_event_map(self, event_name: str, key: str) -> bool:
        if not self._is_key_in_event_map(event_name, key):
            return False
        return self.mapping[event_name][key] != ''

    def _is_key_in_event_map(self, event_name: str, key: str) -> bool:
        if event_name not in self.mapping:
            return False
        return key in self.mapping[event_name]
