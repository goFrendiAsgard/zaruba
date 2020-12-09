from typing import List
import sys
import os
import traceback
import project


# USAGE
# python create_service.py <image> <container> <task>

def create_service(template_path_list: List[str], target_location: str, service_type: str):
    try:
        service_type = service_type if service_type != '' else 'fastapi'
        target_location = target_location if target_location != '' else './{}'.format(service_type)
        service_template = project.get_service_template(template_path_list, service_type)
        service_gen = project.ServiceGen(service_template, target_location)
        service_gen.generate_service()
        print('Service {} ({}) is created successfully'.format(target_location, service_type))
        if os.path.isfile(project.get_main_file_name()):
            ports = service_gen.ports
            task_template = project.get_service_task_template(template_path_list, service_type)
            task_name = project.get_task_name_by_location(target_location)
            task_gen = project.ServiceTaskGen(task_template, task_name, target_location, ports)
            task_gen.generate_service_task()
            print('Task {} is created successfully'.format(task_name))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


if __name__ == '__main__':
    template_path_list = sys.argv[1].split(':') if len(sys.argv) > 1 and sys.argv[1] != '' else []
    location = sys.argv[2] if len(sys.argv) > 2 else ''
    service_type = sys.argv[3] if len(sys.argv) > 3 else ''
    create_service(template_path_list, location, service_type)
