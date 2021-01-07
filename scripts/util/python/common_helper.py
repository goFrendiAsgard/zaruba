import sys

def get_kwargs():
    kwargs = {}
    latest_key = ''
    for arg_index, arg in enumerate(sys.argv):
        if arg_index == 0:
            continue
        pair_parts = arg.split('=')
        if len(pair_parts) > 1:
            key = pair_parts[0]
            value = '='.join(pair_parts[1:])
            kwargs[key] = value
            latest_key = key
            continue
        if latest_key != '' and arg != '':
            kwargs[latest_key] += ' {}'.format(arg)
    return kwargs


def get_argv(index: int, default=''):
    latest_index = 0
    for arg_index, arg in enumerate(sys.argv):
        if '=' in arg and arg_index != 0:
            continue
        if arg == '':
            continue
        if latest_index == index:
            return arg
        latest_index += 1
    return default