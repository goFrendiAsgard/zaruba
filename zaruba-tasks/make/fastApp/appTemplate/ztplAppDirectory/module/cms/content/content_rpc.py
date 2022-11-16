from typing import Any, Optional, Mapping
from core import AuthService
from transport import AppMessageBus, AppRPC
from schema.content import Content, ContentData
from schema.user import User
from module.cms.content.content_service import ContentService

import sys

def register_content_rpc(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, content_service: ContentService):

    @rpc.handle('find_content')
    def find_contents(keyword: str, limit: int, offset: int, current_user_data: Optional[Mapping[str, Any]]) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        content_result = content_service.find(keyword, limit, offset, current_user)
        return content_result.dict()


    @rpc.handle('find_content_by_id')
    def find_content_by_id(id: str, current_user_data: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        content = content_service.find_by_id(id, current_user)
        return None if content is None else content.dict()


    @rpc.handle('insert_content')
    def insert_content(content_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content = ContentData.parse_obj(content_data) 
        new_content = content_service.insert(content, current_user)
        return None if new_content is None else new_content.dict()


    @rpc.handle('update_content')
    def update_content(id: str, content_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content = ContentData.parse_obj(content_data) 
        content.updated_by = current_user.id
        updated_content = content_service.update(id, content, current_user)
        return None if updated_content is None else updated_content.dict()


    @rpc.handle('delete_content')
    def delete_content(id: str, current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        content = content_service.delete(id, current_user)
        return None if content is None else content.dict()


    print('Handle RPC for cms.Content', file=sys.stderr)