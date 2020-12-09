from typing import List
import sys
import traceback

import project


# USAGE
# python create_service_task.py <image> <container> <task>

def create_service_task(template_path_list: List[str], task_location: str, service_type: str, task_name: str, ports=List[str]):
    try:
        service_type = service_type if service_type != '' else 'default'
        task_location = task_location if task_location != '' else './'
        task_name = task_name if task_name != '' else project.get_task_name_by_location(task_location)
        template = project.get_service_task_template(template_path_list, service_type)
        gen = project.ServiceTaskGen(template, task_name, task_location, ports)
        gen.generate_service_task()
        print('Task {} is created successfully'.format(task_name))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


if __name__ == '__main__':
    template_path_list = sys.argv[1].split(':') if len(sys.argv) > 1 and sys.argv[1] != '' else []
    location = sys.argv[2] if len(sys.argv) > 2 else ''
    service_type = sys.argv[3] if len(sys.argv) > 3 else ''
    task_name = sys.argv[4] if len(sys.argv) > 4 else ''
    ports = sys.argv[5].split(',') if len(sys.argv) > 5 and sys.argv[5] != '' else []
    create_service_task(template_path_list, location, service_type, task_name, ports)
