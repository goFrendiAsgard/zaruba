from typing import Any, Optional, Mapping
from core import AuthService
from transport import AppMessageBus, AppRPC
from schemas.contentAttribute import ContentAttribute, ContentAttributeData
from schemas.user import User
from modules.cms.contentAttribute.contentAttributeService import ContentAttributeService

import sys

def register_content_attribute_rpc(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, content_attribute_service: ContentAttributeService):

    @rpc.handle('find_content_attribute')
    def find_content_attributes(keyword: str, limit: int, offset: int, current_user_data: Optional[Mapping[str, Any]]) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        content_attribute_result = content_attribute_service.find(keyword, limit, offset, current_user)
        return content_attribute_result.dict()


    @rpc.handle('find_content_attribute_by_id')
    def find_content_attribute_by_id(id: str, current_user_data: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        content_attribute = content_attribute_service.find_by_id(id, current_user)
        return None if content_attribute is None else content_attribute.dict()


    @rpc.handle('insert_content_attribute')
    def insert_content_attribute(content_attribute_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content_attribute = ContentAttributeData.parse_obj(content_attribute_data) 
        new_content_attribute = content_attribute_service.insert(content_attribute, current_user)
        return None if new_content_attribute is None else new_content_attribute.dict()


    @rpc.handle('update_content_attribute')
    def update_content_attribute(id: str, content_attribute_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content_attribute = ContentAttributeData.parse_obj(content_attribute_data) 
        content_attribute.updated_by = current_user.id
        updated_content_attribute = content_attribute_service.update(id, content_attribute, current_user)
        return None if updated_content_attribute is None else updated_content_attribute.dict()


    @rpc.handle('delete_content_attribute')
    def delete_content_attribute(id: str, current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content_attribute = content_attribute_service.delete(id, current_user)
        return None if content_attribute is None else content_attribute.dict()


    print('Handle RPC for cms.ContentAttribute', file=sys.stderr)