from typing import Optional, List
from modules.auth.role.roleService import RoleService
from modules.auth.role.repos.roleRepo import RoleRepo
from schemas.role import Role, RoleData
from helpers.transport import LocalRPC, LocalMessageBus

################################################
# -- ⚙️ Mock data and objects
################################################

mock_role_data = RoleData(
    name='',
    permissions=[],
    created_by='mock_user_id'
)

mock_role = Role(
    id='mock_user_id',
    name='',
    permissions=[],
    created_by='mock_user_id',
    updated_by='mock_user_id'
)

class MockRoleRepo(RoleRepo):

    def __init__(self):
        self.find_id: Optional[str] = None
        self.find_keyword: Optional[str] = None
        self.count_keyword: Optional[str] = None
        self.find_name: Optional[str] = None
        self.find_limit: Optional[int] = None
        self.find_offset: Optional[int] = None
        self.insert_role_data: Optional[RoleData] = None
        self.update_id: Optional[str] = None
        self.update_role_data: Optional[RoleData] = None
        self.delete_id: Optional[str] = None

    def find_by_id(self, id: str) -> Optional[Role]:
        self.find_id = id
        return mock_role

    def find_by_name(self, name: str) -> Optional[Role]:
        self.find_name = name
        return mock_role

    def find(self, keyword: str, limit: str, offset: int) -> List[Role]:
        self.find_keyword = keyword
        self.find_limit = limit
        self.find_offset = offset
        return [mock_role]

    def count(self, keyword: str) -> int:
        self.count_keyword = keyword
        return 1

    def insert(self, role_data: RoleData) -> Optional[Role]:
        self.insert_role_data = role_data
        return mock_role

    def update(self, id: str, role_data: RoleData) -> Optional[Role]:
        self.update_id = id
        self.update_role_data = role_data
        return mock_role

    def delete(self, id: str) -> Optional[Role]:
        self.delete_id = id
        return mock_role


################################################
# -- 🧪 Test
################################################

def test_role_service_find():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_role_repo = MockRoleRepo()
    role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    role_result = role_service.find('find_keyword', 73, 37)
    # make sure all parameters are passed to repo
    assert mock_role_repo.find_keyword == 'find_keyword'
    assert mock_role_repo.find_limit == 73
    assert mock_role_repo.find_offset == 37
    assert mock_role_repo.count_keyword == 'find_keyword'
    # make sure role_service return the result correctly
    assert role_result.count == 1
    assert len(role_result.rows) == 1
    assert role_result.rows[0] == mock_role


def test_role_service_find_by_id():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_role_repo = MockRoleRepo()
    role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    role = role_service.find_by_id('find_id')
    # make sure all parameters are passed to repo
    assert mock_role_repo.find_id == 'find_id'
    # make sure role_service return the result correctly
    assert role == mock_role


def test_role_service_find_by_name():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_role_repo = MockRoleRepo()
    role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    role = role_service.find_by_name('find_name')
    # make sure all parameters are passed to repo
    assert mock_role_repo.find_name == 'find_name'
    # make sure role_service return the result correctly
    assert role == mock_role


def test_role_service_insert():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_role_repo = MockRoleRepo()
    role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    new_role = role_service.insert(mock_role_data)
    # make sure all parameters are passed to repo
    assert mock_role_repo.insert_role_data == mock_role_data
    # make sure role_service return the result correctly
    assert new_role == mock_role


def test_role_service_update():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_role_repo = MockRoleRepo()
    role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    updated_role = role_service.update('update_id', mock_role_data)
    # make sure all parameters are passed to repo
    assert mock_role_repo.update_id == 'update_id'
    assert mock_role_repo.update_role_data == mock_role_data
    # make sure role_service return the result correctly
    assert updated_role == mock_role


def test_role_service_delete():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_role_repo = MockRoleRepo()
    role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    deleted_role = role_service.delete('delete_id')
    # make sure all parameters are passed to repo
    assert mock_role_repo.delete_id == 'delete_id'
    # make sure role_service return the result correctly
    assert deleted_role == mock_role

