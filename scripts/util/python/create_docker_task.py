from typing import List, Mapping
from helper import cli
import helper.generator as generator
import helper.decoration as decoration
import helper.task as task
import template.additional_docker_task as additional_docker_task
import os


@cli
def create_docker_task(templates: str = '', image_name: str = '', image_prefix: str = '', image_tag = '', container_name: str = '', service_name: str = ''):
    template_path_list = templates.split(':')
    image_name = image_name if image_name != '' else 'nginx'
    container_name = container_name if container_name != '' else generator.get_container_name(image_name)
    service_name = service_name if service_name != '' else container_name
    run_task_name = generator.get_run_task_name(container_name)
    stop_task_name = generator.get_stop_task_name(container_name)
    # get task template
    task_template_location = get_docker_task_template_location(template_path_list, image_name)
    task_template = generator.read_text(task_template_location) + '\n' + additional_docker_task.get_script_template()
    replacement_dict = get_replacement_dict(
        service_name=service_name, image_name=image_name, container_name=container_name, run_task_name=run_task_name, stop_task_name=stop_task_name)
    # write task file
    task_file_content = generator.replace_str(task_template, replacement_dict)
    task_file_name = generator.get_task_file_name(service_name)
    generator.write_text(task_file_name, task_file_content)
    # adjust task env
    adjust_task(task_file_name, run_task_name, image_tag, image_prefix)
    # register tasks
    generator.register_run_task(task_file_name, run_task_name)
    generator.register_run_container_task(task_file_name, run_task_name)
    generator.register_stop_container_task(task_file_name, stop_task_name)
    print('Task {} is created successfully'.format(run_task_name))


def get_replacement_dict(service_name: str, image_name: str, container_name: str, run_task_name: str, stop_task_name: str) -> Mapping[str, str]:
    return {
        'zarubaRunTask': run_task_name,
        'zarubaStopContainerTask': stop_task_name,
        'zarubaImageName': image_name,
        'zarubaContainerName': container_name,
        'zarubaServiceName': service_name
    }


def adjust_task(task_file_name: str, run_task_name: str, image_tag: str, image_prefix: str):
    config = generator.read_config(task_file_name)
    run_docker_task = task.Task(config['tasks'][run_task_name])
    if image_tag != '' or image_prefix != '':
        if image_tag != '':
            run_docker_task.set_config('imageTag', image_tag)
        if image_prefix != '':
            run_docker_task.set_config('imagePrefix', image_prefix)
        config['tasks'][run_task_name] = run_docker_task
        generator.write_config(task_file_name, config)
    generator.write_task_env('.', run_docker_task)


def get_docker_task_template_location(template_path_list: List[str], image_name:str) -> str:
    for template_path in template_path_list:
        template = '{template_path}/task/docker/{image_name}.zaruba.yaml'.format(template_path=template_path, image_name=image_name)
        if os.path.isfile(template):
            return template
    return get_docker_task_template_location(template_path_list, 'default')
    

if __name__ == '__main__':
    create_docker_task()