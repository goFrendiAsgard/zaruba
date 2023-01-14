from typing import Optional, Tuple
from module.cms.content.repo.content_repo import ContentRepo
from schema.content import Content, ContentData, ContentType
from schema.content_type import ContentTypeData
from module.cms.content.content_service import ContentService
from module.cms.content.repo.db_content_repo import DBContentRepo
from module.cms.content_type.repo.db_content_type_repo import DBContentTypeRepo
from helper.transport import LocalRPC, LocalMessageBus, MessageBus
from transport import AppMessageBus, AppRPC
from sqlalchemy import create_engine

def create_content_data(content_type_id: str) -> ContentData:
    # Note: 🤖 Don't delete the following statement
    dummy_content_data = ContentData(
        content_type_id=content_type_id,
        title='',
        attributes={},
        description='',
        created_by=''
    )
    return dummy_content_data


def insert_content_data(content_repo: ContentRepo, content_type_id: str, index: Optional[int] = None) -> Content:
    content_data = create_content_data(content_type_id)
    content_data.title = 'content' if index is None else 'content-{index}'.format(index=index)
    content_data.created_by = 'original_user'
    content_data.updated_by = 'original_user'
    return content_repo.insert(content_data)


def create_mb() -> AppMessageBus:
    mb = AppMessageBus(LocalMessageBus())
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_content_service_components() -> Tuple[ContentService, DBContentTypeRepo, DBContentRepo, ContentType, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    content_type_repo = DBContentTypeRepo(engine=engine, create_all=True)
    content_type = content_type_repo.insert(ContentTypeData(name='article'))
    content_repo = DBContentRepo(engine=engine, create_all=True)
    DBContentTypeRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    content_service = ContentService(mb, rpc, content_repo, content_type_repo)
    return content_service, content_type_repo, content_repo, content_type, mb, rpc
