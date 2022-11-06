from typing import Optional
from transport import AppMessageBus, AppRPC
from schemas.user import User
from schemas.activity import ActivityData
from schemas.content import Content, ContentData, ContentResult
from modules.cms.content.repos.contentRepo import ContentRepo
from fastapi import HTTPException

class ContentService():

    def __init__(self, mb: AppMessageBus, rpc: AppRPC, content_repo: ContentRepo):
        self.mb = mb
        self.rpc = rpc
        self.content_repo = content_repo


    def find(self, keyword: str, limit: int, offset: int, current_user: Optional[User] = None) -> ContentResult:
        count = self.content_repo.count(keyword)
        rows = self.content_repo.find(keyword, limit, offset)
        return ContentResult(count=count, rows=rows)


    def find_by_id(self, id: str, current_user: Optional[User] = None) -> Optional[Content]:
        content = self._find_by_id_or_error(id, current_user)
        return content


    def insert(self, content_data: ContentData, current_user: User) -> Optional[Content]:
        content_data.created_by = current_user.id
        content_data.updated_by = current_user.id
        content_data = self._validate_data(content_data)
        new_content = self.content_repo.insert(content_data)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'insert',
            object = 'content',
            row = new_content.dict(),
            row_id = new_content.id
        ))
        return new_content


    def update(self, id: str, content_data: ContentData, current_user: User) -> Optional[Content]:
        self._find_by_id_or_error(id, current_user)
        content_data.updated_by = current_user.id
        content_data = self._validate_data(content_data, id)
        updated_content = self.content_repo.update(id, content_data)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'update',
            object = 'content',
            row = updated_content.dict(),
            row_id = updated_content.id
        ))
        return updated_content


    def delete(self, id: str, current_user: User) -> Optional[Content]:
        self._find_by_id_or_error(id, current_user)
        deleted_content = self.content_repo.delete(id)
        self.mb.publish_activity(ActivityData(
            user_id = current_user.id,
            activity = 'delete',
            object = 'content',
            row = deleted_content.dict(),
            row_id = deleted_content.id
        ))
        return deleted_content


    def _find_by_id_or_error(self, id: Optional[str] = None, current_user: Optional[User] = None) -> Optional[Content]:
        content = self.content_repo.find_by_id(id)
        if content is None:
            raise HTTPException(
                status_code=404, 
                detail='Content id not found: {}'.format(id)
            )
        return content


    def _validate_data(self, content_data: ContentData, id: Optional[str] = None) -> ContentData:
        # TODO: add your custom logic
        # Example: checking duplication
        # if content_data.some_field is not None:
        #     user = self.user_repo.find_by_some_field(content_data.some_field)
        #     if user is not None and (id is None or user.id != id):
        #         raise HTTPException(
        #             status_code=422, 
        #             detail='some_field already exist: {}'.format(content_data.some_field)
        #         )
        return content_data
