ztpl_app_crud_entity_repo = DBZtplAppCrudEntityRepo(engine=engine, create_all=db_create_all)
ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mb, rpc, ztpl_app_crud_entity_repo)