from helper import cli
from helper.project import DockerProject

@cli
def create_docker_task(template_location: str, image_name: str = '', container_name: str = '', service_name: str = ''):
    if image_name == '' and container_name == '':
        template_file_name = os.path.basename(template_location)
        container_name = template_file_name.split('.')[0]
    dir_name = '.'
    docker_project = DockerProject()
    docker_project.load_from_template(template_location)
    docker_project.generate(dir_name=dir_name, service_name=service_name, image_name=image_name, container_name=container_name)

if __name__ == '__main__':
    create_docker_task()