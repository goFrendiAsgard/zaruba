from typing import Any, Callable, Dict, Mapping, TypedDict
from helpers.transport.serial import create_json_byte_encoder, create_json_decoder

class RMQEventConfig(TypedDict):
    queue_name: str
    exchange_name: str
    rpc_timeout: int
    dead_letter_exchange: str
    dead_letter_queue: str
    ttl: int
    auto_ack: bool
    prefetch_count: int
    encoder: Callable[[Any], bytes]
    decoder: Callable[[bytes], Any]


class RMQEventMap:

    def __init__(self, mapping: Mapping[str, RMQEventConfig]):
        self.mapping = mapping
        self.default_encoder = create_json_byte_encoder()
        self.default_decoder = create_json_decoder()

    def get_exchange_name(self, event_name: str) -> str:
        if event_name in self.mapping and 'exchange_name' in self.mapping[event_name] and self.mapping[event_name]['exchange_name'] != '':
            return self.mapping[event_name]['exchange_name']
        return event_name

    def get_queue_name(self, event_name: str) -> str:
        if event_name in self.mapping and 'queue_name' in self.mapping[event_name] and self.mapping[event_name]['queue_name'] != '':
            return self.mapping[event_name]['queue_name']
        return event_name

    def get_dead_letter_exchange(self, event_name: str) -> str:
        if event_name in self.mapping and 'dead_letter_exchange' in self.mapping[event_name] and self.mapping[event_name]['dead_letter_exchange'] != '':
            return self.mapping[event_name]['dead_letter_exchange']
        return '{}.dlx'.format(self.get_exchange_name(event_name))

    def get_dead_letter_queue(self, event_name: str) -> str:
        if event_name in self.mapping and 'dead_letter_queue' in self.mapping[event_name] and self.mapping[event_name]['dead_letter_queue'] != '':
            return self.mapping[event_name]['dead_letter_queue']
        return '{}.dlx'.format(self.get_queue_name(event_name))

    def get_ttl(self, event_name: str) -> int:
        if event_name in self.mapping and 'ttl' in self.mapping[event_name] and self.mapping[event_name]['ttl'] > 0:
            return self.mapping[event_name]['ttl']
        return 0

    def get_queue_arguments(self, event_name: str) -> Dict:
        args = {}
        if self.get_ttl(event_name) <= 0:
            return {}
        args['x-dead-letter-exchange'] = self.get_dead_letter_exchange(event_name)
        args['x-message-ttl'] = self.get_ttl(event_name)
        return args

    def get_rpc_timeout(self, event_name: str) -> int:
        if event_name in self.mapping and 'rpc_timeout' in self.mapping[event_name] and self.mapping[event_name]['rpc_timeout'] > 0:
            return self.mapping[event_name]['rpc_timeout']
        return 10000

    def get_prefetch_count(self, event_name: str) -> int:
        if event_name in self.mapping and 'prefetch_count' in self.mappping[event_name] and self.mapping[event_name]['prefetch_count'] > 0:
            return self.mapping[event_name]['prefetch_count']
        return 10

    def get_auto_ack(self, event_name: str) -> bool:
        if event_name in self.mapping and 'auto_ack' in self.mapping[event_name]:
            return self.mapping[event_name]['auto_ack']
        return False

    def get_encoder(self, event_name: str) -> Callable:
        if event_name in self.mapping and 'encoder' in self.mapping[event_name]:
            return self.mapping[event_name]['encoder']
        return self.default_encoder

    def get_decoder(self, event_name: str) -> Callable:
        if event_name in self.mapping and 'decoder' in self.mapping[event_name]:
            return self.mapping[event_name]['decoder']
        return self.default_decoder
