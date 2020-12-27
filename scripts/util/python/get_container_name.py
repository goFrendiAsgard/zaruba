from common_helper import get_argv, get_kwargs
from generator_helper import get_container_name

import os, sys, traceback

# USAGE
# python get_container_name.py <image_name>

if __name__ == '__main__':
    image_name = get_argv(1)
    try:
        print(get_container_name(image_name))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)