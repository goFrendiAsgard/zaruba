from typing import Mapping, TypedDict

class RmqEventConfig(TypedDict):
    queue_name: str
    exchange_name: str


class RmqEventMap:

    def __init__(self, mapping: Mapping[str, RmqEventConfig]):
        self.mapping = mapping
    
    def get_exchange_name(self, event_name: str) -> str:
        return self.mapping[event_name]["exchange_name"] if event_name in self.mapping else event_name

    def get_queue_name(self, event_name: str) -> str:
        return self.mapping[event_name]["queue_name"] if event_name in self.mapping else event_name
