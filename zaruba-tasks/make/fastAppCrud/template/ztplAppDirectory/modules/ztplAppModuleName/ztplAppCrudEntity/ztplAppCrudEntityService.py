from typing import Optional
from helpers.transport import RPC, MessageBus
from schemas.user import User
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData, ZtplAppCrudEntityResult
from modules.ztplAppModuleName.ztplAppCrudEntity.repos.ztplAppCrudEntityRepo import ZtplAppCrudEntityRepo
from fastapi import HTTPException

class ZtplAppCrudEntityService():

    def __init__(self, mb: MessageBus, rpc: RPC, ztpl_app_crud_entity_repo: ZtplAppCrudEntityRepo):
        self.mb = mb
        self.rpc = rpc
        self.ztpl_app_crud_entity_repo = ztpl_app_crud_entity_repo


    def find(self, keyword: str, limit: int, offset: int, current_user: Optional[User] = None) -> ZtplAppCrudEntityResult:
        count = self.ztpl_app_crud_entity_repo.count(keyword)
        rows = self.ztpl_app_crud_entity_repo.find(keyword, limit, offset)
        return ZtplAppCrudEntityResult(count=count, rows=rows)


    def find_by_id(self, id: str, current_user: Optional[User] = None) -> Optional[ZtplAppCrudEntity]:
        ztpl_app_crud_entity = self._find_by_id_or_error(id, current_user)
        return ztpl_app_crud_entity


    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user: User) -> Optional[ZtplAppCrudEntity]:
        ztpl_app_crud_entity_data.created_by = current_user.id
        ztpl_app_crud_entity_data.updated_by = current_user.id
        ztpl_app_crud_entity_data = self._validate_data(ztpl_app_crud_entity_data)
        return self.ztpl_app_crud_entity_repo.insert(ztpl_app_crud_entity_data)


    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData, current_user: User) -> Optional[ZtplAppCrudEntity]:
        self._find_by_id_or_error(id, current_user)
        ztpl_app_crud_entity_data.updated_by = current_user.id
        ztpl_app_crud_entity_data = self._validate_data(ztpl_app_crud_entity_data, id)
        return self.ztpl_app_crud_entity_repo.update(id, ztpl_app_crud_entity_data)


    def delete(self, id: str, current_user: User) -> Optional[ZtplAppCrudEntity]:
        self._find_by_id_or_error(id, current_user)
        return self.ztpl_app_crud_entity_repo.delete(id)


    def _find_by_id_or_error(self, id: Optional[str] = None, current_user: Optional[User] = None) -> Optional[ZtplAppCrudEntity]:
        ztpl_app_crud_entity = self.ztpl_app_crud_entity_repo.find_by_id(id)
        if ztpl_app_crud_entity is None:
            raise HTTPException(
                status_code=404, 
                detail='ZtplAppCrudEntity id not found: {}'.format(id)
            )
        return ztpl_app_crud_entity


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
