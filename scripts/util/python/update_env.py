from common_helper import get_argv
from generator_helper import read_config, update_task_env, write_config, write_task_env
from decoration import yellow, normal

import os, sys, traceback

# USAGE
# python update_env.py <project_dir>

def update_env(project_dir:str):
    if not os.path.isabs(project_dir):
        project_dir = os.path.abspath(project_dir)
    for root, dir_names, file_names in os.walk(project_dir):
        for file_name in file_names:
            if not file_name.endswith('.zaruba.yaml'):
                continue
            file_name = os.path.join(root, file_name)
            dir_name = os.path.dirname(file_name)
            config = read_config(file_name)
            if 'tasks' not in config:
                continue
            print('{yellow}Update tasks in "{file_name}"{normal}'.format(yellow=yellow, normal=normal, file_name=file_name))
            for task_name, task in config['tasks'].items():
                if dir_name != project_dir:
                    task = update_task_env(task, file_name)
                write_task_env(project_dir, task)
                config['tasks'][task_name] = task.as_dict()
            print('{yellow}Save updated tasks into "{file_name}"{normal}'.format(yellow=yellow, normal=normal, file_name=file_name))
            write_config(file_name, config)


if __name__ == '__main__':
    try:
        project_dir = get_argv(1, '.')
        update_env(project_dir)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)