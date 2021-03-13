from typing import List
from helper import cli
import helper.generator as generator
import helper.task as task

import os, sys, traceback


service_containerization_tasks = '''

  zarubaRunContainerTask:
    icon: 🐳
    description: Run zarubaServiceName (containerized)
    extend: core.startDockerContainer
    dependencies:
    - zarubaBuildImageTask
    timeout: 1h
    env:
      <<: *zarubaServiceNameEnv
    lconfig:
      ports: *zarubaServiceNamePorts
    config:
      imageName: zarubaServiceName
      imageTag: latest
      containerName: zarubaServiceName
      rebuild: true
      localhost: host.docker.internal
      expose: lconfig.ports
    
  
  zarubaRemoveContainerTask:
    icon: 🐳
    description: Remove container for zarubaServiceName
    extend: core.removeDockerContainer 
    config:
      containerName: zarubaServiceName
  

  zarubaBuildImageTask:
    icon: 🐳
    description: Build image for zarubaServiceName
    extend: core.buildDockerImage
    location: zarubaTaskLocation
    timeout: 1h
    config:
      imageName: zarubaServiceName


  zarubaPushImageTask:
    icon: 🐳
    description: Push zarubaServiceName image
    extend: core.pushDockerImage
    dependencies:
    - zarubaBuildImageTask
    timeout: 1h
    config:
      imageName: zarubaServiceName

'''

@cli
def create_service_task(templates: str = '', location: str = '.', service_name: str = '', service_type: str = '', ports: str = ''):
    template_path_list = templates.split(':')
    service_ports=[service_port for service_port in ports.split(',') if service_port != '']
    location = location if location != '' else '.'
    service_type = service_type if service_type != '' else 'default'
    service_name = service_name if service_name != '' else generator.get_service_name(location)
    run_task_name = generator.get_run_task_name(service_name)
    task_file_name = generator.get_task_file_name(run_task_name)
    task_location = location if os.path.isabs(location) else os.path.join('..', location)
    build_image_task_name = generator.get_build_image_task_name(service_name)
    push_image_task_name = generator.get_push_image_task_name(service_name)
    run_container_task_name = generator.get_run_container_task_name(service_name)
    remove_container_task_name = generator.get_remove_container_task_name(service_name)
    template = get_service_task_template(template_path_list, service_type)
    generator.copy(template, task_file_name)
    add_containerization_tasks(task_file_name)
    generator.replace_all(task_file_name, {
        'zarubaRunTask': run_task_name,
        'zarubaBuildImageTask': build_image_task_name,
        'zarubaPushImageTask': push_image_task_name,
        'zarubaRunContainerTask': run_container_task_name,
        'zarubaRemoveContainerTask': remove_container_task_name,
        'zarubaServiceName': service_name,
        'zarubaTaskLocation': task_location,
        'ZarubaServiceName': service_name.capitalize(),
        'ZARUBA_ENV_PREFIX': generator.get_env_prefix(location),
    })
    config = generator.read_config(task_file_name)
    service_task = task.Task(config['tasks'][run_task_name])
    generator.update_task_env(service_task, task_file_name)
    if len(service_ports) == 0:
        service_ports = service_task.get_possible_ports()
    service_task.add_lconfig_ports(service_ports)
    generator.write_task_env('.', service_task)
    config['tasks'][run_task_name] = service_task.as_dict()
    generator.write_config(task_file_name, config)
    generator.register_run_task(task_file_name, run_task_name)
    generator.register_build_image_task(task_file_name, build_image_task_name)
    generator.register_push_image_task(task_file_name, push_image_task_name)
    generator.register_run_container_task(task_file_name, run_container_task_name)
    generator.register_remove_container_task(task_file_name, remove_container_task_name)
    print('Task {} is created successfully'.format(run_task_name))
    

def get_service_task_template(template_path_list: List[str], service_type:str) -> str:
    for template_path in template_path_list:
        template = '{template_path}/task/service/{service_type}.zaruba.yaml'.format(template_path=template_path, service_type=service_type)
        if os.path.isfile(template):
            return template
    return get_service_task_template(template_path_list, 'default')


def add_containerization_tasks(file_name: str):
    f = open(file_name, 'a')
    f.write(service_containerization_tasks)
    f.close()


if __name__ == '__main__':
    create_service_task()