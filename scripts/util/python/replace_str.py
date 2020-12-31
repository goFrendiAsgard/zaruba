from common_helper import get_argv, get_kwargs
from generator_helper import replace_str
from decoration import yellow, normal

import sys, traceback

# USAGE
# python str_replace_all.py <source> <destination> [replaceKey=replaceValue...]

if __name__ == '__main__':
    try:
        string = get_argv(1)
        replace = get_kwargs()
        print(replace_str(string, replace))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)