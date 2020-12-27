from common_helper import get_argv, get_kwargs
from generator_helper import get_task_file_name

import os, sys, traceback

# USAGE
# python get_task_file_name.py <task_name>

if __name__ == '__main__':
    task_name = get_argv(1)
    try:
        print(get_task_file_name(task_name))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)