from typing import Any, List, Mapping
from helpers.transport import RPC
from schemas.zarubaEntityName import ZarubaEntityName, ZarubaEntityNameData
from repos.zarubaEntityName import ZarubaEntityNameRepo

def handle(rpc: RPC, zaruba_entity_name_repo: ZarubaEntityNameRepo):

    @rpc.handle('find_zaruba_entity_name')
    def find_zaruba_entity_name(keyword: str, limit: int, offset: int) -> List[Mapping[str, Any]]:
        results = zaruba_entity_name_repo.find(keyword, limit, offset)
        return [result.dict() for result in results]


    @rpc.handle('find_zaruba_entity_name_by_id')
    def find_zaruba_entity_name_by_id(id: str) -> Mapping[str, Any]:
        result = zaruba_entity_name_repo.find_by_id(id)
        return None if result is None else result.dict()


    @rpc.handle('insert_zaruba_entity_name')
    def insert_zaruba_entity_name(data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = zaruba_entity_name_repo.insert(ZarubaEntityNameData.parse_obj(data))
        return None if result is None else result.dict()


    @rpc.handle('update_zaruba_entity_name')
    def update_zaruba_entity_name(id: str, data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = zaruba_entity_name_repo.update(id, ZarubaEntityNameData.parse_obj(data))
        return None if result is None else result.dict()


    @rpc.handle('delete_zaruba_entity_name')
    def delete_zaruba_entity_name(id: str) -> Mapping[str, Any]:
        result = zaruba_entity_name_repo.delete(id)
        return None if result is None else result.dict()
    

    print('Handle RPC for zarubaModuleName.zarubaEntityName')

