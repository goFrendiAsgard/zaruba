
@mb.handle('sztplAppEventName')
def handle_zaruba_event_name(message: Mapping[str, Any]):
    print('handle event zarubaEventName with message: {}'.format(message))
