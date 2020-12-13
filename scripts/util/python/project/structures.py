from __future__ import annotations
from typing import List, Mapping, TypedDict


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



class ProjectDict(TypedDict, total=False):
    includes: List[str]
    tasks: Mapping[str, TaskDict]



class Env():

    def __init__(self, env_dict: EnvDict = {}):
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
        if default:
            self._dict['default'] = str(default)
        elif 'default' in self._dict:
            del self._dict['default']
        return self

    
    def get_default(self) -> str:
        if 'default' in self._dict:
            return str(self._dict['default'])
        return ''  


    def as_dict(self) -> EnvDict:
        return self._dict



class Task():

    def __init__(self, task_dict: TaskDict = {}):
        self._dict: TaskDict = task_dict


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


    def set_env(self, key: str, env: Env, override: bool=True) -> Task:
        if override or ('env' not in self._dict) or (key not in self._dict['env']):
            if 'env' not in self._dict:
                self._dict['env'] = {}
            self._dict['env'][key] = env.as_dict()
        return self

    
    def get_all_env(self) -> Mapping[str, Env]:
        env_map: Mapping[str, Env] = {}
        if 'env' in self._dict:
            for key in self._dict['env']:
                env_map[key] = self.get_env(key)
        return env_map
    

    def get_env(self, key) -> Env:
        if 'env' in self._dict and key in self._dict['env']:
            return Env(self._dict['env'][key])
        return Env()

    
    def set_config(self, key: str, val: str) -> Task:
        if 'config'not in self._dict:
            self._dict['config'] = {}
        self._dict['config'][key] = val
        return self

    
    def init_lconfig(self, key: str) -> Task:
        if 'lconfig'not in self._dict:
            self._dict['lconfig'] = {}
        if key not in self._dict['lconfig']:
            self._dict['lconfig'][key] = []

 
    def add_lconfig(self, key: str, *vals: str) -> Task:
        self.init_lconfig(key)
        for val in vals:
            self._dict['lconfig'][key].append(val)
        return self

 
    def add_unique_lconfig(self, key: str, *vals: str) -> Task:
        self.init_lconfig(key)
        for val in vals:
            if val in self._dict['lconfig'][key]:
                continue
            self._dict['lconfig'][key].append(val)
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


    def add_lconfig_ports(self, *ports: str) -> Task:
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



class Template():

    def __init__(self, location: str, is_default: bool):
        self.location = location
        self.is_default = is_default

