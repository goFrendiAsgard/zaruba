from typing import List
from common_helper import get_argv
from generator_helper import copy, replace_all, get_container_name, get_task_file_name, get_run_task_name, read_config, register_run_task, register_run_container_task, write_task_env
from task import Task

import os, sys, traceback


# USAGE
# python create_docker_task.py <templates> <image> <container> <task>

def get_docker_task_template(template_path_list: List[str], image_name:str) -> str:
    for template_path in template_path_list:
        template = '{template_path}/task/docker/{image_name}.zaruba.yaml'.format(template_path=template_path, image_name=image_name)
        if os.path.isfile(template):
            return template
    return get_docker_task_template(template_path_list, 'default')
    

def create_docker_task(template_path_list: List[str], image_name: str, container_name: str, run_task_name: str):
    template = get_docker_task_template(template_path_list, image_name)
    image_name = image_name if image_name != '' else 'nginx'
    container_name = container_name if container_name != '' else get_container_name(image_name)
    run_task_name = run_task_name if run_task_name != '' else get_run_task_name(container_name)
    task_file_name = get_task_file_name(run_task_name)
    copy(template, task_file_name)
    replace_all(task_file_name, {
        'zarubaRunTask': run_task_name,
        'zarubaImageName': image_name,
        'zarubaContainerName': container_name,
    })
    config = read_config(task_file_name)
    task = Task(config['tasks'][run_task_name])
    write_task_env('.', task)
    register_run_task(task_file_name, run_task_name)
    register_run_container_task(task_file_name, run_task_name)
    print('Task {} is created successfully'.format(run_task_name))


if __name__ == '__main__':
    try:
        template_path_list = get_argv(1).split(':')
        image_name = get_argv(2)
        container_name = get_argv(3)
        run_task_name = get_argv(4)
        create_docker_task(template_path_list, image_name, container_name, run_task_name)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
