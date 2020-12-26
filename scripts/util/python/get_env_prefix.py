from common_helper import get_argv, get_kwargs
from generator_helper import get_env_prefix

import os, sys, traceback

# USAGE
# python copy_and_replace.py <location>

if __name__ == '__main__':
    location = get_argv(1)
    try:
        print(get_env_prefix(location))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)