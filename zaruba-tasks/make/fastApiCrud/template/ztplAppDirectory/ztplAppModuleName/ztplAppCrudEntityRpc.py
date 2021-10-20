from typing import Any, List, Mapping
from helpers.transport import RPC
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from repos.ztplAppCrudEntity import ZtplAppCrudEntityRepo

def register_ztpl_app_crud_entity_rpc(rpc: RPC, ztpl_app_crud_entity_repo: ZtplAppCrudEntityRepo):

    @rpc.handle('find_ztpl_app_crud_entity')
    def find_ztpl_app_crud_entity(keyword: str, limit: int, offset: int) -> List[Mapping[str, Any]]:
        results = ztpl_app_crud_entity_repo.find(keyword, limit, offset)
        return [result.dict() for result in results]


    @rpc.handle('find_ztpl_app_crud_entity_by_id')
    def find_ztpl_app_crud_entity_by_id(id: str) -> Mapping[str, Any]:
        result = ztpl_app_crud_entity_repo.find_by_id(id)
        return None if result is None else result.dict()


    @rpc.handle('insert_ztpl_app_crud_entity')
    def insert_ztpl_app_crud_entity(data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = ztpl_app_crud_entity_repo.insert(ZtplAppCrudEntityData.parse_obj(data))
        return None if result is None else result.dict()


    @rpc.handle('update_ztpl_app_crud_entity')
    def update_ztpl_app_crud_entity(id: str, data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = ztpl_app_crud_entity_repo.update(id, ZtplAppCrudEntityData.parse_obj(data))
        return None if result is None else result.dict()


    @rpc.handle('delete_ztpl_app_crud_entity')
    def delete_ztpl_app_crud_entity(id: str) -> Mapping[str, Any]:
        result = ztpl_app_crud_entity_repo.delete(id)
        return None if result is None else result.dict()
    

    print('Handle RPC for ztplAppModuleName.ZtplAppCrudEntity')
