from common_helper import get_argv
from generator_helper import read_config, write_config
from decoration import yellow, normal
from ruamel.yaml import YAML

import sys, traceback

# USAGE
# python add_link.py <source> <destination> [file]

if __name__ == '__main__':
    try:
        source = get_argv(1)
        destination = get_argv(2)
        file_name = get_argv(3, './default.values.yaml')
        config = read_config(file_name)
        print('{yellow}Add link from "{source}" to "{destination}" on "{file_name}"{normal}'.format(yellow=yellow, normal=normal, source=source, destination=destination, file_name=file_name))
        config['link::{}'.format(destination)] = source
        write_config(file_name, config)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)

