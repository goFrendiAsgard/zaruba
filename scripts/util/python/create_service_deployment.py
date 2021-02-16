from typing import Any, List, Mapping
from ruamel.yaml import YAML
from common_helper import get_argv
from generator_helper import get_service_name, get_env_in_location, read_config, write_config
from port_helper import get_possible_ports_env

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


def register_helmfile(service_name: str):
    helmfile_location = os.path.join('helm-deployments', 'helmfile.yaml')
    helmfile_dict = read_config(helmfile_location)
    helmfile_dict['releases'].append({
        'name': service_name,
        'chart': './charts/app',
        'values': ['./values/{}.yaml.gotmpl'.format(service_name)]
    })
    write_config(helmfile_location, helmfile_dict)


def create_service_deployment(location: str, ports: List[str]):
    location = location if location != '' else '.'
    service_name = get_service_name(location)
    value_file_name = os.path.join('helm-deployments', 'values', '{}.yaml.gotmpl'.format(service_name))
    value_dict = read_value_template_dict()
    value_dict['app']['name'] = service_name
    value_dict['app']['container']['image'] = service_name
    # add env
    env_dict = get_env_in_location(location)
    for env_key, env_value in env_dict.items():
        value_dict['app']['container']['env'].append({
            'name': env_key,
            'value': env_value,
        })
    # add ports
    if len(ports) == 0:
        possible_ports_env = get_possible_ports_env(env_dict)
        ports = list(possible_ports_env.values())
    for port in ports:
        value_dict['app']['ports'].append({
            'containerPort': port,
            'servicePort': port
        })
    write_config(value_file_name, value_dict)
    register_helmfile(service_name)


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