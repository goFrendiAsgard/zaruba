from typing import Mapping, Any, List
import re
import os

from dotenv import dotenv_values
from project.structures import Task, ProjectDict, Template, Env
from ruamel.yaml import YAML


def write_task(file_name: str, task_name: str, task: Task):
    project_dict: ProjectDict = {}
    try:
        project_dict = get_dict_from_file(file_name)
    except:
        pass
    if 'tasks' not in project_dict:
        project_dict['tasks'] = {}
    if task_name in project_dict.tasks:
        return False
    project_dict['tasks'][task_name] = task.as_dict()
    save_dict_to_file(file_name, project_dict)


def add_to_main_include(file_name: str) -> bool:
    main_file_name = get_main_file_name()
    main_project_dict = get_dict_from_file(main_file_name)
    if 'includes' not in main_project_dict:
        main_project_dict['includes'] = []
    if file_name in main_project_dict['includes']:
        return False
    main_project_dict['includes'].append(file_name)
    save_dict_to_file(main_file_name, main_project_dict)
    return True


def add_to_main_task(task_name: str) -> bool:
    main_file_name = get_main_file_name()
    main_project_dict = get_dict_from_file(main_file_name)
    if 'tasks' not in main_project_dict:
        main_project_dict['tasks'] = {}
    task = Task(main_project_dict['tasks']['run']) \
        if 'run' in main_project_dict['tasks'] \
        else Task().set_icon('🚅').set_description('Run everything at once')
    task.add_dependency(task_name)
    main_project_dict['tasks']['run'] = task.as_dict()
    save_dict_to_file(main_file_name, main_project_dict)
    return True


def get_main_file_name() -> str:
    main_file_name = 'main.zaruba.yaml'
    return main_file_name


def get_generated_task_file_name(task_name: str) -> str:
    generated_task_file_name = os.path.join('.', 'zaruba-tasks', '{}.zaruba.yaml'.format(task_name))
    if os.path.isfile(generated_task_file_name):
        raise Exception('{} already exists'.format(generated_task_file_name))
    return generated_task_file_name


def create_dir(dirname: str):
    if not os.path.exists(dirname):
        os.makedirs(dirname)


def get_dict_from_file(file_name: str) -> Mapping[str, Any]:
    yaml=YAML()
    f = open(file_name, 'r')
    template_obj = yaml.load(f)
    f.close()
    return template_obj


def save_dict_to_file(file_name: str, dictionary: Mapping[str, Any]):
    yaml=YAML()
    f = open(file_name, 'w')
    yaml.dump(dictionary, f)
    f.close()


def write_task_env_to_file(file_name: str, task: Task):
    existing_envvars: Mapping[str, str] = dotenv_values(file_name) if os.path.isfile(file_name) else {}
    is_first_writing = True
    f_write = open(file_name, 'a')
    for _, env in task.get_all_env().items():
        envvar = env.get_from()
        if envvar == '' or envvar in existing_envvars:
            continue
        value = env.get_default()
        if is_first_writing:
            is_first_writing = False
            f_write.write('\n')
        f_write.write('{}={}\n'.format(envvar, value))
    f_write.close()


def write_task_env(dir_name: str, task: Task):
    env_file_names = [os.path.join(dir_name, f) for f in os.listdir(dir_name) if os.path.isfile(os.path.join(dir_name, f)) and f.endswith('.env')]
    default_file_name = os.path.join(dir_name, '.env')
    if default_file_name not in env_file_names:
        env_file_names.append(default_file_name)
    for file_name in env_file_names:
        write_task_env_to_file(file_name, task)


def get_default_template_location() -> str:
    return os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(__file__)))),
        'templates'
    )


def get_template(template_path_list: List[str], location: str, default_location: str, location_is_dir=False) -> Template:
    real_template_path_list = [os.path.join('.', 'templates')]
    real_template_path_list.extend(template_path_list)
    default_template_location = get_default_template_location()
    real_template_path_list.append(default_template_location)
    for template_path in real_template_path_list:
        template_location = os.path.join(template_path, location)
        if (location_is_dir and os.path.isdir(template_location)) or os.path.isfile(template_location):
            return Template(template_location, False)
    return Template(os.path.join(default_template_location, default_location), True)


def get_service_task_template(template_path_list: List[str], service_type: str) -> Template:
    service_task_template_dir = 'service-task'
    location = os.path.join(service_task_template_dir, '{}.zaruba.yaml'.format(service_type))
    default_location = os.path.join(service_task_template_dir, 'default.zaruba.yaml')
    return get_template(template_path_list, location, default_location, location_is_dir=False)


def get_docker_task_template(template_path_list: List[str], image: str) -> Template:
    docker_task_template_dir = 'docker-task'
    location = os.path.join(docker_task_template_dir, '{}.zaruba.yaml'.format(image))
    default_location = os.path.join(docker_task_template_dir, 'default.zaruba.yaml')
    return get_template(template_path_list, location, default_location, location_is_dir=False)


def get_service_template(template_path_list: List[str], service_type: str) -> Template:
    service_template_dir = 'service'
    location = os.path.join(service_template_dir, service_type)
    default_location = os.path.join(service_template_dir, 'fastapi')
    return get_template(template_path_list, location, default_location, location_is_dir=True)


def get_sanitized_base_name(location: str) -> str:
    abs_location = os.path.abspath(location)
    base_name = os.path.basename(abs_location)
    return re.sub(r'[^A-Za-z0-9]+', ' ', base_name)


def get_env_prefix_by_location(location: str) -> str:
    return get_sanitized_base_name(location).upper().replace(' ', '_')


def get_task_env_name(location: str, raw_env_name: str) -> str:
    upper_env_name = raw_env_name.upper()
    env_prefix = get_env_prefix_by_location(location)
    if upper_env_name.startswith(env_prefix):
        return upper_env_name
    return '_'.join([env_prefix, upper_env_name])


def adjust_task_env(task: Task, task_file_name: str, override: bool = True) -> Task:
    location = task.get_location() if os.path.isabs(task.get_location()) else os.path.join(os.path.dirname(task_file_name), task.get_location())
    for env_file in ('sample.env', 'template.env', 'env.template', '.env'):
        env_path = os.path.join(location, env_file)
        if not os.path.isfile(env_path):
            continue
        raw_env_dict: Mapping[str, str] = dotenv_values(env_path)
        for key, default in raw_env_dict.items():
            env = Env({
                'from': get_task_env_name(location, key),
                'default': default,
            })
            task.set_env(key, env, override)
    return task


def get_service_name_by_location(location: str) -> str:
    capitalized_service_name =  get_sanitized_base_name(location).capitalize().replace(' ', '')
    return capitalized_service_name[0].lower() + capitalized_service_name[1:]


def get_task_name_by_location(location: str) -> str:
    return 'run{}'.format(get_sanitized_base_name(location).capitalize().replace(' ', ''))


def get_container_name_by_image(image_name: str) -> str:
    return get_service_name_by_location(image_name)


def replace_file_content(file_name: str, replace_dict: Mapping[str, str]):
    f_read = open(file_name, 'r')
    lines = f_read.readlines()
    f_read.close()
    for index, line in enumerate(lines):
        for key, val in replace_dict.items():
            lines[index] = line.replace(key, val)
    f_write = open(file_name, 'w')
    f_write.writelines(lines)
    f_write.close()
