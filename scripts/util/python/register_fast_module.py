from common_helper import get_argv

import os, sys, traceback

# USAGE
# python register_fast_module <location> <module>

def register_fast_module(location: str, module: str):
    # read main file
    f_read = open(os.path.abspath(os.path.join(location, 'main.py')), 'r')
    lines = f_read.readlines()
    f_read.close()
    # look for last line with 'import' prefix
    import_found = False
    insert_index = 0
    for index, line in enumerate(lines):
        if line.startswith('import '):
            import_found = True
        elif import_found:
            insert_index = index
            break
    # add importer
    lines.insert(insert_index, 'import {}'.format(module))
    # add initiator
    lines.append('# init {}\n'.format(module))
    lines.append('{}.event.init(mb, engine, DBSession)\n'.format(module))
    lines.append('{}.route.init(app, mb)\n'.format(module))
    # rewrite main file
    f_write = open(os.path.abspath(os.path.join(location, 'main.py')), 'w')
    f_write.writelines(lines)
    f_write.close()


if __name__ == '__main__':
    location = get_argv(1)
    module = get_argv(2)
    try:
        register_fast_module(location, module)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
