from common_helper import get_argv, get_kwargs
from generator_helper import replace_all
from decoration import yellow, normal

import sys, traceback

# USAGE
# python replace_all.py <source> <destination> [replaceKey=replaceValue...]

if __name__ == '__main__':
    try:
        location = get_argv(1)
        replace = get_kwargs()
        print('{yellow}Replace content of "{location}"{normal}'.format(yellow=yellow, normal=normal, location=location))
        replace_all(location, replace)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)