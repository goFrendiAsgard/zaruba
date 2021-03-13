from typing import List, Mapping
from helper import cli
import helper.generator as generator
import helper.decoration as decoration
import helper.task as task

import os


@cli
def update_service_task_env(project_dir='.'):
    update_start_service_task_env(project_dir)
    adjust_start_container_task(project_dir)


def update_start_service_task_env(project_dir:str):
    if not os.path.isabs(project_dir):
        project_dir = os.path.abspath(project_dir)
    for task_root, task_dir_names, task_file_names in os.walk(project_dir):
        for task_file_name in task_file_names:
            if not task_file_name.endswith('.zaruba.yaml'):
                continue
            task_file_name = os.path.join(task_root, task_file_name)
            dir_name = os.path.dirname(task_file_name)
            config = generator.read_config(task_file_name)
            if 'tasks' not in config:
                continue
            for task_name, task_dict in config['tasks'].items():
                current_task = task.Task(task_dict)
                if current_task.get_location() == project_dir:
                    continue
                if current_task.get_extend() != 'core.startService':
                    continue
                print('{yellow}Update task {task_name} in {file_name}{normal}'.format(yellow=decoration.yellow, normal=decoration.normal, task_name=task_name, file_name=task_file_name))
                generator.update_task_env(current_task, task_file_name)
                generator.write_task_env(project_dir, current_task)
                config['tasks'][task_name] = current_task.as_dict()
                generator.write_config(task_file_name, config)
                print('{yellow}Save updated task {task_name} in {file_name}{normal}'.format(yellow=decoration.yellow, normal=decoration.normal, task_name=task_name, file_name=task_file_name))
                config = generator.read_config(task_file_name)


def update_helm_values(project_dir: str, runDockerTask: task.Task):
    helm_values_path = os.path.join(project_dir, 'helm-deployments', 'values')
    task_env = runDockerTask.get_all_env()
    for helm_values_root, helm_values_dir_names, helm_values_file_names in os.walk(helm_values_path):
        for helm_values_file_name in helm_values_file_names:
            abs_helm_values_file_name = os.path.join(helm_values_root, helm_values_file_name)
            helm_values = generator.read_config(abs_helm_values_file_name)
            if 'app' not in helm_values:
                continue
            if 'container' not in helm_values['app']:
                continue
            if 'image' not in helm_values['app']['container']:
                continue
            helm_values_image = helm_values['app']['container']['image']
            if helm_values_image != runDockerTask.get_config('helm'):
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
            print('{yellow}Update helm values {abs_helm_values_file_name}{normal}'.format(yellow=decoration.yellow, normal=decoration.normal, abs_helm_values_file_name=abs_helm_values_file_name))
            helm_values['app']['container']['env']
            generator.write_config(abs_helm_values_file_name, helm_values)
            print('{yellow}Save helm values {abs_helm_values_file_name}{normal}'.format(yellow=decoration.yellow, normal=decoration.normal, abs_helm_values_file_name=abs_helm_values_file_name))


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
            config = generator.read_config(task_file_name)
            if 'tasks' not in config:
                continue
            for task_name, task_dict in config['tasks'].items():
                current_task = task.Task(task_dict)
                if current_task.get_location() == project_dir:
                    continue
                if current_task.get_extend() != 'core.startDockerContainer':
                    continue
                update_helm_values(project_dir, current_task)


if __name__ == '__main__':
    update_service_task_env()