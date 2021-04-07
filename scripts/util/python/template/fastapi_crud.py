from typing import List

script_template = '''
# List {entity}
def list_{entity}(db: Session, skip: int = 0, limit: int = 100):
    return db.query(model.{entity_class}).offset(skip).limit(limit).all()


# Get {entity}
def get_{entity}(db: Session, {entity}_id: int):
    return db.query(model.{entity_class}).filter(model.{entity_class}.id == {entity}_id).first()


# Create {entity}
def create_{entity}(db: Session, {entity}_data: schema.{entity_class}Create):
    db_{entity} = model.{entity_class}({init_property})
    if db_{entity} is None:
        raise Error('Cannot create {entity}')
    db.add(db_{entity})
    db.commit()
    db.refresh(db_{entity})
    return db_{entity}


# Update {entity}
def update_{entity}(db: Session, {entity}_id: int, {entity}_data: schema.{entity_class}Update):
    db_{entity} = get_{entity}(db, {entity}_id)
    if db_{entity} is None:
        return None
    {update_property}
    db.add(db_{entity})
    db.commit()
    db.refresh(db_{entity})
    return db_{entity}


# Delete {entity}
def delete_{entity}(db: Session, {entity}_id: int):
    db_{entity} = get_{entity}(db, {entity}_id)
    if db_{entity} is None:
        return None
    db.delete(db_{entity})
    db.commit()
    return db_{entity}
'''

def get_script(entity_class: str, entity: str, fields: List[str]) -> str:
    update_property = '\n    '.join([
        'db_{entity}.{field} = {entity}_data.{field}'.format(entity = entity, field = field)
        for field in fields 
    ])
    init_property = ', '.join([
        '{field} = {entity}_data.{field}'.format(entity = entity, field = field)
        for field in fields
    ])
    return script_template.format(
        entity_class=entity_class,
        entity=entity,
        init_property=init_property,
        update_property=update_property
    )