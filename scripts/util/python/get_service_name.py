from common_helper import get_argv, get_kwargs
from generator_helper import get_service_name

import os, sys, traceback

# USAGE
# python get_service_name.py <location>

if __name__ == '__main__':
    location = get_argv(1)
    try:
        print(get_service_name(location))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)