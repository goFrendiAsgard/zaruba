from typing import Any, Mapping, List
import sys
import shutil
import traceback
from dotenv import dotenv_values

import project


# USAGE
# python create_service.py <image> <container> <task>

def create_service(template_path_list: List[str], service_location: str, service_type: str, ports=List[str]):
    try:
        service_type = service_type if service_type != '' else 'fastapi'
        service_location = service_location if service_location != '' else './{}'.format(service_type)
        project.create_dir('zaruba-tasks')
        template_dir_name, _ = project.get_service_template(template_path_list, service_type)
        shutil.copytree(template_dir_name, service_location)
        print('Service {} ({}) is created successfully'.format(service_location, service_type))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


if __name__ == '__main__':
    template_path_list = sys.argv[1].split(':') if len(sys.argv) > 1 and sys.argv[1] != '' else []
    location = sys.argv[2] if len(sys.argv) > 2 else ''
    service_type = sys.argv[3] if len(sys.argv) > 3 else ''
    ports = sys.argv[4].split(',') if len(sys.argv) > 4 and sys.argv[4] != '' else []
    create_service(template_path_list, location, service_type, ports)
