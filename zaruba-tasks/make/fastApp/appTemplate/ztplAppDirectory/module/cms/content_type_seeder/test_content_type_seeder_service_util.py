from typing import Tuple
from fastapi.security import OAuth2PasswordBearer
from core.security.service.auth_service import AuthService
from core.security.rule.default_auth_rule import DefaultAuthRule
from core.security.middleware.default_user_fetcher import DefaultUserFetcher
from module.cms.content_type import ContentTypeService, DBContentTypeRepo
from module.cms.content_type.test_content_type_service_util import create_content_type_data
from module.cms.content_type_seeder.content_type_seeder_service import ContentTypeSeederService
from helper.transport import LocalRPC, LocalMessageBus, MessageBus
from transport import AppMessageBus, AppRPC

from sqlalchemy import create_engine


ARTICLE_CONTENT_TYPE_DATA = create_content_type_data()
ARTICLE_CONTENT_TYPE_DATA.name = 'article'


def create_mb() -> AppMessageBus:
    mb = AppMessageBus(LocalMessageBus())
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_content_type_seeder_service_components() -> Tuple[ContentTypeSeederService, ContentTypeService, DBContentTypeRepo, MessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    content_type_repo = DBContentTypeRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    content_type_service = ContentTypeService(mb, rpc, content_type_repo)
    auth_rule = DefaultAuthRule(rpc)
    oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/', auto_error = False)
    user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
    auth_service = AuthService(auth_rule, user_fetcher, 'root')
    content_type_seeder_service = ContentTypeSeederService(auth_service, content_type_service)
    return content_type_seeder_service, content_type_service, content_type_repo, mb, rpc

