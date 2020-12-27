from common_helper import get_argv, get_kwargs
from generator_helper import get_task_name

import os, sys, traceback

# USAGE
# python get_task_name.py <service_or_container>

if __name__ == '__main__':
    service_or_container = get_argv(1)
    try:
        print(get_task_name(service_or_container))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)