from typing import List, Mapping
from common_helper import get_argv
from generator_helper import read_config, update_task_env, write_config, write_task_env
from task import Task
from decoration import yellow, normal

import os, sys, traceback

# USAGE
# python update_env.py <project_dir>


def update_start_service_task_env(project_dir:str):
    if not os.path.isabs(project_dir):
        project_dir = os.path.abspath(project_dir)
    for task_root, task_dir_names, task_file_names in os.walk(project_dir):
        for task_file_name in task_file_names:
            if not task_file_name.endswith('.zaruba.yaml'):
                continue
            task_file_name = os.path.join(task_root, task_file_name)
            dir_name = os.path.dirname(task_file_name)
            config = read_config(task_file_name)
            if 'tasks' not in config:
                continue
            for task_name, task_dict in config['tasks'].items():
                task = Task(task_dict)
                if task.get_location() == project_dir:
                    continue
                if task.get_extend() != 'core.startService':
                    continue
                print('{yellow}Update task {task_name} in {file_name}{normal}'.format(yellow=yellow, normal=normal, task_name=task_name, file_name=task_file_name))
                write_task_env(project_dir, task)
                config['tasks'][task_name] = task.as_dict()
                write_config(task_file_name, config)
                print('{yellow}Save updated task {task_name} in {file_name}{normal}'.format(yellow=yellow, normal=normal, task_name=task_name, file_name=task_file_name))
                config = read_config(task_file_name)


def update_helm_values(project_dir: str, task: Task):
    helm_values_path = os.path.join(project_dir, 'helm-deployments', 'values')
    task_env = task.get_all_env()
    for helm_values_root, helm_values_dir_names, helm_values_file_names in os.walk(helm_values_path):
        for helm_values_file_name in helm_values_file_names:
            abs_helm_values_file_name = os.path.join(helm_values_root, helm_values_file_name)
            helm_values = read_config(abs_helm_values_file_name)
            if 'app' not in helm_values:
                continue
            if 'container' not in helm_values['app']:
                continue
            if 'image' not in helm_values['app']['container']:
                continue
            helm_values_image = helm_values['app']['container']['image']
            if helm_values_image != task.get_config('helm'):
                continue
            helm_env_list: List[Mapping[str, str]] = helm_values['app']['container']['env'] if 'env' in helm_values['app']['container'] else []
            for env_key, env in task_env.items():
                env_exists = False
                for helm_env in helm_env_list:
                    if helm_env['name'] == env:
                        env_exists = True
                    pass
                if not env_exists:
                    helm_env_list.append({'name': env_key, 'value': env.get_default()})
            print('{yellow}Update helm values {abs_helm_values_file_name}{normal}'.format(yellow=yellow, normal=normal, abs_helm_values_file_name=abs_helm_values_file_name))
            helm_values['app']['container']['env']
            write_config(abs_helm_values_file_name, helm_values)
            print('{yellow}Save helm values {abs_helm_values_file_name}{normal}'.format(yellow=yellow, normal=normal, abs_helm_values_file_name=abs_helm_values_file_name))


def adjust_start_container_task(project_dir:str):
    if not os.path.isabs(project_dir):
        project_dir = os.path.abspath(project_dir)
    helm_values_path = os.path.join(project_dir, 'helm-deployments', 'values')
    for task_root, task_dir_names, task_file_names in os.walk(project_dir):
        for task_file_name in task_file_names:
            if not task_file_name.endswith('.zaruba.yaml'):
                continue
            task_file_name = os.path.join(task_root, task_file_name)
            dir_name = os.path.dirname(task_file_name)
            config = read_config(task_file_name)
            if 'tasks' not in config:
                continue
            for task_name, task_dict in config['tasks'].items():
                task = Task(task_dict)
                if task.get_location() == project_dir:
                    continue
                if task.get_extend() != 'core.startDockerContainer':
                    continue
                update_helm_values(project_dir, task)


if __name__ == '__main__':
    try:
        project_dir = get_argv(1, '.')
        update_start_service_task_env(project_dir)
        adjust_start_container_task(project_dir)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)