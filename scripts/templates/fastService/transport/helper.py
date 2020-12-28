from typing import Any, Callable
from transport.interface import MessageBus


def handle(mb: MessageBus, event_name: str) -> Callable[..., Any]:
    def decorated_handler(handler: Callable[[Any], Any]):
        mb.handle(event_name, handler)
    return decorated_handler


def handle_rpc(mb: MessageBus, event_name: str) -> Callable[..., Any]:
    def decorated_handler(handler: Callable[..., Any]) -> Any:
        mb.handle_rpc(event_name, handler)
    return decorated_handler
