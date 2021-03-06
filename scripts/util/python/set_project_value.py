from common_helper import get_argv
from generator_helper import read_config, write_config
from decoration import yellow, normal
from ruamel.yaml import YAML

import sys, traceback

# USAGE
# python add_project_kwarg.py <key> <value> [file]

if __name__ == '__main__':
    try:
        key = get_argv(1)
        value = get_argv(2)
        file_name = get_argv(3, './default.values.yaml')
        print(key, value, file_name)
        config = read_config(file_name)
        print('{yellow}Set "{key}" into "{value}" on "{file_name}"{normal}'.format(yellow=yellow, normal=normal, key=key, value=value, file_name=file_name))
        config[key] = value
        write_config(file_name, config)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)

