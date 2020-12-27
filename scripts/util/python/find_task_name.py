from typing import Mapping
from common_helper import get_argv, get_kwargs
from generator_helper import main_file_name, read_config, write_config
import os, re, sys, traceback

# USAGE
# python find_task_name.py <property=value>

zaruba_home = os.getenv(
    'ZARUBA_HOME', 
    os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))))
)

def is_match(task_dict: Mapping[str, any], query: Mapping[str, str]) -> bool:
    for expected_key, expected_value in query.items():
        keys = expected_key.split('.')
        value = task_dict
        for key in keys:
            if key not in value:
                return False
            value = value[key]
        match = re.search(expected_value, value)
        if not match:
            return False
    return True


def find_task_name(file_name: str, query: Mapping[str, str]):
    config = read_config(file_name)
    if 'tasks' in config:
        for task_name, task_dict in config['tasks'].items():
            if is_match(task_dict, query):
                return task_name
    if 'includes' in config:
        dir_name = os.path.dirname(file_name)
        for included_file in config['includes']:
            included_file = included_file.replace('${ZARUBA_HOME}', zaruba_home)
            if not os.path.isabs(included_file):
                included_file = os.path.abspath(os.path.join(dir_name, included_file))
            task_name = find_task_name(included_file, query)
            if task_name != '':
                return task_name
    return ''


if __name__ == '__main__':
    file_name = get_argv(1, main_file_name)
    query = get_kwargs()
    print(find_task_name(file_name, query))