import re

script_template = '''
    @transport.handle(mb, '{event}')
    def {handler}(msg: Any):
        print('Getting message from {event}', msg)
'''

def get_script(event: str) -> str:
    handler='handle_event_{}'.format(re.sub(r'[^A-Za-z0-9_]+', '_', event).lower())
    return script_template.format(event=event, handler=handler)
