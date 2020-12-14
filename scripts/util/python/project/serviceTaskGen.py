from typing import List, Mapping
from project.structures import Template, Task, Env
from project.helpers import add_to_main_include, add_to_main_task, adjust_task_env, create_dir, get_dict_from_file, get_task_env_name, get_service_name_by_location, get_generated_task_file_name, replace_file_content, save_dict_to_file, write_task_env
import os

class ServiceTaskGen():

    def __init__(self, template: Template, task_name: str, task_location: str, ports: List[str]):
        self.task_name = task_name
        self.task_location = task_location
        self.ports = ports
        self.template_dict = get_dict_from_file(template.location)
        self.task_file_name = get_generated_task_file_name(self.task_name)

    
    def generate_service_task(self):
        create_dir('zaruba-tasks')
        self.create_service_task_file()
        add_to_main_include(self.task_file_name)
        add_to_main_task(self.task_name)

    
    def create_service_task_file(self):
        task = self.load_task()
        # add env from location's file
        task = adjust_task_env(task, self.task_file_name)
        # add lconfig.ports
        ports = self.ports if self.ports else task.get_possible_ports()
        task.add_lconfig_ports(*ports)
        # save project and env
        project_dict = {'tasks': {self.task_name: task.as_dict()}}
        save_dict_to_file(self.task_file_name, project_dict)
        replace_dict = self.create_replace_dict()
        replace_file_content(self.task_file_name, replace_dict)
        write_task_env('.', task)

    
    def load_task(self) -> Task:
        task = Task(self.template_dict['tasks']['runService'])
        if os.path.isabs(self.task_location):
            task.set_location(self.task_location)
        else:
            task.set_location(os.path.join('..', self.task_location))
        # adjust env from template
        for key, env in task.get_all_env().items():
            envvar = env.get_from() if env.get_from() else key
            env.set_from(get_task_env_name(self.task_location, envvar))
            task.set_env(key, env)
        return task
  

    def create_replace_dict(self) -> Mapping[str, str]:
        replace_dict = {
            'MYSERVICE': get_service_name_by_location(self.task_location).upper(),
        }
        return replace_dict
