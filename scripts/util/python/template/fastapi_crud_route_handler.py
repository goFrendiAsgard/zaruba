script_template = '''
    # List {entity} route
    @app.get('/{entity}s/', response_model=List[schema.{entity_caption}])
    def crud_list_{entity}(skip: int = 0, limit: int = 100):
        try:
            db_{entity}_list = mb.call_rpc('list_{entity}', skip, limit)
            return [schema.{entity_caption}.parse_obj(db_{entity}) for db_{entity} in db_{entity}_list]
        except Exception:
            print(traceback.format_exc())
            raise HTTPException(status_code=500, detail='Internal server error')


    # Get {entity} route
    @app.get('/{entity}s/{{{entity}_id}}', response_model=schema.{entity_caption})
    def crud_get_{entity}({entity}_id: int):
        try:
            db_{entity} = mb.call_rpc('get_{entity}', {entity}_id)
            if db_{entity} is None:
                raise HTTPException(status_code=404, detail='{entity_caption} not found')
            return schema.{entity_caption}.parse_obj(db_{entity})
        except Exception:
            print(traceback.format_exc())
            raise HTTPException(status_code=500, detail='Internal server error')


    # Create {entity} route
    @app.post('/{entity}s/', response_model=schema.{entity_caption})
    def crud_create_{entity}({entity}_data: schema.{entity_caption}Create):
        try:
            db_{entity} = mb.call_rpc('create_{entity}', {entity}_data.dict())
            if db_{entity} is None:
                raise HTTPException(status_code=404, detail='{entity_caption} not created')
            return schema.{entity_caption}.parse_obj(db_{entity})
        except Exception:
            print(traceback.format_exc())
            raise HTTPException(status_code=500, detail='Internal server error')


    # Update {entity} route
    @app.put('/{entity}s/{{{entity}_id}}', response_model=schema.{entity_caption})
    def crud_update_{entity}({entity}_id: int, {entity}_data: schema.{entity_caption}Update):
        try:
            db_{entity} = mb.call_rpc('update_{entity}', {entity}_id, {entity}_data.dict())
            if db_{entity} is None:
                raise HTTPException(status_code=404, detail='{entity_caption} not found')
            return schema.{entity_caption}.parse_obj(db_{entity})
        except Exception:
            print(traceback.format_exc())
            raise HTTPException(status_code=500, detail='Internal server error')


    # Delete {entity} route
    @app.delete('/{entity}s/{{{entity}_id}}', response_model=schema.{entity_caption})
    def crud_get_{entity}({entity}_id: int):
        try:
            db_{entity} = mb.call_rpc('delete_{entity}', {entity}_id)
            if db_{entity} is None:
                raise HTTPException(status_code=404, detail='{entity_caption} not found')
            return schema.{entity_caption}.parse_obj(db_{entity})
        except Exception:
            print(traceback.format_exc())
            raise HTTPException(status_code=500, detail='Internal server error')
'''

def get_script(entity: str, entity_caption: str) -> str:
    return script_template.format(entity=entity, entity_caption=entity_caption)