from typing import List
from common_helper import get_argv
from generator_helper import copy_and_replace_all, get_env_prefix, get_task_file_name, get_service_name, get_run_task_name, get_build_image_task_name, get_push_image_task_name, read_config, register_run_task, register_build_image_task, register_push_image_task, update_task_env, write_config, write_task_env
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


def create_service_task(template_path_list: List[str], location: str, service_type: str, run_task_name: str, ports: List[str]):
    location = location if location != '' else '.'
    service_type = service_type if service_type != '' else 'default'
    run_task_name = run_task_name if run_task_name != '' else get_run_task_name(location)
    task_file_name = get_task_file_name(run_task_name)
    task_location = location if os.path.isabs(location) else os.path.join('..', location)
    service_name = get_service_name(location)
    build_image_task_name = get_build_image_task_name(service_name)
    push_image_task_name = get_push_image_task_name(service_name)
    template = get_service_task_template(template_path_list, service_type)
    copy_and_replace_all(template, task_file_name, {
        'zarubaRunTask': run_task_name,
        'zarubaBuildImageTask': build_image_task_name,
        'zarubaPushImageTask': push_image_task_name,
        'zarubaServiceName': service_name,
        'zarubaTaskLocation': task_location,
        'ZarubaServiceName': service_name.capitalize(),
        'ZARUBA_ENV_PREFIX': get_env_prefix(location),
    })
    config = read_config(task_file_name)
    task = Task(config['tasks'][run_task_name])
    update_task_env(task, task_file_name)
    if len(ports) == 0:
        ports = task.get_possible_ports()
    task.add_lconfig_ports(ports)
    write_task_env('.', task)
    config['tasks'][run_task_name] = task.as_dict()
    write_config(task_file_name, config)
    register_run_task(task_file_name, run_task_name)
    register_build_image_task(task_file_name, build_image_task_name)
    register_push_image_task(task_file_name, push_image_task_name)
    print('Task {} is created successfully'.format(run_task_name))
    

if __name__ == '__main__':
    try:
        template_path_list = get_argv(1).split(':')
        location = get_argv(2)
        service_type = get_argv(3)
        run_task_name = get_argv(4)
        raw_ports = get_argv(5).split(':')
        ports=[port for port in raw_ports if port != '']
        create_service_task(template_path_list, location, service_type, run_task_name, ports)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)