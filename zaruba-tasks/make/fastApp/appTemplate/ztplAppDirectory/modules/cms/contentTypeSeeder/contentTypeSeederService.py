from typing import Optional
from starlette.exceptions import HTTPException
from core import AuthService
from modules.cms.contentType import ContentTypeService
from schemas.contentType import ContentType

class ContentTypeSeederService():

    def __init__(self, auth_service: AuthService, content_type_service: ContentTypeService):
        self.auth_service = auth_service
        self.content_type_service = content_type_service


    def seed(self, content_type: ContentType):
        system_user = self.auth_service.get_system_user()
        try:
            self.content_type_service.find_by_name(content_type.name, system_user)
        except HTTPException as error:
            if error.status_code == 404:
                self.content_type_service.insert(content_type, system_user)