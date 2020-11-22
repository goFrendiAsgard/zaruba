from typing import Any
import sys
import os

from ruamel.yaml import YAML
import project


# USAGE
# python create_docker_task.py <image> <container> <task>

def create_docker_task(template_path: str, image: str, container: str, task: str) -> str:
    image = image if image != '' else 'nginx'
    container = container if container != '' else image
    task = task if task != '' else 'run{}'.format(container.capitalize())
    create_dir('docker')
    task_filename = get_taskfilename_or_error(task)
    template_filename = get_actual_template_filename(template_path, image)
    should_override_image = get_should_override_image(template_path, image)
    template_obj = get_template_obj(template_filename)
    create_docker_task_file(task_filename, task, container, image, should_override_image, template_obj)
    project.include(task_filename)
    print('Task {} ({}) is created successfully'.format(task, task_filename))


def create_dir(dirname: str):
    if not os.path.exists(dirname):
        os.makedirs(dirname)


def get_taskfilename_or_error(task: str) -> str:
    task_filename = os.path.join('.', 'docker', '{}.zaruba.yaml'.format(task))
    if os.path.isfile(task_filename):
        raise Exception('{} already exists'.format(task_filename))
    return task_filename


def get_should_override_image(template_path: str, image: str) -> bool:
    template_filename = get_template_filename(template_path, image)
    if not os.path.isfile(template_filename):
        return True
    return False


def get_actual_template_filename(template_path: str, image: str) -> str:
    template_filename = get_template_filename(template_path, image)
    if not os.path.isfile(template_filename):
        template_filename = os.path.join(template_path, 'default.zaruba.yaml')
    return template_filename


def get_template_obj(template_filename: str) -> Any:
    yaml=YAML()
    template_obj = yaml.load(open(template_filename, 'r'))
    return template_obj


def get_template_filename(template_path: str, image: str) -> str:
    return os.path.join(template_path, '{}.zaruba.yaml'.format(image))


def create_docker_task_file(task_filename: str, task: str, container: str, image: str, should_override_image: bool, template_obj: Any):
    obj = {'tasks': {}}
    obj['tasks'][task] = template_obj['tasks']['runContainer']
    obj['tasks'][task]['config']['containerName'] = container
    if should_override_image:
        obj['tasks'][task]['config']['imageName'] = image
    yaml=YAML()
    yaml.dump(obj, open(task_filename, 'w'))


if __name__ == '__main__':
    template_path = os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(sys.argv[0]))),
        'docker-task-template'
    )
    image = sys.argv[1] if len(sys.argv) > 1 else ''
    container = sys.argv[2] if len(sys.argv) > 2 else ''
    task = sys.argv[3] if len(sys.argv) > 3 else ''
    create_docker_task(template_path, image, container, task)
