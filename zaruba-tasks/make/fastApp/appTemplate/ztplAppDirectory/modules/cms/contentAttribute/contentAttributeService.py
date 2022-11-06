from typing import Optional
from transport import AppMessageBus, AppRPC
from schemas.user import User
from schemas.activity import ActivityData
from schemas.contentAttribute import ContentAttribute, ContentAttributeData, ContentAttributeResult
from modules.cms.contentAttribute.repos.contentAttributeRepo import ContentAttributeRepo
from fastapi import HTTPException

class ContentAttributeService():

    def __init__(self, mb: AppMessageBus, rpc: AppRPC, content_attribute_repo: ContentAttributeRepo):
        self.mb = mb
        self.rpc = rpc
        self.content_attribute_repo = content_attribute_repo


    def find(self, keyword: str, limit: int, offset: int, current_user: Optional[User] = None) -> ContentAttributeResult:
        count = self.content_attribute_repo.count(keyword)
        rows = self.content_attribute_repo.find(keyword, limit, offset)
        return ContentAttributeResult(count=count, rows=rows)


    def find_by_id(self, id: str, current_user: Optional[User] = None) -> Optional[ContentAttribute]:
        content_attribute = self._find_by_id_or_error(id, current_user)
        return content_attribute


    def insert(self, content_attribute_data: ContentAttributeData, current_user: User) -> Optional[ContentAttribute]:
        content_attribute_data.created_by = current_user.id
        content_attribute_data.updated_by = current_user.id
        content_attribute_data = self._validate_data(content_attribute_data)
        new_content_attribute = self.content_attribute_repo.insert(content_attribute_data)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'insert',
            object = 'contentAttribute',
            row = new_content_attribute.dict(),
            row_id = new_content_attribute.id
        ))
        return new_content_attribute


    def update(self, id: str, content_attribute_data: ContentAttributeData, current_user: User) -> Optional[ContentAttribute]:
        self._find_by_id_or_error(id, current_user)
        content_attribute_data.updated_by = current_user.id
        content_attribute_data = self._validate_data(content_attribute_data, id)
        updated_content_attribute = self.content_attribute_repo.update(id, content_attribute_data)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'update',
            object = 'contentAttribute',
            row = updated_content_attribute.dict(),
            row_id = updated_content_attribute.id
        ))
        return updated_content_attribute


    def delete(self, id: str, current_user: User) -> Optional[ContentAttribute]:
        self._find_by_id_or_error(id, current_user)
        deleted_content_attribute = self.content_attribute_repo.delete(id)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'delete',
            object = 'contentAttribute',
            row = deleted_content_attribute.dict(),
            row_id = deleted_content_attribute.id
        ))
        return deleted_content_attribute


    def _find_by_id_or_error(self, id: Optional[str] = None, current_user: Optional[User] = None) -> Optional[ContentAttribute]:
        content_attribute = self.content_attribute_repo.find_by_id(id)
        if content_attribute is None:
            raise HTTPException(
                status_code=404, 
                detail='ContentAttribute id not found: {}'.format(id)
            )
        return content_attribute


    def _validate_data(self, content_attribute_data: ContentAttributeData, id: Optional[str] = None) -> ContentAttributeData:
        # TODO: add your custom logic
        # Example: checking duplication
        # if content_attribute_data.some_field is not None:
        #     user = self.user_repo.find_by_some_field(content_attribute_data.some_field)
        #     if user is not None and (id is None or user.id != id):
        #         raise HTTPException(
        #             status_code=422, 
        #             detail='some_field already exist: {}'.format(content_attribute_data.some_field)
        #         )
        return content_attribute_data
