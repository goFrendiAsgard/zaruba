from typing import List, Mapping
from helper import cli
from helper.project import ServiceProject

import helper.generator as generator
import helper.task as task
import template.additional_service_task as additional_service_task

import os, sys, traceback

@cli
def create_service_task(template_location: str, service_name: str, image_name: str, container_name: str, location: str, start_command: str, ports_str: str):
    ports = ports_str.split(',')
    dir_name = '.'
    service_project = ServiceProject()
    service_project.load_from_template(template_location)
    service_project.generate(dir_name=dir_name, service_name=service_name, image_name=image_name, container_name=container_name, location=location, start_command=start_command, ports=ports)


if __name__ == '__main__':
    create_service_task()