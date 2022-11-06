from typing import Optional, Tuple
from modules.cms.contentAttribute.repos.contentAttributeRepo import ContentAttributeRepo
from schemas.contentAttribute import ContentAttribute, ContentAttributeData
from modules.cms.contentAttribute.contentAttributeService import ContentAttributeService
from modules.cms.contentAttribute.repos.dbContentAttributeRepo import DBContentAttributeRepo
from helpers.transport import LocalRPC, LocalMessageBus, MessageBus
from transport import AppMessageBus, AppRPC
from sqlalchemy import create_engine

def create_content_attribute_data() -> ContentAttributeData:
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_content_attribute_data = ContentAttributeData(
        content_id='',
        key='',
        value='',
        created_by=''
    )
    return dummy_content_attribute_data


def insert_content_attribute_data(content_attribute_repo: ContentAttributeRepo, index: Optional[int] = None) -> ContentAttribute:
    content_attribute_data = create_content_attribute_data()
    content_attribute_data.content_id = 'contentAttribute' if index is None else 'contentAttribute-{index}'.format(index=index)
    content_attribute_data.created_by = 'original_user'
    content_attribute_data.updated_by = 'original_user'
    return content_attribute_repo.insert(content_attribute_data)


def create_mb() -> AppMessageBus:
    mb = AppMessageBus(LocalMessageBus())
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_content_attribute_service_components() -> Tuple[ContentAttributeService, DBContentAttributeRepo, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    content_attribute_repo = DBContentAttributeRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    content_attribute_service = ContentAttributeService(mb, rpc, content_attribute_repo)
    return content_attribute_service, content_attribute_repo, mb, rpc
