from typing import Mapping, TypedDict

class RmqEventConfig(TypedDict):
    queue_name: str
    exchange_name: str
    rpc_timeout: int
    dead_letter_exchange: str
    dead_letter_queue: str
    ttl: int
    auto_ack: bool


class RmqEventMap:

    def __init__(self, mapping: Mapping[str, RmqEventConfig]):
        self.mapping = mapping
    
    def get_exchange_name(self, event_name: str) -> str:
        if event_name in self.mapping and "exchange_name" in self.mapping[event_name] and self.mapping[event_name]["exchange_name"] != "":
            return self.mapping[event_name]["exchange_name"]
        return event_name

    def get_queue_name(self, event_name: str) -> str:
        if event_name in self.mapping and "queue_name" in self.mapping[event_name] and self.mapping[event_name]["queue_name"] != "":
            return self.mapping[event_name]["queue_name"]
        return event_name
 
    def get_dead_letter_exchange(self, event_name: str) -> str:
        if event_name in self.mapping and "dead_letter_exchange" in self.mapping[event_name] and self.mapping[event_name]["dead_letter_exchange"] != "":
            return self.mapping[event_name]["dead_letter_exchange"]
        return "{}.dlx".format(self.get_exchange_name(event_name))
 
    def get_dead_letter_queue(self, event_name: str) -> str:
        if event_name in self.mapping and "dead_letter_queue" in self.mapping[event_name] and self.mapping[event_name]["dead_letter_queue"] != "":
            return self.mapping[event_name]["dead_letter_queue"]
        return "{}.dlx".format(self.get_queue_name(event_name))