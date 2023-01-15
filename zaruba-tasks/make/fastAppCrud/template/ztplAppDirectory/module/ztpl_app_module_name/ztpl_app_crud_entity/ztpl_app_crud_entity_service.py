from typing import Optional
from transport import AppMessageBus, AppRPC
from schema.user import User
from schema.activity import ActivityData
from schema.ztpl_app_crud_entity import (
    ZtplAppCrudEntity, ZtplAppCrudEntityData, ZtplAppCrudEntityResult
)
from module.ztpl_app_module_name.ztpl_app_crud_entity.repo.ztpl_app_crud_entity_repo import (
    ZtplAppCrudEntityRepo
)
from fastapi import HTTPException


class ZtplAppCrudEntityService():
    '''
    Service to handle ztpl_app_crud_entity
    '''

    def __init__(
        self,
        mb: AppMessageBus,
        rpc: AppRPC,
        ztpl_app_crud_entity_repo: ZtplAppCrudEntityRepo
    ):
        '''
        Init ZtplAppCrudEntity service.
        '''
        self.mb = mb
        self.rpc = rpc
        self.ztpl_app_crud_entity_repo = ztpl_app_crud_entity_repo

    def find(
        self,
        keyword: str,
        limit: int,
        offset: int,
        current_user: Optional[User] = None
    ) -> ZtplAppCrudEntityResult:
        '''
        Find ztpl_app_crud_entities
        '''
        count = self.ztpl_app_crud_entity_repo.count(keyword)
        rows = [
            self._fulfill_ztpl_app_crud_entity(row)
            for row in self.ztpl_app_crud_entity_repo.find(
                keyword, limit, offset
            )
        ]
        return ZtplAppCrudEntityResult(count=count, rows=rows)

    def find_by_id(
        self,
        id: str,
        current_user: Optional[User] = None
    ) -> Optional[ZtplAppCrudEntity]:
        '''
        Find ztpl_app_crud_entity
        '''
        ztpl_app_crud_entity = self._find_ztpl_app_crud_entity_by_id_or_error(
            id, current_user
        )
        ztpl_app_crud_entity = self._fulfill_ztpl_app_crud_entity(
            ztpl_app_crud_entity
        )
        return ztpl_app_crud_entity

    def insert(
        self,
        ztpl_app_crud_entity_data: ZtplAppCrudEntityData,
        current_user: User
    ) -> Optional[ZtplAppCrudEntity]:
        '''
        Insert ztpl_app_crud_entity
        '''
        ztpl_app_crud_entity_data.created_by = current_user.id
        ztpl_app_crud_entity_data.updated_by = current_user.id
        ztpl_app_crud_entity_data = self._validate_ztpl_app_crud_entity_data(
            ztpl_app_crud_entity_data
        )
        new_ztpl_app_crud_entity = self.ztpl_app_crud_entity_repo.insert(
            ztpl_app_crud_entity_data
        )
        self.mb.publish_activity(ActivityData(
            user_id=current_user.id,
            activity='insert',
            object='ztplAppCrudEntity',
            row=new_ztpl_app_crud_entity.dict(),
            row_id=new_ztpl_app_crud_entity.id
        ))
        new_ztpl_app_crud_entity = self._fulfill_ztpl_app_crud_entity(
            new_ztpl_app_crud_entity
        )
        return new_ztpl_app_crud_entity

    def update(
        self,
        id: str,
        ztpl_app_crud_entity_data: ZtplAppCrudEntityData,
        current_user: User
    ) -> Optional[ZtplAppCrudEntity]:
        '''
        Update ztpl_app_crud_entity
        '''
        self._find_ztpl_app_crud_entity_by_id_or_error(id, current_user)
        ztpl_app_crud_entity_data.updated_by = current_user.id
        ztpl_app_crud_entity_data = self._validate_ztpl_app_crud_entity_data(
            ztpl_app_crud_entity_data, id
        )
        updated_ztpl_app_crud_entity = self.ztpl_app_crud_entity_repo.update(
            id,
            ztpl_app_crud_entity_data
        )
        self.mb.publish_activity(ActivityData(
            user_id=current_user.id,
            activity='update',
            object='ztplAppCrudEntity',
            row=updated_ztpl_app_crud_entity.dict(),
            row_id=updated_ztpl_app_crud_entity.id
        ))
        updated_ztpl_app_crud_entity = self._fulfill_ztpl_app_crud_entity(
            updated_ztpl_app_crud_entity
        )
        return updated_ztpl_app_crud_entity

    def delete(
        self,
        id: str,
        current_user: User
    ) -> Optional[ZtplAppCrudEntity]:
        '''
        Delete ztpl_app_crud_entity
        '''
        self._find_ztpl_app_crud_entity_by_id_or_error(id, current_user)
        deleted_ztpl_app_crud_entity = self.ztpl_app_crud_entity_repo.delete(
            id
        )
        self.mb.publish_activity(ActivityData(
            user_id=current_user.id,
            activity='delete',
            object='ztplAppCrudEntity',
            row=deleted_ztpl_app_crud_entity.dict(),
            row_id=deleted_ztpl_app_crud_entity.id
        ))
        deleted_ztpl_app_crud_entity = self._fulfill_ztpl_app_crud_entity(
            deleted_ztpl_app_crud_entity
        )
        return deleted_ztpl_app_crud_entity

    def _find_ztpl_app_crud_entity_by_id_or_error(
        self,
        id: Optional[str] = None,
        current_user: Optional[User] = None
    ) -> Optional[ZtplAppCrudEntity]:
        '''
        Find ztpl_app_crud_entity or throw an error if not found
        '''
        ztpl_app_crud_entity = self.ztpl_app_crud_entity_repo.find_by_id(id)
        if ztpl_app_crud_entity is None:
            raise HTTPException(
                status_code=404,
                detail='ztplAppCrudEntity id not found: {}'.format(id)
            )
        return ztpl_app_crud_entity

    def _fulfill_ztpl_app_crud_entity(
        self,
        ztpl_app_crud_entity: ZtplAppCrudEntity
    ) -> ZtplAppCrudEntity:
        '''
        Complete ztpl_app_crud_entity.
        You can use this method to add default fields etc.
        '''
        return ztpl_app_crud_entity

    def _validate_ztpl_app_crud_entity_data(
        self,
        ztpl_app_crud_entity_data: ZtplAppCrudEntityData,
        id: Optional[str] = None
    ) -> ZtplAppCrudEntityData:
        '''
        Validate ztpl_app_crud_entity_data.
        You can throw HTTPException when the data is not right
        '''
        return ztpl_app_crud_entity_data
