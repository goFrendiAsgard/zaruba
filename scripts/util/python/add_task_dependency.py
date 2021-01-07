from common_helper import get_argv, get_kwargs
from generator_helper import read_config, write_config
from task import Task

import sys, traceback

# USAGE
# python add_task_dependency.py <file_name> <task_name> <dependency_name>


def add_task_dependency(file_name: str, task_name: str, dependency_name: str):
    config = read_config(file_name)
    if ('tasks' not in config) or (task_name not in config['tasks']):
        raise Exception('There is no task "{}" in "{}"'.format(task_name, file_name))
    task = Task(config['tasks'][task_name])
    task.add_dependency(dependency_name)
    config['tasks'][task_name] = task.as_dict()
    write_config(file_name, config)


if __name__ == '__main__':
    file_name = get_argv(1)
    task_name = get_argv(2)
    dependency_name = get_argv(3)
    try:
       add_task_dependency(file_name, task_name, dependency_name)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)