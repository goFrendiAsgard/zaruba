from typing import List

script_template = '''
# {entity} model
class {entity_class}(Base):
    __tablename__ = '{entity}s'
    id = Column(Integer, primary_key=True, index=True)
    {model_field_declaration}
'''

def get_script(entity_class: str, entity: str, fields: List[str]) -> str:
    model_field_declaration = '\n    '.join([
        '{field} = Column(String)'.format(field = field) 
        for field in fields 
    ])
    return script_template.format(
        entity_class=entity_class,
        entity=entity,
        model_field_declaration=model_field_declaration
    )