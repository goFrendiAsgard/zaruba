from typing import Any, Optional, Mapping
from core import AuthService
from transport import AppMessageBus, AppRPC
from schema.ztpl_app_crud_entity import (
    ZtplAppCrudEntity, ZtplAppCrudEntityData
)
from schema.user import User
from module.ztpl_app_module_name.ztpl_app_crud_entity.ztpl_app_crud_entity_service import (
    ZtplAppCrudEntityService
)
import logging


def register_ztpl_app_crud_entity_rpc(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, ztpl_app_crud_entity_service: ZtplAppCrudEntityService):

    @rpc.handle('find_ztpl_app_crud_entity')
    def find_ztpl_app_crud_entities(
        keyword: str,
        limit: int,
        offset: int,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Mapping[str, Any]:
        current_user = get_user_by_user_data(current_user_data)
        ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(
            keyword, limit, offset, current_user
        )
        return ztpl_app_crud_entity_result.dict()

    @rpc.handle('find_ztpl_app_crud_entity_by_id')
    def find_ztpl_app_crud_entity_by_id(
        id: str,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Optional[Mapping[str, Any]]:
        current_user = get_user_by_user_data(current_user_data)
        ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(
            id, current_user
        )
        return ztpl_app_crud_entity_as_dict(ztpl_app_crud_entity)

    @rpc.handle('insert_ztpl_app_crud_entity')
    def insert_ztpl_app_crud_entity(
        ztpl_app_crud_entity_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ZtplAppCrudEntityData.parse_obj(
            ztpl_app_crud_entity_data)
        new_ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(
            ztpl_app_crud_entity, current_user
        )
        return ztpl_app_crud_entity_as_dict(new_ztpl_app_crud_entity)

    @rpc.handle('update_ztpl_app_crud_entity')
    def update_ztpl_app_crud_entity(
        id: str,
        ztpl_app_crud_entity_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        ztpl_app_crud_entity = ZtplAppCrudEntityData.parse_obj(
            ztpl_app_crud_entity_data)
        ztpl_app_crud_entity.updated_by = current_user.id
        updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update(
            id, ztpl_app_crud_entity, current_user
        )
        return ztpl_app_crud_entity_as_dict(updated_ztpl_app_crud_entity)

    @rpc.handle('delete_ztpl_app_crud_entity')
    def delete_ztpl_app_crud_entity(
        id: str,
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        deleted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete(
            id, current_user
        )
        return ztpl_app_crud_entity_as_dict(deleted_ztpl_app_crud_entity)

    def get_user_by_user_data(
        user_data: Optional[Mapping[str, Any]]
    ) -> Optional[User]:
        if user_data is None:
            return None
        return User.parse_obj(user_data)

    def ztpl_app_crud_entity_as_dict(
        ztpl_app_crud_entity: Optional[ZtplAppCrudEntity]
    ) -> Optional[Mapping[str, Any]]:
        if ztpl_app_crud_entity is None:
            return None
        return ztpl_app_crud_entity.dict()

    logging.info('Register ztplAppModuleName.ztpl_app_crud_entity RPC handler')
