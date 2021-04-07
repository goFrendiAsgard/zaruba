from helper import cli
import helper.generator as generator
import template.fastapi_event_handler as fastapi_event_handler

import os

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
    lines.insert(insert_index, fastapi_event_handler.get_script(event=event) + '\n')
    generator.write_lines(file_name, lines)


if __name__ == '__main__':
    create_fast_event_handler()