from typing import Mapping, TypedDict, Dict

class RMQEventConfig(TypedDict):
    queue_name: str
    exchange_name: str
    rpc_timeout: int
    dead_letter_exchange: str
    dead_letter_queue: str
    ttl: int
    auto_ack: bool


class RmqEventMap:

    def __init__(self, mapping: Mapping[str, RMQEventConfig]):
        self.mapping = mapping


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
        return 30000


    def get_auto_ack(self, event_name: str) -> bool:
        if event_name in self.mapping and 'auto_ack' in self.mapping[event_name] and self.mapping[event_name]['auto_ack'] > 0:
            return self.mapping[event_name]['auto_ack']
        return False
