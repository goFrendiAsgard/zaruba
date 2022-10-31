from typing import Any, Optional, Mapping
from helpers.transport import RPC, MessageBus
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from schemas.user import User
from modules.ztplAppModuleName.ztplAppCrudEntity.ztplAppCrudEntityService import ZtplAppCrudEntityService

import sys

def register_ztpl_app_crud_entity_rpc(mb: MessageBus, rpc: RPC, ztpl_app_crud_entity_service: ZtplAppCrudEntityService):

    @rpc.handle('find_ztpl_app_crud_entity')
    def find_ztpl_app_crud_entities(keyword: str, limit: int, offset: int, current_user_data: Optional[Mapping[str, Any]]) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword, limit, offset, current_user)
        return ztpl_app_crud_entity_result.dict()


    @rpc.handle('find_ztpl_app_crud_entity_by_id')
    def find_ztpl_app_crud_entity_by_id(id: str, current_user_data: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(id, current_user)
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()


    @rpc.handle('insert_ztpl_app_crud_entity')
    def insert_ztpl_app_crud_entity(ztpl_app_crud_entity_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ZtplAppCrudEntityData.parse_obj(ztpl_app_crud_entity_data) 
        new_ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(ztpl_app_crud_entity, current_user)
        return None if new_ztpl_app_crud_entity is None else new_ztpl_app_crud_entity.dict()


    @rpc.handle('update_ztpl_app_crud_entity')
    def update_ztpl_app_crud_entity(id: str, ztpl_app_crud_entity_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ZtplAppCrudEntityData.parse_obj(ztpl_app_crud_entity_data) 
        ztpl_app_crud_entity.updated_by = current_user.id
        updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update(id, ztpl_app_crud_entity, current_user)
        return None if updated_ztpl_app_crud_entity is None else updated_ztpl_app_crud_entity.dict()


    @rpc.handle('delete_ztpl_app_crud_entity')
    def delete_ztpl_app_crud_entity(id: str, current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete(id, current_user)
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()


    print('Handle RPC for ztplAppModuleName.ZtplAppCrudEntity', file=sys.stderr)