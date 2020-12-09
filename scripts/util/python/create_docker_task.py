from typing import List
import sys
import traceback

import project


# USAGE
# python create_docker_task.py <image> <container> <task>

def create_docker_task(template_path_list: List[str], image_name: str, container_name: str, task_name: str):
    try:
        image_name = image_name if image_name != '' else 'nginx'
        container_name = container_name if container_name != '' else project.get_container_name_by_image(image_name)
        task_name = task_name if task_name != '' else 'run{}'.format(container_name.capitalize())
        template = project.get_docker_task_template(template_path_list, image_name)
        gen = project.DockerTaskGen(template, task_name, image_name, container_name)
        gen.generate_docker_task()
        print('Task {} is created successfully'.format(task_name))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


if __name__ == '__main__':
    template_path_list = sys.argv[1].split(':') if len(sys.argv) > 1 and sys.argv[1] != '' else []
    image = sys.argv[2] if len(sys.argv) > 2 else ''
    container = sys.argv[3] if len(sys.argv) > 3 else ''
    task_name = sys.argv[4] if len(sys.argv) > 4 else ''
    create_docker_task(template_path_list, image, container, task_name)
