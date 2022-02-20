from typing import Any, Callable, Mapping, TypedDict
from confluent_kafka import avro

import datetime

class KafkaAvroEventConfig(TypedDict):
    topic: str
    group_id: str
    key_schema_str: str
    key_schema_file: str
    value_schema_str: str
    value_schema_file: str
    key_maker: Callable[[Any], Any]


class KafkaAvroEventMap:

    def __init__(self, mapping: Mapping[str, KafkaAvroEventConfig]):
        self.mapping = mapping

    def get_topic(self, event_name: str) -> str:
        if event_name in self.mapping and 'topic' in self.mapping[event_name] and self.mapping[event_name]['topic'] != '':
            return self.mapping[event_name]['topic']
        return event_name

    def get_group_id(self, event_name: str) -> str:
        if event_name in self.mapping and 'group_id' in self.mapping[event_name] and self.mapping[event_name]['group_id'] != '':
            return self.mapping[event_name]['group_id']
        return 'default'
    
    def get_key_schema(self, event_name: str) -> Any:
        if event_name in self.mapping and 'key_schema_str' in self.mapping[event_name] and self.mapping[event_name]['key_schema_str'] != '':
            return avro.loads(self.mapping[event_name]['key_schema_str'])
        if event_name in self.mapping and 'key_schema_file' in self.mapping[event_name] and self.mapping[event_name]['key_schema_file'] != '':
            return avro.load(self.mapping[event_name]['key_schema_file'])
        return avro.loads('{"type": "int"}')
    
    def get_key_maker(self, event_name: str) -> Callable[[Any], Any]:
        if event_name in self.mapping and 'key_maker' in self.mapping[event_name] and callable(self.mapping[event_name]['key_maker']):
            return self.mapping[event_name]['key_maker']
        return lambda msg: datetime.datetime.now().timestamp()
        
    def get_value_schema(self, event_name: str) -> Any:
        if event_name in self.mapping and 'value_schema_str' in self.mapping[event_name] and self.mapping[event_name]['value_schema_str'] != '':
            return avro.loads(self.mapping[event_name]['value_schema_str'])
        if event_name in self.mapping and 'value_schema_file' in self.mapping[event_name] and self.mapping[event_name]['value_schema_file'] != '':
            return avro.load(self.mapping[event_name]['value_schema_file'])
        return avro.load('./avro/{}.avro'.format(event_name))
         