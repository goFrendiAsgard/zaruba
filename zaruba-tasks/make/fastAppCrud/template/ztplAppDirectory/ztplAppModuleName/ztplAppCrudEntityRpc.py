from typing import Any, Optional, Mapping
from helpers.transport import RPC
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from schemas.user import User
from repos.ztplAppCrudEntity import ZtplAppCrudEntityRepo
from ztplAppModuleName.ztplAppCrudEntityService import ZtplAppCrudEntityService

def register_ztpl_app_crud_entity_rpc(rpc: RPC, ztpl_app_crud_entity_repo: ZtplAppCrudEntityRepo):

    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(ztpl_app_crud_entity_repo)

    @rpc.handle('find_ztpl_app_crud_entity')
    def find_ztpl_app_crud_entities(keyword: str, limit: int, offset: int) -> Mapping[str, Any]:
        ztpl_app_crud_result = ztpl_app_crud_entity_service.find(keyword, limit, offset)
        return ztpl_app_crud_result.dict()

    @rpc.handle('find_ztpl_app_crud_entity_by_id')
    def find_ztpl_app_crud_entity_by_id(id: str) -> Optional[Mapping[str, Any]]:
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(id)
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()

    @rpc.handle('insert_ztpl_app_crud_entity')
    def insert_ztpl_app_crud_entity(ztpl_app_crud_entity_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ZtplAppCrudEntityData.parse_obj(ztpl_app_crud_entity_data) 
        ztpl_app_crud_entity.created_by = current_user.id
        new_ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(ztpl_app_crud_entity)
        return None if new_ztpl_app_crud_entity is None else new_ztpl_app_crud_entity.dict()

    @rpc.handle('update_ztpl_app_crud_entity')
    def update_ztpl_app_crud_entity(id: str, ztpl_app_crud_entity_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ZtplAppCrudEntityData.parse_obj(ztpl_app_crud_entity_data) 
        ztpl_app_crud_entity.updated_by = current_user.id
        updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update(id, ztpl_app_crud_entity)
        return None if updated_ztpl_app_crud_entity is None else updated_ztpl_app_crud_entity.dict()

    @rpc.handle('delete_ztpl_app_crud_entity')
    def delete_ztpl_app_crud_entity(id: str, current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete(id)
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()

    print('Handle RPC for ztplAppModuleName.ZtplAppCrudEntity')