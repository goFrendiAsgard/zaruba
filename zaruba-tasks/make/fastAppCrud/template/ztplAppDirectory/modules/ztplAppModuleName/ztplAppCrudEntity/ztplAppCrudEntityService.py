from typing import Optional
from helpers.transport import RPC, MessageBus
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData, ZtplAppCrudEntityResult
from modules.ztplAppModuleName.ztplAppCrudEntity.repos.ztplAppCrudEntityRepo import ZtplAppCrudEntityRepo
from fastapi import HTTPException

class ZtplAppCrudEntityService():

    def __init__(self, mb: MessageBus, rpc: RPC, ztpl_app_crud_entity_repo: ZtplAppCrudEntityRepo):
        self.mb = mb
        self.rpc = rpc
        self.ztpl_app_crud_entity_repo = ztpl_app_crud_entity_repo

    def find(self, keyword: str, limit: int, offset: int) -> ZtplAppCrudEntityResult:
        count = self.ztpl_app_crud_entity_repo.count(keyword)
        rows = self.ztpl_app_crud_entity_repo.find(keyword, limit, offset)
        return ZtplAppCrudEntityResult(count=count, rows=rows)

    def find_by_id(self, id: str) -> Optional[ZtplAppCrudEntity]:
        return self.ztpl_app_crud_entity_repo.find_by_id(id)

    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        ztpl_app_crud_entity_data = self._validate_data(ztpl_app_crud_entity_data)
        return self.ztpl_app_crud_entity_repo.insert(ztpl_app_crud_entity_data)

    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        ztpl_app_crud_entity_data = self._validate_data(ztpl_app_crud_entity_data, id)
        return self.ztpl_app_crud_entity_repo.update(id, ztpl_app_crud_entity_data)

    def delete(self, id: str) -> Optional[ZtplAppCrudEntity]:
        return self.ztpl_app_crud_entity_repo.delete(id)

    def _validate_data(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData, id: Optional[str] = None) -> ZtplAppCrudEntityData:
        # TODO: add your custom logic
        # Example: checking duplication
        # if ztpl_app_crud_entity_data.some_field is not None:
        #     user = self.user_repo.find_by_some_field(ztpl_app_crud_entity_data.some_field)
        #     if user is not None and (id is None or user.id != id):
        #         raise HTTPException(
        #             status_code=422, 
        #             detail='some_field already exist: {}'.format(ztpl_app_crud_entity_data.some_field)
        #         )
        return ztpl_app_crud_entity_data
