from common_helper import get_argv

import os, sys, traceback

# USAGE
# python register_fast_module <location> <module>

init_module_template = '''
# init {module}
{module}.model.Base.metadata.create_all(bind=engine)
if enable_event_handler:
    {module}.event.init(mb, DBSession)
if enable_route_handler:
    {module}.route.init(app, mb)

'''

def register_fast_module(location: str, module: str):
    file_name = os.path.abspath(os.path.join(location, 'main.py'))
    f_read = open(file_name, 'r')
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
    lines.insert(insert_index, 'import {}'.format(module))
    lines.append(init_module_template.format(module=module))
    f_write = open(file_name, 'w')
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
