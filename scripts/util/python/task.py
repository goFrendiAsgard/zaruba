from __future__ import annotations
from typing import List, Mapping, TypedDict
from dotenv import dotenv_values
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
        self.add_unique_lconfig(
            'ports', 
            '{open_template} .GetEnv "{env_name}" {close_template}'.format(
                open_template='{{', 
                close_template='}}', 
                env_name=port
            )
        )
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
        ports: List[str] = []
        for key, env in self.get_all_env().items():
            val = env.get_default()
            if val.isnumeric():
                if int(val) in [3306, 5432, 5672, 15672, 27017, 6379, 9200, 9300, 7001, 7199, 9042, 9160]:
                    continue
                if int(val) in [80, 443] or int(val) >= 3000:
                    ports.append(key)
        return ports


    def as_dict(self) -> TaskDict:
        return self._dict
