from typing import Any, List, Mapping
from helpers.transport import RPC
from schemas.ztplAppCrudEntity import ZtplAppCrudEntityData
from schemas.user import User
from repos.ztplAppCrudEntity import ZtplAppCrudEntityRepo
from ztplAppModuleName.ztplAppCrudEntityService import ZtplAppCrudEntityService

def register_ztpl_app_crud_entity_rpc(rpc: RPC, ztpl_app_crud_entity_repo: ZtplAppCrudEntityRepo):

    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(ztpl_app_crud_entity_repo)

    @rpc.handle('find_ztpl_app_crud_entity')
    def find_ztpl_app_crud_entity(keyword: str, limit: int, offset: int, current_user_data: Mapping[str, Any]) -> List[Mapping[str, Any]]:
        ztpl_app_crud_entities = ztpl_app_crud_entity_service.find(keyword, limit, offset)
        return [ztpl_app_crud_entity.dict() for ztpl_app_crud_entity in ztpl_app_crud_entities]

    @rpc.handle('find_ztpl_app_crud_entity_by_id')
    def find_ztpl_app_crud_entity_by_id(id: str, current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(id)
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()

    @rpc.handle('insert_ztpl_app_crud_entity')
    def insert_ztpl_app_crud_entity(ztpl_app_crud_entity_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(ZtplAppCrudEntityData.parse_obj(ztpl_app_crud_entity_data))
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()

    @rpc.handle('update_ztpl_app_crud_entity')
    def update_ztpl_app_crud_entity(id: str, ztpl_app_crud_entity_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.update(id, ZtplAppCrudEntityData.parse_obj(ztpl_app_crud_entity_data))
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()

    @rpc.handle('delete_ztpl_app_crud_entity')
    def delete_ztpl_app_crud_entity(id: str, current_user_data: Mapping[str, Any]) -> Mapping[str, Any]:
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete(id)
        return None if ztpl_app_crud_entity is None else ztpl_app_crud_entity.dict()

    print('Handle RPC for ztplAppModuleName.ZtplAppCrudEntity')