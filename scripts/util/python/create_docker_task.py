from typing import Any, List
import sys
import re
import traceback

import project


# USAGE
# python create_docker_task.py <image> <container> <task>

def create_docker_task(template_path_list: List[str], image: str, container: str, task_name: str):
    try:
        image = image if image != '' else 'nginx'
        container = container if container != '' else get_default_container_name(image)
        task_name = task_name if task_name != '' else 'run{}'.format(container.capitalize())
        project.create_dir('zaruba-tasks')
        task_file_name = project.get_task_file_name(task_name)
        template_file_name, should_override_image = project.get_docker_task_template(template_path_list, image)
        template_dict = project.get_dict_from_file(template_file_name)
        create_docker_task_file(task_file_name, task_name, container, image, should_override_image, template_dict)
        print('Task {} ({}) is created successfully'.format(task_name, task_file_name))
        project.add_to_main_include(task_file_name)
        project.add_to_main_task(task_name)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


def get_default_container_name(image: str) -> str:
    capitalized_alphanum = re.sub(r'[^A-Za-z0-9]+', ' ', image).capitalize()
    return capitalized_alphanum.replace(' ', '')


def create_docker_task_file(task_file_name: str, task_name: str, container: str, image: str, should_override_image: bool, template_obj: Any):
    task = project.Task(template_obj['tasks']['runContainer'])
    task.set_config('containerName', container)
    if should_override_image:
        task.set_config('imageName', image)
    project_dict = {'tasks': {task_name: task.as_dict()}}
    project.save_dict_to_file(task_file_name, project_dict)
    project.write_task_env('.env', task)
    project.write_task_env('template.env', task)


if __name__ == '__main__':
    template_path_list = sys.argv[1].split(':') if len(sys.argv) > 1 and sys.argv[1] != '' else []
    image = sys.argv[2] if len(sys.argv) > 2 else ''
    container = sys.argv[3] if len(sys.argv) > 3 else ''
    task_name = sys.argv[4] if len(sys.argv) > 4 else ''
    create_docker_task(template_path_list, image, container, task_name)
