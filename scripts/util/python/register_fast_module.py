from helper import cli
import helper.generator as generator

import os

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

@cli
def register_fast_module(location: str, module: str):
    file_name = os.path.abspath(os.path.join(location, 'main.py'))
    lines = generator.read_lines(file_name)
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
    generator.write_lines(file_name, lines)


if __name__ == '__main__':
    register_fast_module()