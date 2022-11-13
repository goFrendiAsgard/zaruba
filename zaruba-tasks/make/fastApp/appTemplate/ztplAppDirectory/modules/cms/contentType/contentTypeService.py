from typing import Optional
from transport import AppMessageBus, AppRPC
from schemas.user import User
from schemas.activity import ActivityData
from schemas.contentType import ContentType, ContentTypeData, ContentTypeResult
from modules.cms.contentType.repos.contentTypeRepo import ContentTypeRepo
from fastapi import HTTPException

class ContentTypeService():

    def __init__(self, mb: AppMessageBus, rpc: AppRPC, content_type_repo: ContentTypeRepo):
        self.mb = mb
        self.rpc = rpc
        self.content_type_repo = content_type_repo


    def find(self, keyword: str, limit: int, offset: int, current_user: Optional[User] = None) -> ContentTypeResult:
        count = self.content_type_repo.count(keyword)
        rows = [self._fulfill(row) for row in self.content_type_repo.find(keyword, limit, offset)]
        return ContentTypeResult(count=count, rows=rows)


    def find_by_id(self, id: str, current_user: Optional[User] = None) -> Optional[ContentType]:
        content_type = self._find_by_id_or_error(id, current_user)
        content_type = self._fulfill(content_type)
        return content_type


    def find_by_name(self, name: str, current_user: Optional[User] = None) -> Optional[ContentType]:
        content_type = self.content_type_repo.find_by_name(name)
        if content_type is None:
            raise HTTPException(
                status_code=404, 
                detail='content type not found: {}'.format(name)
            )
        content_type = self._fulfill(content_type)
        return content_type


    def insert(self, content_type_data: ContentTypeData, current_user: User) -> Optional[ContentType]:
        content_type_data.created_by = current_user.id
        content_type_data.updated_by = current_user.id
        content_type_data = self._validate_data(content_type_data)
        new_content_type = self.content_type_repo.insert(content_type_data)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'insert',
            object = 'contentType',
            row = new_content_type.dict(),
            row_id = new_content_type.id
        ))
        new_content_type = self._fulfill(new_content_type)
        return new_content_type


    def update(self, id: str, content_type_data: ContentTypeData, current_user: User) -> Optional[ContentType]:
        self._find_by_id_or_error(id, current_user)
        content_type_data.updated_by = current_user.id
        content_type_data = self._validate_data(content_type_data, id)
        updated_content_type = self.content_type_repo.update(id, content_type_data)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'update',
            object = 'contentType',
            row = updated_content_type.dict(),
            row_id = updated_content_type.id
        ))
        updated_content_type = self._fulfill(updated_content_type)
        return updated_content_type


    def delete(self, id: str, current_user: User) -> Optional[ContentType]:
        self._find_by_id_or_error(id, current_user)
        deleted_content_type = self.content_type_repo.delete(id)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'delete',
            object = 'contentType',
            row = deleted_content_type.dict(),
            row_id = deleted_content_type.id
        ))
        deleted_content_type = self._fulfill(deleted_content_type)
        return deleted_content_type


    def _find_by_id_or_error(self, id: Optional[str] = None, current_user: Optional[User] = None) -> Optional[ContentType]:
        content_type = self.content_type_repo.find_by_id(id)
        if content_type is None:
            raise HTTPException(
                status_code=404, 
                detail='content type id not found: {}'.format(id)
            )
        return content_type


    def _fulfill(self, content_type: ContentType) -> ContentType:
        return content_type


    def _validate_data(self, content_type_data: ContentTypeData, id: Optional[str] = None) -> ContentTypeData:
        if content_type_data.name is not None:
            content_type = self.content_type_repo.find_by_name(content_type_data.name)
            if content_type is not None and (id is None or content_type.id != id):
                raise HTTPException(
                    status_code=422, 
                    detail='content type with the same name already exist: {}'.format(content_type_data.name)
                )
        return content_type_data
