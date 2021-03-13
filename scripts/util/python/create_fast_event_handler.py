from helper import cli
import helper.generator as generator

import os, re


handle_event_template = '''
    @transport.handle(mb, '{event}')
    def {handler}(msg: Any):
        print('Getting message from {event}', msg)

'''

@cli
def create_fast_event_handler(location: str, module: str, event: str):
    file_name = os.path.abspath(os.path.join(location, module, 'event.py'))
    lines = generator.read_lines(file_name)
    # look for line with 'def init(' prefix
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            insert_index = index + 1
            break
    if insert_index == -1:
        raise Exception('init function not found in {}'.format(file_name))
    lines.insert(insert_index, handle_event_template.format(
        event=event,
        handler='handle_event_{}'.format(re.sub(r'[^A-Za-z0-9_]+', '_', event).lower())
    ))
    generator.write_lines(file_name, lines)


if __name__ == '__main__':
    create_fast_event_handler()