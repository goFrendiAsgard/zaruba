from typing import Mapping, Any, List
from dotenv import dotenv_values
from task import Task
from ruamel.yaml import YAML

import os, re, shutil, sys, traceback

main_file_name = 'main.zaruba.yaml'

def read_config(file_name: str) -> Mapping[str, Any]:
    yaml=YAML()
    f = open(file_name, 'r')
    template_obj = yaml.load(f)
    f.close()
    return template_obj


def write_config(file_name: str, dictionary: Mapping[str, Any]):
    yaml=YAML()
    f = open(file_name, 'w')
    yaml.dump(dictionary, f)
    f.close()


def register_task(file_name: str, task_name: str) -> bool:
    main_config = read_config(main_file_name)
    if 'includes' not in main_config:
        main_config['includes'] = []
    if file_name not in main_config['includes']:
        main_config['includes'].append(file_name)
    if 'tasks' not in main_config:
        main_config['tasks'] = {}
    main_task = Task(main_config['tasks']['run']) if 'run' in main_config['tasks'] else Task().set_icon('🚅').set_description('Run everything at once')
    main_task.add_dependency(task_name)
    main_config['tasks']['run'] = main_task.as_dict()
    write_config(main_file_name, main_config)


def update_task_env(task: Task, task_file_name: str) -> Task:
    location = task.get_location() 
    if not os.path.isabs(location):
        location = os.path.join(os.path.dirname(task_file_name), task.get_location())
    for env_file in ('sample.env', 'template.env', 'env.template', '.env'):
        env_path = os.path.join(location, env_file)
        if not os.path.isfile(env_path):
            continue
        local_env: Mapping[str, str] = dotenv_values(env_path)
        for env_key, env_value in local_env.items():
            env = task.get_env(env_key)
            if not env.get_from():
                env.set_from(get_task_env_name(location, env_key))
            if not env.get_default():
                env.set_default(env_value)
            task.set_env(env_key, env, env_value)
    return task


def get_task_file_name(task_name: str) -> str:
    generated_task_file_name = os.path.join('.', 'zaruba-tasks', '{}.zaruba.yaml'.format(task_name))
    return generated_task_file_name


def write_task_env(dir_name: str, task: Task):
    env_file_names = [os.path.join(dir_name, f) for f in os.listdir(dir_name) if os.path.isfile(os.path.join(dir_name, f)) and f.endswith('.env')]
    default_file_name = os.path.join(dir_name, '.env')
    if default_file_name not in env_file_names:
        env_file_names.append(default_file_name)
    for file_name in env_file_names:
        existing_envvars: Mapping[str, str] = dotenv_values(file_name) if os.path.isfile(file_name) else {}
        is_first_writing = True
        f_write = open(file_name, 'a')
        for env_key, env in task.get_all_env().items():
            env_from = env.get_from()
            if env_from == '' or env_from in existing_envvars:
                continue
            if is_first_writing:
                is_first_writing = False
                f_write.write('\n')
            env_value = task.get_env_value(env_key)
            f_write.write('{}={}\n'.format(env_from, env_value))
        f_write.close()


def get_alphanum_basename(location: str) -> str:
    abs_location = os.path.abspath(location)
    base_name = os.path.basename(abs_location)
    return re.sub(r'[^A-Za-z0-9]+', ' ', base_name)


def get_env_prefix(location: str) -> str:
    return get_alphanum_basename(location).upper().replace(' ', '_')


def get_task_env_name(location: str, env_name: str) -> str:
    upper_env_name = env_name.upper()
    env_prefix = get_env_prefix(location)
    if upper_env_name.startswith(env_prefix):
        return upper_env_name
    return '_'.join([env_prefix, upper_env_name])


def get_service_name(location: str) -> str:
    capital_service_name =  get_alphanum_basename(location).capitalize().replace(' ', '')
    return capital_service_name[0].lower() + capital_service_name[1:]


def get_container_name(image_name: str) -> str:
    capital_container_name =  get_alphanum_basename(image_name).capitalize().replace(' ', '')
    return capital_container_name[0].lower() + capital_container_name[1:]


def get_task_name(service_or_container: str) -> str:
    return 'run{}'.format(service_or_container.capitalize())


def replace_content(file_name: str, replace_dict: Mapping[str, str]):
    if not replace_dict:
        return
    f_read = open(file_name, 'r')
    content = f_read.read()
    f_read.close()
    for key, val in replace_dict.items():
        content = content.replace(key, val)
    f_write = open(file_name, 'w')
    f_write.write(content)
    f_write.close()


def copy_and_replace(source: str, destination: str, replace_dict: Mapping[str, str]):
    source = os.path.abspath(source)
    destination = os.path.abspath(destination)
    # create destination's parent in case of it doesn't exist
    destination_dir = os.path.dirname(destination)
    if not os.path.exists(destination_dir):
        os.makedirs(destination_dir)
    # the source is a file
    if os.path.isfile(source):
        shutil.copy(source, destination)
        replace_content(destination, replace_dict)
        return
    # the source is a directory, but the destination doesn't exist
    if os.path.isdir(destination):
        for root, source_dir_names, source_file_names in os.walk(source):
            for file_name in source_file_names:
                copy_and_replace(os.path.join(root, file_name), os.path.join(destination, file_name), replace_dict)
            for dir_name in source_dir_names:
                copy_and_replace(os.path.join(root, dir_name), os.path.join(destination, dir_name), replace_dict)
        return
    # the source is a directory and the destination is not exist
    shutil.copytree(source, destination)
    for root, dir_names, file_names in os.walk(destination):
        for file_name in file_names:
            replace_content(os.path.join(root, file_name), replace_dict)