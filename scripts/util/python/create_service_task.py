from typing import List
from common_helper import get_argv
from generator_helper import copy_and_replace_all, get_env_prefix, get_task_file_name, get_service_name, get_task_name, read_config, register_task, update_task_env, write_config, write_task_env
from task import Task

import os, sys, traceback


# USAGE
# python create_service_task.py <templates> <location> <service_type> <task>

def get_service_task_template(template_path_list: List[str], service_type:str) -> str:
    for template_path in template_path_list:
        template = '{template_path}/task/service/{service_type}.zaruba.yaml'.format(template_path=template_path, service_type=service_type)
        if os.path.isfile(template):
            return template
    return get_service_task_template(template_path_list, 'default')


def create_service_task(template_path_list: List[str], location: str, service_type: str, task_name: str, ports: List[str]):
    location = location if location != '' else '.'
    service_type = service_type if service_type != '' else 'default'
    task_name = task_name if task_name != '' else get_task_name(location)
    task_file_name = get_task_file_name(task_name)
    task_location = location if os.path.isabs(location) else os.path.join('..', location)
    service_name = get_service_name(location)
    template = get_service_task_template(template_path_list, service_type)
    copy_and_replace_all(template, task_file_name, {
        'zarubaTaskName': task_name,
        'zarubaServiceName': service_name,
        'zarubaTaskLocation': task_location,
        'ZarubaServiceName': service_name.capitalize(),
        'ZARUBA_ENV_PREFIX': get_env_prefix(location),
    })
    config = read_config(task_file_name)
    task = Task(config['tasks'][task_name])
    update_task_env(task, task_file_name)
    if len(ports) == 0:
        ports = task.get_possible_ports()
    task.add_lconfig_ports(ports)
    write_task_env('.', task)
    config['tasks'][task_name] = task.as_dict()
    write_config(task_file_name, config)
    register_task(task_file_name, task_name)
    print('Task {} is created successfully'.format(task_name))
    

if __name__ == '__main__':
    try:
        template_path_list = get_argv(1).split(':')
        location = get_argv(2)
        service_type = get_argv(3)
        task_name = get_argv(4)
        raw_ports = get_argv(5).split(':')
        ports=[port for port in raw_ports if port != '']
        create_service_task(template_path_list, location, service_type, task_name, ports)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)