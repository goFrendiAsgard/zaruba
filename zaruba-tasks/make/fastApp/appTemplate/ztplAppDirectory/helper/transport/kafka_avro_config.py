from typing import Any, Callable, Mapping, TypedDict
from confluent_kafka import avro
from helper.transport.serial import (
    create_json_string_encoder, create_json_decoder
)

import datetime
import os


class KafkaAvroEventConfig(TypedDict):
    topic: str
    group_id: str
    key_schema_str: str
    key_schema_file: str
    value_schema_str: str
    value_schema_file: str
    key_maker: Callable[[Any], Any]
    encoder: Callable[[Any], Any]
    decoder: Callable[[Any], Any]


VALUE_SCHEMA_TEMPLATE = '''
{
    "namespace": "{namespace}",
    "name": "value",
    "type": "record",
    "fields": [
        {
            "name": "data",
            "type": "string"
        }
    ]
}
'''


KEY_SCHEMA_TEMPLATE = '''
{
    "namespace": "{namespace}",
    "name": "key",
    "type": "record",
    "fields" : [
        {
            "name" : "created_at",
            "type" : "float"
        }
    ]
}
'''


class KafkaAvroEventMap:

    def __init__(self, mapping: Mapping[str, KafkaAvroEventConfig]):
        self.mapping = mapping
        encode = create_json_string_encoder()
        self.default_encoder = lambda msg: {'data': encode(msg)}
        decode = create_json_decoder()
        self.default_decoder = lambda msg: decode(msg['data'])

    def get_topic(self, event_name: str) -> str:
        if self._is_key_set_in_event_map(event_name, 'topic'):
            return self.mapping[event_name]['topic']
        return event_name

    def get_group_id(self, event_name: str) -> str:
        if self._is_key_set_in_event_map(event_name, 'group_id'):
            return self.mapping[event_name]['group_id']
        return 'default'

    def get_key_schema(self, event_name: str) -> Any:
        if self._is_key_set_in_event_map(event_name, 'key_schema_str'):
            return avro.loads(self.mapping[event_name]['key_schema_str'])
        if self._is_key_set_in_event_map(event_name, 'key_schema_file'):
            return avro.load(self.mapping[event_name]['key_schema_file'])
        schema_str = KEY_SCHEMA_TEMPLATE.replace('{namespace}', event_name)
        return avro.loads(schema_str)

    def get_key_maker(self, event_name: str) -> Callable[[Any], Any]:
        if self._is_key_callable_in_event_map(event_name, 'key_maker'):
            return self.mapping[event_name]['key_maker']
        return lambda msg: {"created_at": datetime.datetime.now().timestamp()}

    def get_value_schema(self, event_name: str) -> Any:
        if self._is_key_set_in_event_map(event_name, 'value_schema_str'):
            return avro.loads(self.mapping[event_name]['value_schema_str'])
        if self._is_key_set_in_event_map(event_name, 'value_schema_file'):
            return avro.load(self.mapping[event_name]['value_schema_file'])
        default_schema_file = './avro/{}.avro'.format(event_name)
        if os.path.exists(default_schema_file):
            return avro.load(default_schema_file)
        schema_str = VALUE_SCHEMA_TEMPLATE.replace('{namespace}', event_name)
        return avro.loads(schema_str)

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
