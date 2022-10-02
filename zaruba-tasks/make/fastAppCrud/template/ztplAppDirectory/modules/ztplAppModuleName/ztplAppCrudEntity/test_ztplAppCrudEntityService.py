from typing import Optional, List
from modules.ztplAppModuleName.ztplAppCrudEntity.ztplAppCrudEntityService import ZtplAppCrudEntityService
from modules.ztplAppModuleName.ztplAppCrudEntity.repos.ztplAppCrudEntityRepo import ZtplAppCrudEntityRepo
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from helpers.transport import LocalRPC, LocalMessageBus

################################################
# -- ⚙️ Mock data and objects
################################################

mock_ztpl_app_crud_entity_data = ZtplAppCrudEntityData(
    created_by='mock_user_id'
)

mock_ztpl_app_crud_entity = ZtplAppCrudEntity(
    id='mock_ztpl_app_crud_entity_id',
    created_by='mock_user_id',
    updated_by='mock_user_id'
)

class MockZtplAppCrudEntityRepo(ZtplAppCrudEntityRepo):

    def __init__(self):
        self.find_id: Optional[str] = None
        self.find_keyword: Optional[str] = None
        self.count_keyword: Optional[str] = None
        self.find_limit: Optional[int] = None
        self.find_offset: Optional[int] = None
        self.insert_ztpl_app_crud_entity_data: Optional[ZtplAppCrudEntityData] = None
        self.update_id: Optional[str] = None
        self.update_ztpl_app_crud_entity_data: Optional[ZtplAppCrudEntityData] = None
        self.delete_id: Optional[str] = None

    def find_by_id(self, id: str) -> Optional[ZtplAppCrudEntity]:
        self.find_id = id
        return mock_ztpl_app_crud_entity

    def find(self, keyword: str, limit: str, offset: int) -> List[ZtplAppCrudEntity]:
        self.find_keyword = keyword
        self.find_limit = limit
        self.find_offset = offset
        return [mock_ztpl_app_crud_entity]

    def count(self, keyword: str) -> int:
        self.count_keyword = keyword
        return 1

    def insert(self, ztplAppCrudEntity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        self.insert_ztpl_app_crud_entity_data = ztplAppCrudEntity_data
        return mock_ztpl_app_crud_entity

    def update(self, id: str, ztplAppCrudEntity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        self.update_id = id
        self.update_ztpl_app_crud_entity_data = ztplAppCrudEntity_data
        return mock_ztpl_app_crud_entity

    def delete(self, id: str) -> Optional[ZtplAppCrudEntity]:
        self.delete_id = id
        return mock_ztpl_app_crud_entity


################################################
# -- 🧪 Test
################################################

def test_ztpl_app_crud_entity_service_find():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_ztpl_app_crud_entity_repo = MockZtplAppCrudEntityRepo()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mock_mb, mock_rpc, mock_ztpl_app_crud_entity_repo)
    ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find('find_keyword', 73, 37)
    # make sure all parameters are passed to repo
    assert mock_ztpl_app_crud_entity_repo.find_keyword == 'find_keyword'
    assert mock_ztpl_app_crud_entity_repo.find_limit == 73
    assert mock_ztpl_app_crud_entity_repo.find_offset == 37
    assert mock_ztpl_app_crud_entity_repo.count_keyword == 'find_keyword'
    # make sure ztpl_app_crud_entity_service return the result correctly
    assert ztpl_app_crud_entity_result.count == 1
    assert len(ztpl_app_crud_entity_result.rows) == 1
    assert ztpl_app_crud_entity_result.rows[0] == mock_ztpl_app_crud_entity


def test_ztpl_app_crud_entity_service_find_by_id():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_ztpl_app_crud_entity_repo = MockZtplAppCrudEntityRepo()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mock_mb, mock_rpc, mock_ztpl_app_crud_entity_repo)
    ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id('find_id')
    # make sure all parameters are passed to repo
    assert mock_ztpl_app_crud_entity_repo.find_id == 'find_id'
    # make sure ztpl_app_crud_entity_service return the result correctly
    assert ztpl_app_crud_entity == mock_ztpl_app_crud_entity


def test_ztpl_app_crud_entity_service_insert():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_ztpl_app_crud_entity_repo = MockZtplAppCrudEntityRepo()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mock_mb, mock_rpc, mock_ztpl_app_crud_entity_repo)
    new_ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(mock_ztpl_app_crud_entity_data)
    # make sure all parameters are passed to repo
    assert mock_ztpl_app_crud_entity_repo.insert_ztpl_app_crud_entity_data == mock_ztpl_app_crud_entity_data
    # make sure ztpl_app_crud_entity_service return the result correctly
    assert new_ztpl_app_crud_entity == mock_ztpl_app_crud_entity


def test_ztpl_app_crud_entity_service_update():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_ztpl_app_crud_entity_repo = MockZtplAppCrudEntityRepo()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mock_mb, mock_rpc, mock_ztpl_app_crud_entity_repo)
    updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update('update_id', mock_ztpl_app_crud_entity_data)
    # make sure all parameters are passed to repo
    assert mock_ztpl_app_crud_entity_repo.update_id == 'update_id'
    assert mock_ztpl_app_crud_entity_repo.update_ztpl_app_crud_entity_data == mock_ztpl_app_crud_entity_data
    # make sure ztpl_app_crud_entity_service return the result correctly
    assert updated_ztpl_app_crud_entity == mock_ztpl_app_crud_entity


def test_ztpl_app_crud_entity_service_delete():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_ztpl_app_crud_entity_repo = MockZtplAppCrudEntityRepo()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mock_mb, mock_rpc, mock_ztpl_app_crud_entity_repo)
    deleted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete('delete_id')
    # make sure all parameters are passed to repo
    assert mock_ztpl_app_crud_entity_repo.delete_id == 'delete_id'
    # make sure ztpl_app_crud_entity_service return the result correctly
    assert deleted_ztpl_app_crud_entity == mock_ztpl_app_crud_entity

