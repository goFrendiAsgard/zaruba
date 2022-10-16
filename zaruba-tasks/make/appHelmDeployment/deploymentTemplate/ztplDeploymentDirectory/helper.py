from typing import Any, Mapping
import os
import json

def load_config(file_path: str) -> Mapping[str, Any]:
    with open(file_path) as f:
        return json.load(f)


def get_probe(env_prefix: str, config: Mapping[str, Any]) -> Mapping[str, Any]:
    method = os.getenv('{}_METHOD'.format(env_prefix), '')
    probe_config = {}
    if method == 'exec':
        probe_config = get_exec_probe(env_prefix)
    elif method == 'httpGet':
        probe_config = get_http_get_probe(env_prefix, config)
    elif method == 'tcpSocket':
        probe_config = get_http_get_probe(env_prefix, config)
    if len(probe_config) == 0:
        return probe_config
    # add properties
    probe_config['failureThreshold'] = int(os.getenv('{}_FAILURE_THRESHOLD'.format(env_prefix), '3'))
    probe_config['initialDelaySeconds'] = int(os.getenv('{}_INITIAL_DELAY_SECONDS'.format(env_prefix), '3'))
    probe_config['periodSeconds'] = int(os.getenv('{}_PERIOD_SECONDS'.format(env_prefix), '10'))
    probe_config['successThreshold'] = int(os.getenv('{}_SUCCESS_THRESHOLD'.format(env_prefix), '1'))
    probe_config['timeoutSeconds'] = int(os.getenv('{}_TIMEOUT_SECONDS'.format(env_prefix), '1'))
    return probe_config


def get_http_get_probe(env_prefix: str, config: Mapping[str, Any]) -> Mapping[str, Any]:
    port = get_default_port_by_env_and_config('{}_HTTP_GET_PORT'.format(env_prefix), config)
    if port is None:
        return {}
    return {
        'httpGet': {
            'path': os.getenv('{}_HTTP_GET_PATH'.format(env_prefix), '/'),
            'port': port, 
            'scheme': os.getenv('{}_HTTP_GET_SCHEME'.format(env_prefix), 'http'),
        }
    }


def get_tcp_socket_probe(env_prefix: str, config: Mapping[str, Any]) -> Mapping[str, Any]:
    port = get_default_port_by_env_and_config('{}_TCP_SOCKET_PORT'.format(env_prefix), config)
    if port is None:
        return {}
    return {
        'tcpSocket': {
            'port': port, 
        }
    }


def get_default_port_by_env_and_config(env_name: str, config: Mapping[str, Any]) -> Any:
    port = os.getenv(env_name, '')
    if port != '':
        if port.isnumeric():
            return int(port)
        return port
    if 'ports' in config and len(config['ports']) > 0 and 'name' in config['ports'][0] and config['ports'][0]['name']:
        return config['ports'][0]['name']
    return None


def get_exec_probe(env_prefix: str) -> Mapping[str, Any]:
    return {
        'exec': {
            'command': os.getenv('{}_EXEC_COMMAND'.format(env_prefix), 'echo "default probe"')
        }
    }

