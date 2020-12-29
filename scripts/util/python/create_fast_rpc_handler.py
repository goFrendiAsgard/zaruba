from common_helper import get_argv

import os, re, sys, traceback

# USAGE
# python create_fast_rpc_handler <location> <module>

def create_fast_rpc_handler(location: str, module: str, event: str):
    file_name = os.path.abspath(os.path.join(location, module, 'event.py'))
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
        '    @transport.handle_rpc(mb, \'{event}\')'.format(event=event),
        '    def handle_{handle_name}(msg: Any) -> Any:'.format(handle_name=re.sub(r'[^A-Za-z0-9_]+', '_', event).lower()),
        '        print(\'Getting message from {event}\', msg)'.format(event=event),
        '        return \'ok\'',
    ]) + '\n\n')
    # rewrite main file
    f_write = open(file_name, 'w')
    f_write.writelines(lines)
    f_write.close()


if __name__ == '__main__':
    location = get_argv(1)
    module = get_argv(2)
    event = get_argv(3)
    try:
        create_fast_rpc_handler(location, module, event)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
