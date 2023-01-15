from typing import Any, Optional, Mapping
from core import AuthService
from transport import AppMessageBus, AppRPC
from schema.content_type import ContentType, ContentTypeData
from schema.user import User
from module.cms.content_type.content_type_service import ContentTypeService

import logging


def register_content_type_rpc(
    mb: AppMessageBus, 
    rpc: AppRPC, 
    auth_service: AuthService, 
    content_type_service: ContentTypeService
):

    @rpc.handle('find_content_type')
    def find_content_types(
        keyword: str,
        limit: int,
        offset: int,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Mapping[str, Any]:
        current_user = _get_user_from_dict(current_user_data)
        content_type_result = content_type_service.find(
            keyword, limit, offset, current_user
        )
        return content_type_result.dict()

    @rpc.handle('find_content_type_by_id')
    def find_content_type_by_id(
        id: str,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user_from_dict(current_user_data)
        content_type = content_type_service.find_by_id(
            id, current_user
        )
        return _content_type_as_dict(content_type)

    @rpc.handle('insert_content_type')
    def insert_content_type(
        content_type_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content_type = ContentTypeData.parse_obj(content_type_data)
        new_content_type = content_type_service.insert(
            content_type, current_user
        )
        return _content_type_as_dict(new_content_type)

    @rpc.handle('update_content_type')
    def update_content_type(
        id: str,
        content_type_data: Mapping[str, Any],
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content_type = ContentTypeData.parse_obj(content_type_data)
        content_type.updated_by = current_user.id
        updated_content_type = content_type_service.update(
            id, content_type, current_user
        )
        return _content_type_as_dict(updated_content_type)

    @rpc.handle('delete_content_type')
    def delete_content_type(
        id: str,
        current_user_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        deleted_content_type = content_type_service.delete(id, current_user)
        return _content_type_as_dict(deleted_content_type)

    def _get_user_from_dict(
        user_data: Optional[Mapping[str, Any]]
    ) -> Optional[User]:
        if user_data is None:
            return None
        return User.parse_obj(user_data)

    def _content_type_as_dict(
        content_type: Optional[ContentType]
    ) -> Optional[Mapping[str, Any]]:
        if content_type is None:
            return None
        return content_type.dict()

    logging.info('Register cms.content_type RPC handler')
