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
      imageName: &zarubaServiceNameImage zarubaImageName
      imageTag: latest
      containerName: &zarubaServiceNameContainer zarubaContainerName
      rebuild: true
      localhost: host.docker.internal
      expose: lconfig.ports
    
  
  zarubaRemoveContainerTask:
    icon: 🐳
    description: Remove zarubaServiceName's container
    extend: core.removeDockerContainer 
    config:
      containerName: *zarubaServiceNameContainer
  

  zarubaBuildImageTask:
    icon: 🐳
    description: Build zarubaServiceName's image
    extend: core.buildDockerImage
    location: zarubaTaskLocation
    timeout: 1h
    config:
      imageName: *zarubaServiceNameImage


  zarubaPushImageTask:
    icon: 🐳
    description: Push zarubaServiceName's image
    extend: core.pushDockerImage
    dependencies:
    - zarubaBuildImageTask
    timeout: 1h
    config:
      imageName: *zarubaServiceNameImage

'''

@cli
def create_service_task(templates: str = '', location: str = '.', service_name: str = '', service_type: str = '', ports: str = '', image_name: str = ''):
    template_path_list = templates.split(':')
    service_ports=[service_port for service_port in ports.split(',') if service_port != '']
    location = location if location != '' else '.'
    service_type = service_type if service_type != '' else 'default'
    print('SERVICE NAME AWAL', service_name)
    service_name = service_name if service_name != '' else generator.get_service_name(location)
    print('SERVICE NAME AKHIR', service_name)
    container_name = service_name
    image_name = image_name if image_name != '' else container_name.lower()
    run_task_name = generator.get_run_task_name(service_name)
    task_location = location if os.path.isabs(location) else os.path.join('..', location)
    build_image_task_name = generator.get_build_image_task_name(service_name)
    push_image_task_name = generator.get_push_image_task_name(service_name)
    run_container_task_name = generator.get_run_container_task_name(service_name)
    remove_container_task_name = generator.get_remove_container_task_name(service_name)
    task_template = get_service_task_template(template_path_list, service_type)
    task_file_name = generator.get_task_file_name(service_name)
    task_file_content = generator.replace_str(generator.read_text(task_template) + service_containerization_tasks, {
        'zarubaRunTask': run_task_name,
        'zarubaBuildImageTask': build_image_task_name,
        'zarubaPushImageTask': push_image_task_name,
        'zarubaRunContainerTask': run_container_task_name,
        'zarubaRemoveContainerTask': remove_container_task_name,
        'zarubaServiceName': service_name,
        'zarubaContainerName': container_name,
        'zarubaImageName': image_name,
        'zarubaTaskLocation': task_location,
        'ZarubaServiceName': service_name.capitalize(),
        'ZARUBA_ENV_PREFIX': generator.get_env_prefix(location),
    })
    generator.write_text(task_file_name, task_file_content)
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


if __name__ == '__main__':
    create_service_task()