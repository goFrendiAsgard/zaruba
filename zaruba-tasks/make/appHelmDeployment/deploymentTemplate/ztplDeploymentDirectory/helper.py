from typing import Any, Mapping, Optional
import os
import json

DEFAULT_CONFIG = {
    'env': [],
    'image.repository': '',
    'image.tag': '',
    'ports': [],
    'service.ports': []
}

class Config:

    def __init__(self, config_file_path: str):
        self.config_map: Mapping[str, Any] = self.load_config(config_file_path)


    def load_config(self, file_path: str) -> Mapping[str, Any]:
        if not os.path.exists(file_path):
            return DEFAULT_CONFIG
        with open(file_path) as f:
            return json.load(f)
        

    def get(self, key: str, default_value: Any = None) -> Any:
        return self.config_map.get(key, default_value)
    

    def get_config_map(self) -> Mapping[str, Any]:
        return self.config_map


class Probe:

    def __init__(self, config: Config):
        self.config = config


    def get_liveness_config(self) -> Mapping[str, Any]:
        return self._get_probe_config('LIVENESS_PROBE')


    def get_readiness_config(self) -> Mapping[str, Any]:
        return self._get_probe_config('READINESS_PROBE')


    def _get_probe_config(self, env_prefix: str) -> Mapping[str, Any]:
        method = os.getenv('{}_METHOD'.format(env_prefix), '')
        probe_config = {}
        if method == 'exec':
            probe_config = self._create_exec_probe(env_prefix)
        elif method == 'httpGet':
            probe_config = self._create_http_get_probe(env_prefix)
        elif method == 'tcpSocket':
            probe_config = self._create_http_get_probe(env_prefix)
        if len(probe_config) == 0:
            return probe_config
        # add properties
        probe_config['failureThreshold'] = int(os.getenv('{}_FAILURE_THRESHOLD'.format(env_prefix), '3'))
        probe_config['initialDelaySeconds'] = int(os.getenv('{}_INITIAL_DELAY_SECONDS'.format(env_prefix), '3'))
        probe_config['periodSeconds'] = int(os.getenv('{}_PERIOD_SECONDS'.format(env_prefix), '10'))
        probe_config['successThreshold'] = int(os.getenv('{}_SUCCESS_THRESHOLD'.format(env_prefix), '1'))
        probe_config['timeoutSeconds'] = int(os.getenv('{}_TIMEOUT_SECONDS'.format(env_prefix), '1'))
        return probe_config


    def _create_http_get_probe(self, env_prefix: str) -> Mapping[str, Any]:
        port = self._create_default_port_by_env_and_config('{}_HTTP_GET_PORT'.format(env_prefix))
        if port is None:
            return {}
        return {
            'httpGet': {
                'path': os.getenv('{}_HTTP_GET_PATH'.format(env_prefix), '/'),
                'port': port, 
                'scheme': os.getenv('{}_HTTP_GET_SCHEME'.format(env_prefix), 'http'),
            }
        }


    def _create_tcp_socket_probe(self, env_prefix: str) -> Mapping[str, Any]:
        port = self._create_default_port_by_env_and_config('{}_TCP_SOCKET_PORT'.format(env_prefix))
        if port is None:
            return {}
        return {
            'tcpSocket': {
                'port': port, 
            }
        }


    def _create_exec_probe(self, env_prefix: str) -> Mapping[str, Any]:
        return {
            'exec': {
                'command': os.getenv('{}_EXEC_COMMAND'.format(env_prefix), 'echo "default probe"')
            }
        }


    def _create_default_port_by_env_and_config(self, env_name: str) -> Any:
        port = os.getenv(env_name, '')
        if port != '':
            if port.isnumeric():
                return int(port)
            return port
        config_map = self.config.get_config_map()
        if 'ports' in config_map and len(config_map['ports']) > 0 and 'name' in config_map['ports'][0] and config_map['ports'][0]['name']:
            return config_map['ports'][0]['name']
        return None
