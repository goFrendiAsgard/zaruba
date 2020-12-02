from typing import Mapping, Any
import os
from project.structure import Task, Env, ProjectDict, EnvDict
from ruamel.yaml import YAML


def add_task(file_name: str, task_name: str, task: Task):
    project_dict: ProjectDict = {}
    try:
        project_dict = get_dict(file_name)
    except:
        pass
    if 'tasks' not in project_dict:
        project_dict['tasks'] = {}
    if task_name in project_dict.tasks:
        return False
    project_dict['tasks'][task_name] = task.as_dict()
    write_dict(file_name, project_dict)


def add_to_include(file_name: str) -> bool:
    main_file_name = get_main_file_name()
    main_project_dict = get_dict(main_file_name)
    if 'includes' not in main_project_dict:
        main_project_dict['includes'] = []
    if file_name in main_project_dict['includes']:
        return False
    main_project_dict['includes'].append(file_name)
    write_dict(main_file_name, main_project_dict)
    return True


def add_to_run_task(task_name: str) -> bool:
    main_file_name = get_main_file_name()
    main_project_dict = get_dict(main_file_name)
    if 'tasks' not in main_project_dict:
        main_project_dict['tasks'] = {}
    task = Task(main_project_dict['tasks']['run']) \
        if 'run' in main_project_dict['tasks'] \
        else Task().set_icon('🚅').set_description('Run everything at once')
    task.add_dependency(task_name)
    main_project_dict['tasks']['run'] = task.as_dict()
    write_dict(main_file_name, main_project_dict)
    return True


def get_main_file_name() -> str:
    main_file_name = 'main.zaruba.yaml'
    return main_file_name


def create_dir(dirname: str):
    if not os.path.exists(dirname):
        os.makedirs(dirname)


def get_dict(file_name: str) -> Mapping[str, Any]:
    yaml=YAML()
    f = open(file_name, 'r')
    template_obj = yaml.load(f)
    f.close()
    return template_obj


def write_dict(file_name: str, dictionary: Mapping[str, Any]):
    yaml=YAML()
    f = open(file_name, 'w')
    yaml.dump(dictionary, f)
    f.close()

def write_task_env(file_name: str, task: Task):
    f = open(file_name, 'a')
    f.write('\n')
    for _, env in task.get_all_env().items():
        envvar = env.get_from()
        if envvar == '':
            continue
        value = env.get_default()
        f.write('{}={}\n'.format(envvar, value))
    f.close()

