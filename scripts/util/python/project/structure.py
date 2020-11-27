from __future__ import annotations
from typing import List, Mapping, Any, Type, TypedDict


EnvDict = TypedDict('envDict', {'from': str, 'default': str}, total=False)


class TaskDict(TypedDict, total=False):
    icon: str
    extend: str
    description: str
    dependencies: List[str]
    env: EnvDict
    config: Mapping[str, str]
    lconfig: Mapping[str, List[str]]



class ProjectDict(TypedDict, total=False):
    includes: List[str]
    tasks: Mapping[str, TaskDict]



class Env():

    def __init__(self, envvar: str, default: str):
        self.envvar: str = envvar
        self.default: str = default
    

    def to_dict(self) -> EnvDict:
        dictionary: EnvDict = {}
        if self.envvar:
            dictionary['from'] = self.envvar
        if self.default:
            dictionary['default'] = self.default
        return dictionary



class Task():

    def __init__(self):
        self.icon: str = ''
        self.extend: str = ''
        self.description: str = ''
        self.dependencies: List[str] = []
        self.env: Mapping[str, Env] = {}
        self.config: Mapping[str, str] = {}
        self.lconfig: Mapping[str, List[str]] = {}


    def set_icon(self, icon: str) -> Task:
        self.icon = icon
        return self
    

    def set_parent_task(self, parent_task: str) -> Task:
        self.extend = parent_task
        return self

    
    def add_dependency(self, dependency: str) -> Task:
        if dependency not in self.dependencies:
            self.dependencies.append(dependency)
        return self 


    def set_env(self, key: str, envvar: str, default: str) -> Task:
        self.env[key] = Env(envvar, default)
        return self

    
    def set_config(self, key: str, val: str) -> Task:
        self.config[key] = val
        return self

 
    def add_lconfig(self, key: str, *vals: str) -> Task:
        if key not in self.lconfig:
            self.lconfig[key] = []
        for val in vals:
            self.lconfig[key].append(val)
        return self

    
    def to_dict(self) -> TaskDict:
        dictionary: TaskDict = {}
        if self.icon:
            dictionary['icon'] = self.icon
        if self.extend:
            dictionary['extend'] = self.extend
        if self.description:
            dictionary['description'] = self.description
        if self.dependencies:
            dictionary['dependencies'] = self.dependencies
        if self.env:
            dictionary['env'] = {key: env.to_dict() for key, env in self.env.items()}
        if self.config:
            dictionary['config'] = self.config
        if self.lconfig:
            dictionary['lconfig'] = self.lconfig
        return dictionary
