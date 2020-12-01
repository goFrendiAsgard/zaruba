from __future__ import annotations
from typing import List, Mapping, Any, Type, TypedDict


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
            self._dict['default'] = default
        elif 'default' in self._dict:
            del self._dict['default']
        return self

    
    def get_default(self) -> str:
        if 'default' in self._dict:
            return self._dict['default']
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


    def set_env(self, key: str, env: Env) -> Task:
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

 
    def add_lconfig(self, key: str, *vals: str) -> Task:
        if 'lconfig'not in self._dict:
            self._dict['lconfig'] = {}
        if key not in self._dict['lconfig']:
            self._dict['lconfig'][key] = []
        for val in vals:
            self._dict['lconfig'][key].append(val)
        return self

    
    def as_dict(self) -> TaskDict:
        return self._dict
