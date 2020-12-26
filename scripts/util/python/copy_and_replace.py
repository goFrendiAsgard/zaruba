from common_helper import get_argv, get_kwargs
from generator_helper import copy_and_replace

import os, sys, traceback

# USAGE
# python copy_and_replace.py <source> <destination> [replaceKey=replaceValue...]

if __name__ == '__main__':
    try:
        source = get_argv(1)
        destination = get_argv(2)
        replace = get_kwargs()
        copy_and_replace(source, destination, replace)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)