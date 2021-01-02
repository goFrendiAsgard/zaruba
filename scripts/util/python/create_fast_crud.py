from typing import List
from common_helper import get_argv

import os, re, sys, traceback

# USAGE
# python create_fast_crud <location> <module> <entity> <fields>

route_template = '''

    @app.get('/{entity}s/', response_model=List[schema.{entity_caption}])
    def crud_list_{entity}(skip: int = 0, limit: int = 100):
        db_{entity}_list = mb.call_rpc('list_{entity}', skip, limit)
        return [schema.{entity_caption}.parse_obj(db_{entity}) for db_{entity} in db_{entity}_list]

    @app.get('/{entity}s/{{{entity}_id}}', response_model=schema.{entity_caption})
    def crud_get_{entity}({entity}_id: int):
        db_{entity} = mb.call_rpc('get_{entity}', {entity}_id)
        if db_{entity} is None:
            raise HTTPException(status_code=404, detail='{entity_caption} not found')
        return schema.{entity_caption}.parse_obj(db_{entity})

    @app.post('/{entity}s/', response_model=schema.{entity_caption})
    def crud_create_{entity}({entity}_data: schema.{entity_caption}Create):
        db_{entity} = mb.call_rpc('create_{entity}', {entity}_data.dict())
        if db_{entity} is None:
            raise HTTPException(status_code=404, detail='{entity_caption} not found')
        return schema.{entity_caption}.parse_obj(db_{entity})

    @app.put('/{entity}s/{{{entity}_id}}', response_model=schema.{entity_caption})
    def crud_create_{entity}({entity}_id: int, {entity}_data: schema.{entity_caption}Create):
        db_{entity} = mb.call_rpc('update_{entity}', {entity}_id, {entity}_data.dict())
        if db_{entity} is None:
            raise HTTPException(status_code=404, detail='{entity_caption} not found')
        return schema.{entity_caption}.parse_obj(db_{entity})

    @app.delete('/{entity}s/{{{entity}_id}}', response_model=schema.{entity_caption})
    def crud_get_{entity}({entity}_id: int):
        db_{entity} = mb.call_rpc('delete_{entity}', {entity}_id)
        if db_{entity} is None:
            raise HTTPException(status_code=404, detail='{entity_caption} not found')
        return schema.{entity_caption}.parse_obj(db_{entity})

'''

event_template = '''

    @transport.handle_rpc(mb, 'list_{entity}')
    @database.handle(DBSession)
    def crud_rpc_list_{entity}(db: Session, skip: int = 0, limit: int = 100) -> List[Mapping[str, Any]]:
        db_{entity}_list = crud.list_{entity}(db, skip = skip, limit = limit)
        return [schema.{entity_class}.from_orm(db_{entity}).dict() for db_{entity} in db_{entity}_list]

    @transport.handle_rpc(mb, 'get_{entity}')
    @database.handle(DBSession)
    def crud_rpc_get_{entity}(db: Session, {entity}_id: int) -> Mapping[str, Any]:
        db_{entity} = crud.get_{entity}(db, {entity}_id = {entity}_id)
        if db_{entity} is None:
            return None
        return schema.{entity_class}.from_orm(db_{entity}).dict()

    @transport.handle_rpc(mb, 'create_{entity}')
    @database.handle(DBSession)
    def crud_rpc_create_{entity}(db: Session, {entity}_dict: Mapping[str, Any]) -> Mapping[str, Any]:
        db_{entity} = crud.create_{entity}(db, {entity}_data = schema.{entity_class}Create.parse_obj({entity}_dict))
        if db_{entity} is None:
            return None
        return schema.{entity_class}.from_orm(db_{entity}).dict()

    @transport.handle_rpc(mb, 'update_{entity}')
    @database.handle(DBSession)
    def crud_rpc_update_{entity}(db: Session, {entity}_id: int, {entity}_dict: Mapping[str, Any]) -> Mapping[str, Any]:
        db_{entity} = crud.update_{entity}(db, {entity}_id = {entity}_id, {entity}_data = schema.{entity_class}Update.parse_obj({entity}_dict))
        if db_{entity} is None:
            return None
        return schema.{entity_class}.from_orm(db_{entity}).dict()

    @transport.handle_rpc(mb, 'delete_{entity}')
    @database.handle(DBSession)
    def crud_rpc_delete_{entity}(db: Session, {entity}_id: int) -> Mapping[str, Any]:
        db_{entity} = crud.delete_{entity}(db, {entity}_id = {entity}_id)
        if db_{entity} is None:
            return None
        return schema.{entity_class}.from_orm(db_{entity}).dict()

'''

crud_template = '''

def list_{entity}(db: Session, skip: int = 0, limit: int = 100):
    return db.query(model.{entity_class}).offset(skip).limit(limit).all()

def get_{entity}(db: Session, {entity}_id: int):
    return db.query(model.{entity_class}).filter(model.{entity_class}.id == {entity}_id).first()

def create_{entity}(db: Session, {entity}_data: schema.{entity_class}Create):
    db_{entity} = model.{entity_class}({init_property})
    if db_{entity} is None:
        return None
    db.add(db_{entity})
    db.commit()
    db.refresh(db_{entity})
    return db_{entity}

def update_{entity}(db: Session, {entity}_id: int, {entity}_data: schema.{entity_class}Update):
    db_{entity} = get_{entity}(db, {entity}_id)
    if db_{entity} is None:
        return None
    {update_property}
    db.add(db_{entity})
    db.commit()
    db.refresh(db_{entity})
    return db_{entity}

def delete_{entity}(db: Session, {entity}_id: int):
    db_{entity} = get_{entity}(db, {entity}_id)
    if db_{entity} is None:
        return None
    db.delete(db_{entity})
    db.commit()
    return db_{entity}

'''

model_template = '''

class {entity_class}(Base):
    __tablename__ = '{entity}s'
    id = Column(Integer, primary_key=True, index=True)
    {model_field_declaration}

'''

schema_template = '''

class {entity_class}Base(BaseModel):
    {schema_field_declaration}

class {entity_class}Create({entity_class}Base):
    pass

class {entity_class}Update({entity_class}Base):
    pass

class {entity_class}({entity_class}Base):
    id: int
    class Config:
        orm_mode = True

'''

def create_fast_crud(location: str, module: str, entity: str, fields: List[str]):
    # declare substitutions
    indentation = '    '
    indented_new_line = '\n' + indentation
    entity = re.sub(r'[^A-Za-z0-9_]+', '_', entity).lower()
    entity_class = entity.capitalize()
    entity_caption = entity.replace('_', ' ').capitalize()
    model_field_declaration = indented_new_line.join([
        '{field} = Column(String)'.format(field = field) 
        for field in fields 
    ])
    schema_field_declaration = 'pass'
    if len(fields) > 0:
        schema_field_declaration = indented_new_line.join([
            '{field} : str'.format(field = field) 
            for field in fields 
        ])
    update_property = indented_new_line.join([
        'db_{entity}.{field} = {entity}_data.{field}'.format(entity = entity, field = field)
        for field in fields 
    ])
    init_property = ', '.join([
        '{field} = {entity}_data.{field}'.format(entity = entity, field = field)
        for field in fields
    ])
    # create files
    create_schema(location, module, entity_class, schema_field_declaration)
    create_model(location, module, entity_class, entity, model_field_declaration)
    create_crud(location, module, entity_class, entity, init_property, update_property)
    create_route(location, module, entity, entity_caption)
    create_event(location, module, entity_class, entity)


def create_event(location: str, module: str, entity_class: str, entity: str):
    file_name = os.path.abspath(os.path.join(location, module, 'event.py'))
    f_read = open(file_name, 'r')
    lines = f_read.readlines()
    f_read.close()
    # look for line with 'def init(' prefix
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            insert_index = index + 1
            break
    if insert_index == -1:
        raise Exception('init function not found in {}'.format(file_name))
    # add event handler
    lines.insert(insert_index, event_template.format(
        entity_class=entity_class,
        entity=entity
    ))
    f_write = open(file_name, 'w')
    f_write.writelines(lines)
    f_write.close()


def create_route(location: str, module: str, entity: str, entity_caption: str):
    file_name = os.path.abspath(os.path.join(location, module, 'route.py'))
    f_read = open(file_name, 'r')
    lines = f_read.readlines()
    f_read.close()
    # look for line with 'def init(' prefix
    insert_index = -1
    for index, line in enumerate(lines):
        if line.startswith('def init('):
            insert_index = index + 1
            break
    if insert_index == -1:
        raise Exception('init function not found in {}'.format(file_name))
    lines.insert(insert_index, route_template.format(
        entity=entity,
        entity_caption=entity_caption
    ))
    f_write = open(file_name, 'w')
    f_write.writelines(lines)
    f_write.close()


def create_schema(location: str, module: str, entity_class: str, schema_field_declaration=str):
    # create schema
    file_name = os.path.abspath(os.path.join(location, module, 'schema.py'))
    f_write = open(file_name, 'a')
    f_write.write(schema_template.format(
        entity_class=entity_class, 
        schema_field_declaration=schema_field_declaration
    ))
    f_write.close()


def create_model(location: str, module: str, entity_class: str, entity: str, model_field_declaration: str):
    file_name = os.path.abspath(os.path.join(location, module, 'model.py'))
    f_write = open(file_name, 'a')
    f_write.write(model_template.format(
        entity_class=entity_class,
        entity=entity,
        model_field_declaration=model_field_declaration
    ))
    f_write.close()


def create_crud(location: str, module: str, entity_class: str, entity: str, init_property: str, update_property: str):
    file_name = os.path.abspath(os.path.join(location, module, 'crud.py'))
    f_write = open(file_name, 'a')
    f_write.write(crud_template.format(
        entity_class=entity_class,
        entity=entity,
        init_property=init_property,
        update_property=update_property
    ))
    f_write.close()


if __name__ == '__main__':
    location = get_argv(1)
    module = get_argv(2)
    entity = get_argv(3)
    str_fields = get_argv(4)
    fields = str_fields.split(',') if str_fields != '' else []
    try:
        create_fast_crud(location, module, entity, fields)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
