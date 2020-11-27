from typing import Mapping, Any
import os
from project.structure import Task, ProjectDict
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
    project_dict['tasks'][task_name] = task.to_dict()
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
    if 'run' not in main_project_dict['tasks']:
        task = Task()
        task.icon = '🚒' 
        task.description = 'Run everything at once'
        task_dict = task.to_dict()
        main_project_dict['tasks']['run'] = task_dict
    if 'dependencies' not in main_project_dict['tasks']['run']:
        main_project_dict['tasks']['run']['dependencies'] = []
    if task_name in main_project_dict['tasks']['run']['dependencies']:
        return False
    main_project_dict['tasks']['run']['dependencies'].append(task_name)
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

