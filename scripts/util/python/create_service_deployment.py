from typing import Any, List, Mapping
from helper import cli
import helper.generator as generator
import helper.port as port
from ruamel.yaml import YAML

import os

# USAGE
# python create_service_deployment.py

@cli
def create_service_deployment(service_location: str = '', ports: str = ''):
    service_location = service_location if service_location != '' else '.'
    service_ports=[service_port for service_port in ports.split(',') if service_port != '']
    service_name = generator.get_service_name(service_location).lower()
    value_file_name = os.path.join('helm-deployments', 'values', '{}.yaml.gotmpl'.format(service_name))
    value_dict = read_value_template_dict()
    value_dict['app']['name'] = service_name
    value_dict['app']['container']['image'] = service_name
    # add env
    env_dict = generator.get_env_in_location(service_location)
    for env_key, env_value in env_dict.items():
        value_dict['app']['container']['env'].append({
            'name': env_key,
            'value': env_value,
        })
    # add ports
    if len(service_ports) == 0:
        possible_ports_env = port.get_possible_ports_env(env_dict)
        service_ports = list(possible_ports_env.values())
    for service_port in service_ports:
        value_dict['app']['ports'].append({
            'containerPort': service_port,
            'servicePort': service_port
        })
    generator.write_config(value_file_name, value_dict)
    register_helmfile(service_name)


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
    helmfile_dict = generator.read_config(helmfile_location)
    helmfile_dict['releases'].append({
        'name': service_name,
        'chart': './charts/app',
        'values': ['./values/{}.yaml.gotmpl'.format(service_name)]
    })
    generator.write_config(helmfile_location, helmfile_dict)


if __name__ == '__main__':
    create_service_deployment()