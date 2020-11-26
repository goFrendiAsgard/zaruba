from project.task import Task
from ruamel.yaml import YAML


def add_task(filename: str, task_name: str, task: Task):
    yaml = YAML()
    project = {}
    try:
        project = yaml.load(open(filename, 'r'))
    except:
        pass
    if 'tasks' not in project:
        project['tasks'] = {}
    if task_name in project.tasks:
        return False
    project['tasks'][task_name] = task.to_dict()
    yaml.dump(project, open(filename, 'w'))


def add_include(filename: str) -> bool:
    yaml = YAML()
    main_filename = 'main.zaruba.yaml'
    main_project = yaml.load(open(main_filename, 'r'))
    if filename in main_project:
        return False
    if 'includes' not in main_project:
        main_project['includes'] = []
    main_project['includes'].append(filename)
    yaml.dump(main_project, open(main_filename, 'w'))
    return True
