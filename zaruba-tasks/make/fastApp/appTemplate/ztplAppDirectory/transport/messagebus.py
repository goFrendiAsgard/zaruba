from typing import Any, Callable, List
from helper.transport import MessageBus
from schema.activity import ActivityData


class AppMessageBus(MessageBus):
    '''
    Wrapping for MesssageBus with special methods to support app use case.

    Feel free to add methods as necessary
    '''

    def __init__(self, mb: MessageBus, activity_events: List[str] = []):
        '''
        Initiate a new AppMessageBus.

        Keyword arguments:
        - mb -- MessageBus you want to use.
        - activity_events -- List of event name to be triggered.
            whenever AppMessageBus.publish_activity is performed (default: []).
        '''
        self.mb = mb
        self.activity_events = activity_events

    def handle(self, event_name: str) -> Callable[..., Any]:
        '''
        Decorator to handle an event.

        Keyword arguments:
        - event_name -- Name of the event you want to handle.
        '''
        return self.mb.handle(event_name)

    def publish(self, event_name: str, message: Any) -> Any:
        '''
        Publish an event_name containing a message.

        Keyword arguments:
        - event_name -- Name of the event you want to publish.
        - message -- Message you want to published,
            should be a dictionary, a list, or primitive data,
            cannot contain objects.
        '''
        return self.mb.publish(event_name, message)

    def shutdown(self) -> Any:
        '''
        Shutdown the AppMessageBus.
        '''
        return self.mb.shutdown()

    def get_error_count(self) -> int:
        '''
        Get how many error has been occurred while publish/handle events.
        '''
        return self.mb.get_error_count()

    def is_failing(self) -> bool:
        '''
        Get whether AppMessageBus is failing (and should be terminated) or not.
        '''
        return self.mb.is_failing()

    def broadcast(self, event_names: List[str], message: Any):
        '''
        Publish multiple events containing the same message.

        Keyword arguments:
        - event_names -- List of event names.
        - message -- Message you want to published.
        '''
        for event_name in event_names:
            self.mb.publish(event_name, message)

    def publish_activity(self, activity_data: ActivityData):
        '''
        Publish message to `new_activty` event
        and other specified activity_events.

        Keyword arguments:
        - activity_data -- Instance of ActivityData,
            containing an activity data to be published.
        '''
        self.mb.publish('new_activity', activity_data.dict())
        self.broadcast(self.activity_events, activity_data.dict())
