from typing import Any, Callable
import abc


class MessageBus(abc.ABC):

    @abc.abstractmethod
    def handle(self, event_name: str) -> Callable[..., Any]:
        '''
        Decorator to handle an event.

        Keyword arguments:
        - event_name -- Name of the event you want to handle.
        '''
        pass

    @abc.abstractmethod
    def publish(self, event_name: str, message: Any) -> Any:
        '''
        Publish an event_name containing a message.

        Keyword arguments:
        - event_name -- Name of the event you want to publish.
        - message -- Message you want to published,
            should be a dictionary, a list, or primitive data,
            cannot contain objects.
        '''
        pass

    @abc.abstractmethod
    def shutdown(self) -> Any:
        '''
        Shutdown the MessageBus
        '''
        pass

    @abc.abstractclassmethod
    def get_error_count(self) -> int:
        '''
        Get how many error has been occurred while publish/handle events.
        '''
        pass

    @abc.abstractclassmethod
    def is_failing(self) -> bool:
        '''
        Get whether MessageBus is failing (and should be terminated) or not.
        '''
        pass
