from helper import cli
import helper.generator as generator

import os, re


handle_route_template = '''
    @app.get('{url}')
    def {handler}():
        return 'response of {url}'

'''

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
    lines.insert(insert_index, handle_route_template.format(
        url=url,
        handler='handle_route_{}'.format(re.sub(r'[^A-Za-z0-9_]+', '_', url).lower())
    ))
    generator.write_lines(file_name, lines)


if __name__ == '__main__':
    create_fast_route()