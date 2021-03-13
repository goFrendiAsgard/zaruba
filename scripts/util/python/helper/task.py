from __future__ import annotations
from typing import List, Mapping, TypedDict
from dotenv import dotenv_values
from .port import get_service_possible_ports_env
import os

EnvDict = TypedDict('envDict', {'from': str, 'default': str}, total=False)

class TaskDict(TypedDict, total=False):
    icon: str
    extend: str
    description: str
    location: str
    dependencies: List[str]
    env: EnvDict
    config: Mapping[str, str]
    lconfig: Mapping[str, List[str]]
    start: List[str]
    check: list[str]
    private: bool
    timeout: str


class Env():

    def __init__(self, env_dict: EnvDict):
        self._dict: EnvDict = env_dict

    def set_from(self, envvar: str) -> Env:
        if envvar:
            self._dict['from'] = envvar
        elif 'from' in self._dict:
            del self._dict['from']
        return self

    def get_from(self) -> str:
        if 'from' in self._dict:
            return self._dict['from']
        return ''

    def set_default(self, default: str) -> Env:
        self._dict['default'] = str(default)
        return self

    def get_default(self) -> str:
        if 'default' in self._dict:
            return str(self._dict['default'])
        return ''  

    def as_dict(self) -> EnvDict:
        return self._dict


class Task():

    def __init__(self, task_dict: TaskDict):
        self._dict: TaskDict = task_dict
        self._env_values: Mapping[str, str] = {}

    def set_icon(self, icon: str) -> Task:
        if icon:
            self._dict['icon'] = icon
        elif 'icon' in self._dict:
            del self._dict['icon']
        return self

    def set_extend(self, extend: str) -> Task:
        if extend:
            self._dict['extend'] = extend
        elif 'extend' in self._dict:
            del self._dict['extend']
        return self
    
    def get_extend(self) -> str:
        return self._dict.get('extend', '')

    def set_description(self, description: str) -> Task:
        if description:
            self._dict['description'] = description
        elif 'description' in self._dict:
            del self._dict['description']
        return self

    def set_location(self, location: str) -> Task:
        if location:
            self._dict['location'] = location
        elif 'location' in self._dict:
            del self._dict['location']
        return self

    def add_dependency(self, *dependencies: str) -> Task:
        if 'dependencies' not in self._dict:
            self._dict['dependencies'] = []
        for dependency in dependencies:
            if dependency not in self._dict['dependencies']:
                self._dict['dependencies'].append(dependency)
        return self     

    def set_env(self, env_key: str, env: Env, env_value: str = '') -> Task:
        if 'env' not in self._dict:
            self._dict['env'] = {}
        if env_key not in self._dict['env']:
            self._dict['env'][env_key] = env.as_dict()
        self._env_values[env_key] = env_value
        return self

    def get_all_env(self) -> Mapping[str, Env]:
        env_map: Mapping[str, Env] = {}
        if 'env' in self._dict:
            for key in self._dict['env']:
                env_map[key] = self.get_env(key)
        return env_map

    def get_env(self, env_key: str) -> Env:
        if ('env' in self._dict) and (env_key in self._dict['env']):
            return Env(self._dict['env'][env_key])
        return Env({})

    def get_env_value(self, env_key: str) -> str:
        if env_key in self._env_values:
            value = self._env_values[env_key]
            if value:
                return value
        env = self.get_env(env_key)
        return env.get_default()

    def set_config(self, config_key: str, val: str) -> Task:
        if 'config'not in self._dict:
            self._dict['config'] = {}
        self._dict['config'][config_key] = val
        return self
    
    def get_config(self, config_key: str) -> str:
        if ('config' in self._dict) and (config_key in self._dict['config']):
            return self._dict['config'][config_key]
        return ''

    def init_lconfig(self, lconfig_key: str) -> Task:
        if 'lconfig'not in self._dict:
            self._dict['lconfig'] = {}
        if lconfig_key not in self._dict['lconfig']:
            self._dict['lconfig'][lconfig_key] = []

    def add_lconfig(self, lconfig_key: str, *vals: str) -> Task:
        self.init_lconfig(lconfig_key)
        for val in vals:
            self._dict['lconfig'][lconfig_key].append(val)
        return self

    def add_unique_lconfig(self, lconfig_key: str, *vals: str) -> Task:
        self.init_lconfig(lconfig_key)
        for val in vals:
            if val in self._dict['lconfig'][lconfig_key]:
                continue
            self._dict['lconfig'][lconfig_key].append(val)
        return self

    def add_lconfig_port(self, port: str) -> Task:
        port = port.strip()
        if port == '':
            return self
        if port.isnumeric():
            self.add_unique_lconfig('ports', port)
            return self
        # port is envvar
        new_port = '{{ .GetEnv "' + port + '" }}'
        if 'lconfig' in self._dict and 'ports' in self._dict['lconfig']:
            existing_ports = self._dict['lconfig']['ports']
            for existing_port in existing_ports:
                if new_port in existing_port:
                    return self
        self.add_unique_lconfig('ports', new_port)
        return self

    def add_lconfig_ports(self, ports: List[str]) -> Task:
        for port in ports:
            self.add_lconfig_port(port)
        return self

    def get_location(self) -> str:
        if 'location' in self._dict:
            return self._dict['location']
        return ''

    def get_possible_ports(self) -> List[str]:
        env_dict = {}
        for env_key, env in self.get_all_env().items():
            env_dict[env_key] = env.get_default()
        possible_ports_env = get_service_possible_ports_env(env_dict)
        return list(possible_ports_env.keys())

    def as_dict(self) -> TaskDict:
        return self._dict
