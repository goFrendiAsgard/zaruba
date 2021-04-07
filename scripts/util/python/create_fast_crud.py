from typing import List
from helper import cli
import helper.generator as generator
import template.fastapi_crud_event_handler as fastapi_crud_event_handler
import template.fastapi_crud_route_handler as fastapi_crud_route_handler
import template.fastapi_crud_model as fastapi_crud_model
import template.fastapi_crud_schema as fastapi_crud_schema
import template.fastapi_crud as fastapi_crud

import os, re

@cli
def create_fast_crud(location: str, module: str, entity: str, str_fields: str):
    fields = str_fields.split(',') if str_fields != '' else []
    # declare substitutions
    entity = re.sub(r'[^A-Za-z0-9_]+', '_', entity).lower()
    entity_class = entity.capitalize()
    entity_caption = entity.replace('_', ' ').capitalize()
    # create files
    create_schema(location, module, entity_class, entity, fields)
    create_model(location, module, entity_class, entity, fields)
    create_crud(location, module, entity_class, entity, fields)
    create_route(location, module, entity, entity_caption)
    create_event(location, module, entity_class, entity)


def create_event(location: str, module: str, entity_class: str, entity: str):
    file_name = os.path.abspath(os.path.join(location, module, 'event.py'))
    lines = generator.read_lines(file_name)
    # look for line with 'def init(' prefix
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            insert_index = index + 1
            break
    if insert_index == -1:
        raise Exception('init function not found in {}'.format(file_name))
    # add event handler
    lines.insert(
        insert_index, 
        '\n' + fastapi_crud_event_handler.get_script(
            entity_class=entity_class,
            entity=entity
        ) + '\n'
    )
    generator.write_lines(file_name, lines)


def create_route(location: str, module: str, entity: str, entity_caption: str):
    file_name = os.path.abspath(os.path.join(location, module, 'route.py'))
    lines = generator.read_lines(file_name)
    # look for line with 'def init(' prefix
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            insert_index = index + 1
            break
    if insert_index == -1:
        raise Exception('init function not found in {}'.format(file_name))
    lines.insert(
        insert_index,
        '\n' + fastapi_crud_route_handler.get_script(
            entity=entity,
            entity_caption = entity_caption
        ) + '\n'
    )
    generator.write_lines(file_name, lines)


def create_schema(location: str, module: str, entity_class: str, entity: str, fields=List[str]):
    # create schema
    file_name = os.path.abspath(os.path.join(location, module, 'schema.py'))
    text = generator.read_text(file_name)
    text += fastapi_crud_schema.get_script(
        entity_class=entity_class, 
        entity=entity,
        fields=fields
    )
    generator.write_text(file_name, text)


def create_model(location: str, module: str, entity_class: str, entity: str, fields: List[str]):
    file_name = os.path.abspath(os.path.join(location, module, 'model.py'))
    text = generator.read_text(file_name)
    text += '\n' + fastapi_crud_model.get_script(
        entity_class=entity_class,
        entity=entity,
        fields=fields
    ) + '\n'
    generator.write_text(file_name, text)


def create_crud(location: str, module: str, entity_class: str, entity: str, fields: List[str]):
    file_name = os.path.abspath(os.path.join(location, module, 'crud.py'))
    text = generator.read_text(file_name)
    text += '\n' + fastapi_crud.get_script(
        entity_class=entity_class,
        entity=entity,
        fields=fields
    ) + '\n'
    generator.write_text(file_name, text)


if __name__ == '__main__':
    create_fast_crud()