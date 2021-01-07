from common_helper import get_argv

import os, re, sys, traceback

# USAGE
# python create_fast_event_handler <location> <module>

handle_event_template = '''
    @transport.handle(mb, '{event}')
    def {handler}(msg: Any):
        print('Getting message from {event}', msg)

'''

def create_fast_event_handler(location: str, module: str, event: str):
    file_name = os.path.abspath(os.path.join(location, module, 'event.py'))
    f_read = open(file_name, 'r')
    lines = f_read.readlines()
    f_read.close()
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
    f_write = open(file_name, 'w')
    f_write.writelines(lines)
    f_write.close()


if __name__ == '__main__':
    location = get_argv(1)
    module = get_argv(2)
    event = get_argv(3)
    try:
        create_fast_event_handler(location, module, event)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
