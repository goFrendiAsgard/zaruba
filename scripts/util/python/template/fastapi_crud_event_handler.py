script_template = '''
    # List {entity} message handler
    @transport.handle_rpc(mb, 'list_{entity}')
    @database.handle(DBSession)
    def crud_rpc_list_{entity}(db: Session, skip: int = 0, limit: int = 100) -> List[Mapping[str, Any]]:
        try:
            db_{entity}_list = crud.list_{entity}(db, skip = skip, limit = limit)
            return [schema.{entity_class}.from_orm(db_{entity}).dict() for db_{entity} in db_{entity}_list]
        except Exception:
            print(traceback.format_exc())
            raise


    # Get {entity} message handler
    @transport.handle_rpc(mb, 'get_{entity}')
    @database.handle(DBSession)
    def crud_rpc_get_{entity}(db: Session, {entity}_id: int) -> Mapping[str, Any]:
        try:
            db_{entity} = crud.get_{entity}(db, {entity}_id = {entity}_id)
            if db_{entity} is None:
                return None
            return schema.{entity_class}.from_orm(db_{entity}).dict()
        except Exception:
            print(traceback.format_exc())
            raise


    # Create {entity} message handler
    @transport.handle_rpc(mb, 'create_{entity}')
    @database.handle(DBSession)
    def crud_rpc_create_{entity}(db: Session, {entity}_dict: Mapping[str, Any]) -> Mapping[str, Any]:
        try:
            db_{entity} = crud.create_{entity}(db, {entity}_data = schema.{entity_class}Create.parse_obj({entity}_dict))
            if db_{entity} is None:
                return None
            return schema.{entity_class}.from_orm(db_{entity}).dict()
        except Exception:
            print(traceback.format_exc())
            raise


    # Update {entity} message handler
    @transport.handle_rpc(mb, 'update_{entity}')
    @database.handle(DBSession)
    def crud_rpc_update_{entity}(db: Session, {entity}_id: int, {entity}_dict: Mapping[str, Any]) -> Mapping[str, Any]:
        try:
            db_{entity} = crud.update_{entity}(db, {entity}_id = {entity}_id, {entity}_data = schema.{entity_class}Update.parse_obj({entity}_dict))
            if db_{entity} is None:
                return None
            return schema.{entity_class}.from_orm(db_{entity}).dict()
        except Exception:
            print(traceback.format_exc())
            raise


    # Delete {entity} message handler
    @transport.handle_rpc(mb, 'delete_{entity}')
    @database.handle(DBSession)
    def crud_rpc_delete_{entity}(db: Session, {entity}_id: int) -> Mapping[str, Any]:
        try:
            db_{entity} = crud.delete_{entity}(db, {entity}_id = {entity}_id)
            if db_{entity} is None:
                return None
            return schema.{entity_class}.from_orm(db_{entity}).dict()
        except Exception:
            print(traceback.format_exc())
            raise
'''

def get_script(entity_class: str, entity: str) -> str:
    return script_template.format(entity_class=entity_class, entity=entity)