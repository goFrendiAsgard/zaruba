from typing import List
from helper import cli
import helper.generator as generator
import helper.decoration as decoration
import helper.task as task
import os


@cli
def create_docker_task(templates: str = '', image_name: str = '', image_prefix: str = '', image_tag = '', container_name: str = '', service_name: str = ''):
    template_path_list = templates.split(':')
    task_template = get_docker_task_template(template_path_list, image_name)
    image_name = image_name if image_name != '' else 'nginx'
    container_name = container_name if container_name != '' else generator.get_container_name(image_name)
    run_task_name = 'run' + service_name.capitalize() if service_name != '' else generator.get_run_task_name(container_name)
    task_file_name = generator.get_task_file_name(run_task_name)
    task_file_content = generator.replace_str(generator.read_text(task_template), {
        'zarubaRunTask': run_task_name,
        'zarubaImageName': image_name,
        'zarubaContainerName': container_name,
    })
    generator.write_text(task_file_name, task_file_content)
    config = generator.read_config(task_file_name)
    run_docker_task = task.Task(config['tasks'][run_task_name])
    generator.write_task_env('.', run_docker_task)
    generator.register_run_task(task_file_name, run_task_name)
    generator.register_run_container_task(task_file_name, run_task_name)
    if image_tag != '' or image_prefix != '':
        if image_tag != '':
            run_docker_task.set_config('imageTag', image_tag)
        if image_prefix != '':
            run_docker_task.set_config('imagePrefix', image_prefix)
        config['tasks'][run_task_name] = run_docker_task
        generator.write_config(task_file_name, config)
    print('Task {} is created successfully'.format(run_task_name))


def get_docker_task_template(template_path_list: List[str], image_name:str) -> str:
    for template_path in template_path_list:
        template = '{template_path}/task/docker/{image_name}.zaruba.yaml'.format(template_path=template_path, image_name=image_name)
        if os.path.isfile(template):
            return template
    return generator.get_docker_task_template(template_path_list, 'default')
    

if __name__ == '__main__':
    create_docker_task()