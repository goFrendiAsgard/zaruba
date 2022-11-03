from typing import Any, Callable, List
from helpers.transport import MessageBus
from schemas.activity import ActivityData

class AppMessageBus(MessageBus):
    '''
    MesssageBus with special methods to support app use case.
    Feel free to add methods as necessary
    '''

    def __init__(self, mb: MessageBus, activity_events: List[str] = []):
        self.mb = mb
        self.activity_events = activity_events


    def handle(self, event_name: str) -> Callable[..., Any]:
        return self.mb.handle(event_name)


    def publish(self, event_name: str, message: Any) -> Any:
        return self.mb.publish(event_name, message)


    def shutdown(self) -> Any:
        return self.mb.shutdown()


    def get_error_count(self) -> int:
        return self.mb.get_error_count()


    def is_failing(self) -> bool:
        return self.mb.is_failing()


    def publish_activity(self, activity_data: ActivityData):
        self.mb.publish('new_activity', activity_data.dict())
        self.mb.publish('new_activity', activity_data.dict())
        for activity_event in self.activity_events:
            self.mb.publish(activity_event, activity_data.dict())
