from typing import Any, Mapping, List
import sys
import os
import re
import traceback
from dotenv import dotenv_values

import project


# USAGE
# python create_docker_task.py <image> <container> <task>

def create_service_task(template_path_list: List[str], task_location: str, service_type: str, task_name: str, ports=List[str]):
    try:
        service_type = service_type if service_type != '' else 'default'
        task_location = task_location if task_location != '' else './'
        task_name = task_name if task_name != '' else get_default_task_name(task_location)
        project.create_dir('zaruba-tasks')
        task_file_name = get_taskfile_name_or_error(task_name)
        template_file_name, _ = project.get_service_task_template(template_path_list, service_type)
        template_dict = project.get_dict_from_file(template_file_name)
        create_service_task_file(task_file_name, task_name, task_location, template_dict, ports)
        print('Task {} ({}) is created successfully'.format(task_name, task_file_name))
        project.add_to_main_include(task_file_name)
        project.add_to_main_task(task_name)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


def get_taskfile_name_or_error(task: str) -> str:
    task_file_name = os.path.join('.', 'zaruba-tasks', '{}.zaruba.yaml'.format(task))
    if os.path.isfile(task_file_name):
        raise Exception('{} already exists'.format(task_file_name))
    return task_file_name


def create_service_task_file(task_file_name: str, task_name: str, task_location: str, template_obj: Any, ports: List[str]):
    task = project.Task(template_obj['tasks']['runService'])
    if os.path.isabs(task_location):
        task.set_location(task_location)
    else:
        task.set_location(os.path.join('..', task_location))
    # adjust env from template
    for key, env in task.get_all_env().items():
        envvar = env.get_from() if env.get_from() else key
        env.set_from(get_env_from(task_location, envvar))
        task.set_env(key, env)
    # add env from location's file
    for env_file in ('sample.env', 'template.env', 'env.template', '.env'):
        env_path = os.path.join(task_location, env_file)
        if not os.path.isfile(env_path):
            continue
        raw_env_dict: Mapping[str, str] = dotenv_values(env_path)
        for key, default in raw_env_dict.items():
            env = project.Env({
                'from': get_env_from(task_location, key),
                'default': default,
            })
            task.set_env(key, env)
    # add lconfig.ports
    ports = ports if ports else task.get_possible_ports()
    task.add_lconfig_ports(*ports)
    # save project and env
    project_dict = {'tasks': {task_name: task.as_dict()}}
    project.save_dict_to_file(task_file_name, project_dict)
    project.write_task_env('.env', task)
    project.write_task_env('template.env', task)


def get_env_from(location: str, env_name: str) -> str:
    upper_env_name = env_name.upper()
    env_prefix = get_env_prefix(location)
    if upper_env_name.startswith(env_prefix):
        return upper_env_name
    return '_'.join([env_prefix, upper_env_name])


def get_env_prefix(location: str) -> str:
    upper_alphanum = get_location_base_name(location).upper()
    return upper_alphanum.replace(' ', '_')


def get_default_task_name(location: str) -> str:
    capitalized_alphanum = get_location_base_name(location).capitalize()
    return 'run{}'.format(capitalized_alphanum.replace(' ', ''))


def get_location_base_name(location: str) -> str:
    abs_location = os.path.abspath(location)
    base_name = os.path.basename(abs_location)
    return re.sub(r'[^A-Za-z0-9]+', ' ', base_name)


if __name__ == '__main__':
    template_path_list = sys.argv[1].split(':') if len(sys.argv) > 1 and sys.argv[1] != '' else []
    location = sys.argv[2] if len(sys.argv) > 2 else ''
    service_type = sys.argv[3] if len(sys.argv) > 3 else ''
    task_name = sys.argv[4] if len(sys.argv) > 4 else ''
    ports = sys.argv[5].split(',') if len(sys.argv) > 5 and sys.argv[5] != '' else []
    create_service_task(template_path_list, location, service_type, task_name, ports)
