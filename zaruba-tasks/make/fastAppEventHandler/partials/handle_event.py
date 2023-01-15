
@mb.handle('ztplAppEventName')
def handle_ztpl_app_event_name(message: Mapping[str, Any]):
    logging.info('handle event ztplAppEventName with message: {}'.format(message))
