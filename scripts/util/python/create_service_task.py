from typing import List, Mapping
from helper import cli

import helper.generator as generator
import helper.task as task
import template.additional_service_task as additional_service_task

import os, sys, traceback

@cli
def create_service_task(templates: str = '', location: str = '.', service_name: str = '', service_type: str = '', ports: str = '', image_name: str = ''):
    template_path_list = templates.split(':')
    service_ports=[service_port for service_port in ports.split(',') if service_port != '']
    location = location if location != '' else '.'
    service_type = service_type if service_type != '' else 'default'
    service_name = service_name if service_name != '' else generator.get_service_name(location)
    run_task_name = generator.get_run_task_name(service_name)
    build_image_task_name = generator.get_build_image_task_name(service_name)
    push_image_task_name = generator.get_push_image_task_name(service_name)
    run_container_task_name = generator.get_run_container_task_name(service_name)
    remove_container_task_name = generator.get_remove_container_task_name(service_name)
    stop_container_task_name = generator.get_stop_container_task_name(service_name)
    # get task template
    task_template_location = get_service_task_template_location(template_path_list, service_type)
    task_template = generator.read_text(task_template_location) + '\n' + additional_service_task.get_script_template()
    replacement_dict = get_replacement_dict(
        location=location, service_name=service_name, image_name=image_name, run_task_name=run_task_name, build_image_task_name = build_image_task_name, push_image_task_name = push_image_task_name, run_container_task_name=run_container_task_name, remove_container_task_name=remove_container_task_name, stop_container_task_name=stop_container_task_name)
    # write task file
    task_file_content = generator.replace_str(task_template, replacement_dict)
    task_file_name = generator.get_task_file_name(service_name)
    generator.write_text(task_file_name, task_file_content)
    # adjust task port and env
    adjust_task(task_file_name, run_task_name, service_ports)
    # register tasks
    generator.register_run_task(task_file_name, run_task_name)
    generator.register_build_image_task(task_file_name, build_image_task_name)
    generator.register_push_image_task(task_file_name, push_image_task_name)
    generator.register_run_container_task(task_file_name, run_container_task_name)
    generator.register_remove_container_task(task_file_name, remove_container_task_name)
    generator.register_stop_container_task(task_file_name, stop_container_task_name)
    print('Task {} is created successfully'.format(run_task_name))


def get_replacement_dict(location: str, service_name: str, image_name: str, run_task_name: str, build_image_task_name: str, push_image_task_name: str, run_container_task_name: str, remove_container_task_name: str, stop_container_task_name: str) -> Mapping[str, str]:
    container_name = service_name
    image_name = image_name if image_name != '' else container_name.lower()
    task_location = location if os.path.isabs(location) else os.path.join('..', location)
    return {
        'zarubaRunTask': run_task_name,
        'zarubaBuildImageTask': build_image_task_name,
        'zarubaPushImageTask': push_image_task_name,
        'zarubaRunContainerTask': run_container_task_name,
        'zarubaRemoveContainerTask': remove_container_task_name,
        'zarubaStopContainerTask': stop_container_task_name,
        'zarubaServiceName': service_name,
        'zarubaContainerName': container_name,
        'zarubaImageName': image_name,
        'zarubaTaskLocation': task_location,
        'ZarubaServiceName': service_name.capitalize(),
        'ZARUBA_ENV_PREFIX': generator.get_env_prefix(location),
    }


def adjust_task(task_file_name: str, run_task_name: str, service_ports: List[str]):
    config = generator.read_config(task_file_name)
    service_task = task.Task(config['tasks'][run_task_name])
    generator.update_task_env(service_task, task_file_name)
    if len(service_ports) == 0:
        service_ports = service_task.get_possible_ports()
    service_task.add_lconfig_ports(service_ports)
    generator.write_task_env('.', service_task)
    config['tasks'][run_task_name] = service_task.as_dict()
    generator.write_config(task_file_name, config)
    

def get_service_task_template_location(template_path_list: List[str], service_type:str) -> str:
    for template_path in template_path_list:
        template = '{template_path}/task/service/{service_type}.zaruba.yaml'.format(template_path=template_path, service_type=service_type)
        if os.path.isfile(template):
            return template
    return get_service_task_template_location(template_path_list, 'default')


if __name__ == '__main__':
    create_service_task()