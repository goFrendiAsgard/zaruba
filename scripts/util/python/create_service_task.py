from typing import Any, Mapping
import sys
import os
import re
import traceback
from dotenv import dotenv_values

import project


# USAGE
# python create_docker_task.py <image> <container> <task>

def create_service_task(template_path: str, location: str, service_type: str, task_name: str):
    try:
        service_type = service_type if service_type != '' else 'default'
        location = location if location != '' else './'
        task_name = task_name if task_name != '' else get_default_task_name(location)
        project.create_dir('zaruba-tasks')
        task_file_name = get_taskfile_name_or_error(task_name)
        template_file_name = get_template_or_default_file_name(template_path, service_type)
        template_dict = project.get_dict(template_file_name)
        create_service_task_file(task_file_name, task_name, location, template_dict)
        print('Task {} ({}) is created successfully'.format(task_name, task_file_name))
        add_to_include(task_file_name)
        add_to_run_task(task_name)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


def get_default_task_name(location: str) -> str:
    abs_location = os.path.abspath(location)
    basename = os.path.basename(abs_location)
    capitalized_alphanum = re.sub(r'[^A-Za-z0-9]+', ' ', basename).capitalize()
    return 'run{}'.format(capitalized_alphanum.replace(' ', ''))


def add_to_include(task_file_name: str):
    if project.add_to_include(task_file_name):
        print('{} has been added to main.zaruba.yaml'.format(task_file_name))
        return
    print('{} already exists in main.zaruba.yaml'.format(task_file_name))


def add_to_run_task(task_name: str):
    if project.add_to_run_task(task_name):
        print('{} has been added to main.zaruba.yaml'.format(task_name))
        return
    print('{} already exists in main.zaruba.yaml'.format(task_name))


def get_taskfile_name_or_error(task: str) -> str:
    task_file_name = os.path.join('.', 'zaruba-tasks', '{}.zaruba.yaml'.format(task))
    if os.path.isfile(task_file_name):
        raise Exception('{} already exists'.format(task_file_name))
    return task_file_name


def get_template_or_default_file_name(template_path: str, image: str) -> str:
    template_file_name = get_template_file_name(template_path, image)
    if not os.path.isfile(template_file_name):
        template_file_name = os.path.join(template_path, 'default.zaruba.yaml')
    return template_file_name


def get_template_file_name(template_path: str, image: str) -> str:
    return os.path.join(template_path, '{}.zaruba.yaml'.format(image))


def create_service_task_file(task_file_name: str, task_name: str, location: str, template_obj: Any):
    task_dict = {'tasks': {}}
    task_dict['tasks'][task_name] = template_obj['tasks']['runService']
    if os.path.isabs(location):
        task_dict['tasks'][task_name]['location'] = location
    else:
        task_dict['tasks'][task_name]['location'] = os.path.join('..', location)
    for env_file in ('.env', 'template.env', 'env.template'):
        env_path = os.path.join(location, env_file)
        if not os.path.isfile(env_path):
            continue
        env_dict: Mapping[str, str] = dotenv_values(env_path)
        for key, val in env_dict.items():
            if 'env' not in task_dict['tasks'][task_name]:
                task_dict['tasks'][task_name]['env'] = {}
            task_dict['tasks'][task_name]['env'][key] = {'default': val}
    project.write_dict(task_file_name, task_dict)


if __name__ == '__main__':
    template_path = os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(sys.argv[0]))),
        'service-task-template'
    )
    location = sys.argv[1] if len(sys.argv) > 1 else ''
    service_type = sys.argv[2] if len(sys.argv) > 2 else ''
    task_name = sys.argv[3] if len(sys.argv) > 3 else ''
    create_service_task(template_path, location, service_type, task_name)
