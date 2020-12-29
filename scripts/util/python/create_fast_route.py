from common_helper import get_argv

import os, re, sys, traceback

# USAGE
# python create_fast_route <location> <module>

def create_fast_route(location: str, module: str, url: str):
    file_name = os.path.abspath(os.path.join(location, module, 'route.py'))
    # read main file
    f_read = open(file_name, 'r')
    lines = f_read.readlines()
    f_read.close()
    # look for last line with 'import' prefix
    function_found = False
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            function_found = True
        elif function_found and line.startswith('    '):
            insert_index = index
            break
    # add route handler
    lines.insert(insert_index, '\n' + '\n'.join([
        '    @app.get(\'{url}\')'.format(url=url),
        '    def handle_{handle_name}():'.format(handle_name=re.sub(r'[^A-Za-z0-9_]+', '_', url).lower()),
        '        return \'response of {url}\''.format(url=url)
    ]) + '\n\n')
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
