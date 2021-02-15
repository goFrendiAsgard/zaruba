from typing import Any, List, Mapping
from ruamel.yaml import YAML
from common_helper import get_argv
from generator_helper import get_service_name, get_env_in_location, write_config

import os, sys, traceback

# USAGE
# python create_service_deployment.py


def read_value_template_dict() -> Mapping[str, Mapping[str, Any]]:
    yaml = YAML()
    file_name = os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))),
        'templates',
        'helmValues',
        'app.yaml.gotmpl'
    )
    f = open(file_name, 'r')
    value_dict = yaml.load(f)
    f.close()
    return value_dict


def create_service_deployment(location: str, ports: List[str]):
    location = location if location != '' else '.'
    service_name = get_service_name(location)
    value_file_name = os.path.join('helm-deployments', 'values', '{}.yaml.gotmpl'.format(service_name))
    value_dict = read_value_template_dict()
    value_dict['app']['name'] = service_name
    value_dict['app']['container']['image'] = service_name
    env_dict = get_env_in_location(location)
    for env_key, env_value in env_dict.items():
        value_dict['app']['container']['env'].append({
            'name': env_key,
            'value': env_value,
        })
    # TODO: get possible ports if ports keyword in not available (sementara pake if gini dulu)
    if len(ports) == 0:
        ports = ['8080']
    for port in ports:
        value_dict['app']['ports'].append({
            'containerPort': port,
            'servicePort': port
        })
    write_config(value_file_name, value_dict)
    # TODO: register values to helmfile


if __name__ == '__main__':
    try:
        location = get_argv(1)
        raw_ports = get_argv(2).split(',')
        ports=[port for port in raw_ports if port != '']
        create_service_deployment(location, ports)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)