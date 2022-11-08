from typing import Optional, Tuple
from modules.cms.contentType.repos.contentTypeRepo import ContentTypeRepo
from schemas.contentType import ContentType, ContentTypeData
from modules.cms.contentType.contentTypeService import ContentTypeService
from modules.cms.contentType.repos.dbContentTypeRepo import DBContentTypeRepo
from helpers.transport import LocalRPC, LocalMessageBus, MessageBus
from transport import AppMessageBus, AppRPC
from sqlalchemy import create_engine

def create_content_type_data() -> ContentTypeData:
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_content_type_data = ContentTypeData(
        name='',
        template='',
        attributes=[],
        created_by=''
    )
    return dummy_content_type_data


def insert_content_type_data(content_type_repo: ContentTypeRepo, index: Optional[int] = None) -> ContentType:
    content_type_data = create_content_type_data()
    content_type_data.name = 'contentType' if index is None else 'contentType-{index}'.format(index=index)
    content_type_data.created_by = 'original_user'
    content_type_data.updated_by = 'original_user'
    return content_type_repo.insert(content_type_data)


def create_mb() -> AppMessageBus:
    mb = AppMessageBus(LocalMessageBus())
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_content_type_service_components() -> Tuple[ContentTypeService, DBContentTypeRepo, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    content_type_repo = DBContentTypeRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    content_type_service = ContentTypeService(mb, rpc, content_type_repo)
    return content_type_service, content_type_repo, mb, rpc
