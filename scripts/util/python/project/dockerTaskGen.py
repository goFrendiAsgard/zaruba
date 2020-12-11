from project.structures import Template, Task, Env
from project.helpers import add_to_main_include, add_to_main_task, create_dir, get_dict_from_file, get_task_file_name, save_dict_to_file, write_task_env


class DockerTaskGen():

    def __init__(self, template: Template, task_name: str, image_name: str, container_name: str):
        self.is_default_template = template.is_default
        self.image_name = image_name
        self.container_name = container_name
        self.task_name = task_name
        self.template_dict = get_dict_from_file(template.location)
        self.task_file_name = get_task_file_name(self.task_name)

    
    def generate_docker_task(self):
        create_dir('zaruba-tasks')
        self.create_service_task_file()
        add_to_main_include(self.task_file_name)
        add_to_main_task(self.task_name)

    
    def create_service_task_file(self):
        task = self.load_task()
        project_dict = {'tasks': {self.task_name: task.as_dict()}}
        save_dict_to_file(self.task_file_name, project_dict)
        write_task_env('.', task)

    
    def load_task(self) -> Task:
        task = Task(self.template_dict['tasks']['runContainer'])
        task.set_config('containerName', self.container_name)
        if self.is_default_template:
            task.set_config('imageName', self.image_name)
        return task
