from helper import cli
import helper.generator as generator
import template.fastapi_route_handler as fastapi_route_handler

import os, re

@cli
def create_fast_route(location: str, module: str, url: str):
    file_name = os.path.abspath(os.path.join(location, module, 'route.py'))
    lines = generator.read_lines(file_name)
    # look for line with 'def init(' prefix
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            insert_index = index + 1
            break
    if insert_index == -1:
        raise Exception('init function not found in {}'.format(file_name))
    lines.insert(insert_index, fastapi_route_handler.get_script(url=url) + '\n')
    generator.write_lines(file_name, lines)


if __name__ == '__main__':
    create_fast_route()