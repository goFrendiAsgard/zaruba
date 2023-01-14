from typing import Any, Callable, Dict, Mapping, TypedDict
from helper.transport.serial import (
    create_json_byte_encoder, create_json_decoder
)


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
        if self._is_key_gt_0_in_event_map(event_name, 'exchange_name'):
            return self.mapping[event_name]['exchange_name']
        return event_name

    def get_queue_name(self, event_name: str) -> str:
        if self._is_key_gt_0_in_event_map(event_name, 'queue_name'):
            return self.mapping[event_name]['queue_name']
        return event_name

    def get_dead_letter_exchange(self, event_name: str) -> str:
        if self._is_key_gt_0_in_event_map(event_name, 'dead_letter_exchange'):
            return self.mapping[event_name]['dead_letter_exchange']
        return '{}.dlx'.format(self.get_exchange_name(event_name))

    def get_dead_letter_queue(self, event_name: str) -> str:
        if self._is_key_gt_0_in_event_map(event_name, 'dead_letter_queue'):
            return self.mapping[event_name]['dead_letter_queue']
        return '{}.dlx'.format(self.get_queue_name(event_name))

    def get_ttl(self, event_name: str) -> int:
        if self._is_key_gt_0_in_event_map(event_name, 'ttl'):
            return self.mapping[event_name]['ttl']
        return 0

    def get_queue_arguments(self, event_name: str) -> Dict:
        args = {}
        if self.get_ttl(event_name) <= 0:
            return {}
        args['x-dead-letter-exchange'] = self.get_dead_letter_exchange(
            event_name
        )
        args['x-message-ttl'] = self.get_ttl(event_name)
        return args

    def get_rpc_timeout(self, event_name: str) -> int:
        if self._is_key_gt_0_in_event_map(event_name, 'rpc_timeout'):
            return self.mapping[event_name]['rpc_timeout']
        return 10000

    def get_prefetch_count(self, event_name: str) -> int:
        if self._is_key_gt_0_in_event_map(event_name, 'prefetch_count'):
            return self.mapping[event_name]['prefetch_count']
        return 10

    def get_auto_ack(self, event_name: str) -> bool:
        if self._is_key_set_in_event_map(event_name, 'auto_ack'):
            return self.mapping[event_name]['auto_ack']
        return False

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

    def _is_key_gt_0_in_event_map(self, event_name: str, key: str) -> bool:
        if not self._is_key_in_event_map(event_name, key):
            return False
        return self.mapping[event_name][key] > 0

    def _is_key_set_in_event_map(self, event_name: str, key: str) -> bool:
        if not self._is_key_in_event_map(event_name, key):
            return False
        return self.mapping[event_name][key] != ''

    def _is_key_in_event_map(self, event_name: str, key: str) -> bool:
        if event_name not in self.mapping:
            return False
        return key in self.mapping[event_name]
