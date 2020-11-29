from typing import Any
import sys
import os
import re
import traceback

import project


# USAGE
# python create_docker_task.py <image> <container> <task>

def create_docker_task(template_path: str, image: str, container: str, task_name: str):
    try:
        image = image if image != '' else 'nginx'
        container = container if container != '' else get_default_container_name(image)
        task_name = task_name if task_name != '' else 'run{}'.format(container.capitalize())
        project.create_dir('tasks')
        task_file_name = get_taskfile_name_or_error(task_name)
        template_file_name = get_template_or_default_file_name(template_path, image)
        should_override_image = get_should_override_image(template_path, image)
        template_dict = project.get_dict(template_file_name)
        create_docker_task_file(task_file_name, task_name, container, image, should_override_image, template_dict)
        print('Task {} ({}) is created successfully'.format(task_name, task_file_name))
        add_to_include(task_file_name)
        add_to_run_task(task_name)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


def get_default_container_name(image: str) -> str:
    capitalized_alphanum = re.sub(r'[^A-Za-z0-9]+', ' ', image).capitalize()
    return capitalized_alphanum.replace(' ', '')


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


def get_taskfile_name_or_error(task_name: str) -> str:
    task_file_name = os.path.join('.', 'tasks', '{}.zaruba.yaml'.format(task_name))
    if os.path.isfile(task_file_name):
        raise Exception('{} already exists'.format(task_file_name))
    return task_file_name


def get_should_override_image(template_path: str, image: str) -> bool:
    template_file_name = get_template_file_name(template_path, image)
    if not os.path.isfile(template_file_name):
        return True
    return False


def get_template_or_default_file_name(template_path: str, image: str) -> str:
    template_file_name = get_template_file_name(template_path, image)
    if not os.path.isfile(template_file_name):
        template_file_name = os.path.join(template_path, 'default.zaruba.yaml')
    return template_file_name


def get_template_file_name(template_path: str, image: str) -> str:
    return os.path.join(template_path, '{}.zaruba.yaml'.format(image))


def create_docker_task_file(task_file_name: str, task_name: str, container: str, image: str, should_override_image: bool, template_obj: Any):
    task_dict = {'tasks': {}}
    task_dict['tasks'][task_name] = template_obj['tasks']['runContainer']
    task_dict['tasks'][task_name]['config']['containerName'] = container
    if should_override_image:
        task_dict['tasks'][task_name]['config']['imageName'] = image
    project.write_dict(task_file_name, task_dict)


if __name__ == '__main__':
    template_path = os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(sys.argv[0]))),
        'docker-task-template'
    )
    image = sys.argv[1] if len(sys.argv) > 1 else ''
    container = sys.argv[2] if len(sys.argv) > 2 else ''
    task_name = sys.argv[3] if len(sys.argv) > 3 else ''
    create_docker_task(template_path, image, container, task_name)
