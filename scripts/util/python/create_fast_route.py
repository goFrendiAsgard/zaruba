from common_helper import get_argv

import os, re, sys, traceback

# USAGE
# python create_fast_route <location> <module>


handle_route_template = '''
    @app.get('{url}')
    def {handler}():
        return 'response of {url}'

'''

def create_fast_route(location: str, module: str, url: str):
    file_name = os.path.abspath(os.path.join(location, module, 'route.py'))
    # read main file
    f_read = open(file_name, 'r')
    lines = f_read.readlines()
    f_read.close()
    # look for last line with 'def init(' prefix
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            insert_index = index + 1
            break
    if insert_index == -1:
        raise Exception('init function not found in {}'.format(file_name))
    # add route handler
    lines.insert(insert_index, handle_route_template.format(
        url=url,
        handler='handle_route_{}'.format(re.sub(r'[^A-Za-z0-9_]+', '_', url).lower())
    ))
    # rewrite main file
    f_write = open(file_name, 'w')
    f_write.writelines(lines)
    f_write.close()


if __name__ == '__main__':
    location = get_argv(1)
    module = get_argv(2)
    url = get_argv(3)
    try:
        create_fast_route(location, module, url)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
