from typing import Optional, Tuple
from modules.cms.content.repos.contentRepo import ContentRepo
from schemas.content import Content, ContentData
from modules.cms.content.contentService import ContentService
from modules.cms.content.repos.dbContentRepo import DBContentRepo
from helpers.transport import LocalRPC, LocalMessageBus, MessageBus
from transport import AppMessageBus, AppRPC
from sqlalchemy import create_engine

def create_content_data() -> ContentData:
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_content_data = ContentData(
        type_id='',
        title='',
        attributes={},
        description='',
        created_by=''
    )
    return dummy_content_data


def insert_content_data(content_repo: ContentRepo, index: Optional[int] = None) -> Content:
    content_data = create_content_data()
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


def init_test_content_service_components() -> Tuple[ContentService, DBContentRepo, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    content_repo = DBContentRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    content_service = ContentService(mb, rpc, content_repo)
    return content_service, content_repo, mb, rpc
