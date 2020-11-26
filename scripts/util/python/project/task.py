from __future__ import annotations
from typing import List, Mapping, TypedDict, Any

class Env():
    key: str
    default: str

class Task():

    def __init__(self):
        self.icon: str = ''
        self.extend: str = ''
        self.description: str = ''
        self.dependencies: List[str] = []
        self.env: Mapping[str, TypedDict('Env', {'from': str, 'default': str})] = {}
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
        self.env[key] = {
            'from': envvar,
            'default': default,
        }
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

    
    def to_dict(self) -> Mapping[str, Any]:
        dictionary = {}
        if self.icon:
            dictionary['icon'] = self.icon
        if self.extends:
            dictionary['extend'] = self.extend
        if self.description:
            dictionary['description'] = self.description
        if self.dependenceies:
            dictionary['dependencies'] = self.dependencies
        if self.env:
            dictionary['env'] = self.env
        if self.config:
            dictionary['config'] = self.config
        if self.lconfig:
            dictionary['lconfig'] = self.lconfig
        return dictionary

