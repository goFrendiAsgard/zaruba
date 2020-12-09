from typing import List, Mapping, TypedDict
from project.structures import Template
from project.helpers import get_dict_from_file, replace_file_content, save_dict_to_file, get_service_name_by_location
import os
import shutil


class LinkDict(TypedDict, total=False):
    source: str
    destination: str


class Link():
    def __init__(self, source: str, destination: str):
        self.source = source
        self.destination = destination


class ServiceGen():

    def __init__(self, template: Template, target_location: str):
        config = get_dict_from_file(os.path.join(template.location, 'config.yaml'))
        self.target_location = target_location
        self.service_location = os.path.join(template.location, config.get('service', './service'))
        self.module_location = os.path.join(template.location, config.get('module', './module'))
        self.links: Mapping[str, Link] = {}
        links: Mapping[str, LinkDict] = config.get('links', {})
        for location, link in links.items():
            location = os.path.join(template.location, location)
            self.links[location] = Link(
                link.get('source', ''),
                link.get('destination', '')
            )
        self.replace_dict = self.create_replace_dict()
        self.ports: List[str] = []
        ports: List[str] = config.get('ports', [])
        for port in ports:
            for key, val in self.replace_dict.items():
                self.ports.append(port.replace(key, val))
    

    def generate_service(self):
        if os.path.isdir(self.target_location):
            raise Exception('{} is not empty'.format(self.target_location))
        self.create_service()
        self.create_link()

    
    def create_service(self):
        shutil.copytree(self.service_location, self.target_location)
        for root, dir_names, file_names in os.walk(os.path.abspath(self.target_location)):
            for file_name in file_names:
                file_name = os.path.join(root, file_name)
                replace_file_content(file_name, self.replace_dict)


    def create_replace_dict(self) -> Mapping[str, str]:
        replace_dict = {
            'SERVICE': get_service_name_by_location(self.target_location).upper(),
        }
        return replace_dict
            
    
    def create_link(self):
        kwargs = get_dict_from_file('default.kwargs.yaml')
        for location, link in self.links.items():
            source = os.path.abspath(os.path.join('.', link.source))
            destination = os.path.abspath(os.path.join(self.target_location, link.destination))
            shutil.copytree(location, source)
            shutil.copytree(location, destination)
            for root, dir_names, file_names in os.walk(destination):
                for file_name in file_names:
                    os.chmod(os.path.join(root, file_name), 0o555)
                for dir_name in dir_names:
                    os.chmod(os.path.join(root, dir_name), 0o555)
                os.chmod(root, 0o555)
            kwargs['link::{}'.format(destination)] = source
        save_dict_to_file('default.kwargs.yaml', kwargs)

