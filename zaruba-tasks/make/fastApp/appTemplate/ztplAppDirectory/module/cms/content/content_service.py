from typing import Optional
from transport import AppMessageBus, AppRPC
from schema.user import User
from schema.activity import ActivityData
from schema.content import Content, ContentData, ContentResult
from schema.content_type import ContentType
from module.cms.content.repo.content_repo import ContentRepo
from module.cms.content_type.repo.content_type_repo import ContentTypeRepo
from fastapi import HTTPException
from markdown import markdown
from jinja2 import Template


class ContentService():

    def __init__(
        self,
        mb: AppMessageBus,
        rpc: AppRPC,
        content_repo: ContentRepo,
        content_type_repo: ContentTypeRepo
    ):
        self.mb = mb
        self.rpc = rpc
        self.content_repo = content_repo
        self.content_type_repo = content_type_repo

    def find(
        self,
        keyword: str,
        limit: int,
        offset: int,
        current_user: Optional[User] = None
    ) -> ContentResult:
        count = self.content_repo.count(keyword)
        rows = [
            self._fulfill(row)
            for row in self.content_repo.find(keyword, limit, offset)
        ]
        return ContentResult(count=count, rows=rows)

    def find_by_id(
        self, id: str, current_user: Optional[User] = None
    ) -> Optional[Content]:
        content = self._find_by_id_or_error(id, current_user)
        content = self._fulfill(content)
        return content

    def insert(
        self, content_data: ContentData, current_user: User
    ) -> Optional[Content]:
        content_data.created_by = current_user.id
        content_data.updated_by = current_user.id
        content_data = self._validate_data(content_data)
        new_content = self.content_repo.insert(content_data)
        self.mb.publish_activity(ActivityData(
            user_id=current_user.id,
            activity='insert',
            object='content',
            row=new_content.dict(),
            row_id=new_content.id
        ))
        new_content = self._fulfill(new_content)
        return new_content

    def update(
        self, id: str, content_data: ContentData, current_user: User
    ) -> Optional[Content]:
        self._find_by_id_or_error(id, current_user)
        content_data.updated_by = current_user.id
        content_data = self._validate_data(content_data, id)
        updated_content = self.content_repo.update(id, content_data)
        self.mb.publish_activity(ActivityData(
            user_id=current_user.id,
            activity='update',
            object='content',
            row=updated_content.dict(),
            row_id=updated_content.id
        ))
        updated_content = self._fulfill(updated_content)
        return updated_content

    def delete(self, id: str, current_user: User) -> Optional[Content]:
        self._find_by_id_or_error(id, current_user)
        deleted_content = self.content_repo.delete(id)
        self.mb.publish_activity(ActivityData(
            user_id=current_user.id,
            activity='delete',
            object='content',
            row=deleted_content.dict(),
            row_id=deleted_content.id
        ))
        deleted_content = self._fulfill(deleted_content)
        return deleted_content

    def _find_by_id_or_error(
        self,
        id: Optional[str] = None,
        current_user: Optional[User] = None
    ) -> Optional[Content]:
        content = self.content_repo.find_by_id(id)
        if content is None:
            raise HTTPException(
                status_code=404,
                detail='content id not found: {}'.format(id)
            )
        return content

    def _fulfill(self, content: Content) -> Content:
        content.content_type = self.content_type_repo.find_by_id(
            content.content_type_id
        )
        if content.content_type is None:
            content.content_type = ContentType(id='default', name='default')
        jinja_template = Template(content.content_type.template)
        try:
            content.html_content = markdown(
                jinja_template.render(**content.dict())
            )
        except Exception:
            content.html_content = 'Cannot render content'
        return content

    def _validate_data(
        self, content_data: ContentData, id: Optional[str] = None
    ) -> ContentData:
        content_type = self.content_type_repo.find_by_id(
            content_data.content_type_id
        )
        if content_type is None:
            raise HTTPException(
                status_code=422,
                detail='content type is not exist: {}'.format(
                    content_data.content_type_id
                )
            )
        for cta in content_type.attributes:
            if cta.name not in content_data.attributes:
                content_data.attributes[cta.name] = cta.default_value
        return content_data
