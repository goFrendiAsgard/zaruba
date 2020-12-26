import sys

def get_kwargs():
    kwargs = {}
    latest_key = ''
    for index, arg in enumerate(sys.argv):
        if index == 0:
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
    if len(sys.argv) > index and sys.argv[index] != '':
        return sys.argv[index]
    return default