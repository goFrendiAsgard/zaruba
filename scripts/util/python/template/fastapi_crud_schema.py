from typing import List

script_template = '''
# {entity} schema

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

def get_script(entity_class: str, entity: str, fields: List[str]) -> str:
    schema_field_declaration = 'pass'
    if len(fields) > 0:
        schema_field_declaration = '\n    '.join([
            '{field} : str'.format(field = field) 
            for field in fields 
        ])
    return script_template.format(
        entity_class=entity_class, 
        entity=entity,
        schema_field_declaration=schema_field_declaration
    )