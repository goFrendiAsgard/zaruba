from typing import Optional, List
from modules.auth.role.roleService import RoleService
from modules.auth.role.repos.roleRepo import RoleRepo
from modules.auth.user.userService import DefaultUserService
from modules.auth.user.repos.userRepo import UserRepo
from schemas.user import User, UserData, UserWithoutPassword
from schemas.role import Role, RoleData
from helpers.transport import LocalRPC, LocalMessageBus

import datetime

################################################
# -- ⚙️ Mock data and objects
################################################

mock_role_data = RoleData(
    name='',
    permissions=['role_permission'],
    created_by='mock_user_id'
)

mock_role = Role(
    id='mock_role_id',
    name='',
    permissions=['role_permission'],
    created_by='mock_user_id',
    updated_by='mock_user_id'
)

mock_user_data = UserData(
    username='',
    email='',
    phone_number='',
    permissions=[],
    role_ids=['mock_role_id'],
    active=True,
    full_name='',
    created_by='mock_user_id'
)

mock_user = User(
    id='mock_user_id',
    username='',
    email='',
    phone_number='',
    permissions=[],
    role_ids=['mock_role_id'],
    active=True,
    full_name='',
    created_by='mock_user_id',
    updated_by='mock_user_id'
)


class MockRoleRepo(RoleRepo):

    def __init__(self):
        self.find_ids: Optional[str] = []
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
        self.find_ids.append(id)
        if id == mock_role.id:
            return mock_role
        return None

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



class MockUserRepo(UserRepo):

    def __init__(self):
        self.find_id: Optional[str] = None
        self.find_username: Optional[str] = None
        self.find_identity: Optional[str] = None
        self.find_password: Optional[str] = None
        self.find_keyword: Optional[str] = None
        self.count_keyword: Optional[str] = None
        self.find_limit: Optional[int] = None
        self.find_offset: Optional[int] = None
        self.insert_user_data: Optional[UserData] = None
        self.update_id: Optional[str] = None
        self.update_user_data: Optional[UserData] = None
        self.delete_id: Optional[str] = None

    def find_by_id(self, id: str) -> Optional[UserWithoutPassword]:
        self.find_id = id
        return mock_user

    def find_by_username(self, username: str) -> Optional[UserWithoutPassword]:
        self.find_username = username
        return mock_user

    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[UserWithoutPassword]:
        self.find_identity = identity
        self.find_password = password
        return mock_user

    def find(self, keyword: str, limit: str, offset: int) -> List[UserWithoutPassword]:
        self.find_keyword = keyword
        self.find_limit = limit
        self.find_offset = offset
        return [mock_user]

    def count(self, keyword: str) -> int:
        self.count_keyword = keyword
        return 1

    def insert(self, user_data: UserData) -> Optional[UserWithoutPassword]:
        self.insert_user_data = user_data
        return mock_user

    def update(self, id: str, user_data: UserData) -> Optional[UserWithoutPassword]:
        self.update_id = id
        self.update_user_data = user_data
        return mock_user

    def delete(self, id: str) -> Optional[UserWithoutPassword]:
        self.delete_id = id
        return mock_user


################################################
# -- 🧪 Test
################################################

def test_user_service_get_guest():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    guest_user = user_service.get_guest()
    # make sure user_service return the result correctly
    assert guest_user.id == 'guest'
    assert guest_user.active
    assert guest_user.username == 'guest_username'
    assert guest_user.created_at == datetime.datetime.min
    assert guest_user.updated_at == datetime.datetime.min
    assert len(guest_user.permissions) == 0
    assert len(guest_user.role_ids) == 0


def test_user_service_find():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user_result = user_service.find('find_keyword', 73, 37)
    # make sure all parameters are passed to repo
    assert mock_user_repo.find_keyword == 'find_keyword'
    assert mock_user_repo.find_limit == 73
    assert mock_user_repo.find_offset == 37
    assert mock_user_repo.count_keyword == 'find_keyword'
    # make sure user_service return the result correctly
    assert user_result.count == 1
    assert len(user_result.rows) == 1
    assert user_result.rows[0].id == mock_user.id


def test_user_service_find_by_id():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = user_service.find_by_id('find_id')
    # make sure all parameters are passed to repo
    assert mock_user_repo.find_id == 'find_id'
    # make sure user_service return the result correctly
    assert user == mock_user


def test_user_service_find_by_username():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = user_service.find_by_username('find_username')
    # make sure all parameters are passed to repo
    assert mock_user_repo.find_username == 'find_username'
    # make sure user_service return the result correctly
    assert user == mock_user


def test_user_service_find_by_identity_and_password():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = user_service.find_by_identity_and_password('find_identity', 'find_password')
    # make sure all parameters are passed to repo
    assert mock_user_repo.find_identity == 'find_identity'
    assert mock_user_repo.find_password == 'find_password'
    # make sure user_service return the result correctly
    assert user == mock_user


def test_user_service_insert():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    new_user = user_service.insert(mock_user_data)
    # make sure all parameters are passed to repo
    assert mock_user_repo.insert_user_data == mock_user_data
    # make sure user_service return the result correctly
    assert new_user == mock_user


def test_user_service_update():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    updated_user = user_service.update('update_id', mock_user_data)
    # make sure all parameters are passed to repo
    assert mock_user_repo.update_id == 'update_id'
    assert mock_user_repo.update_user_data == mock_user_data
    # make sure user_service return the result correctly
    assert updated_user == mock_user


def test_user_service_delete():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    deleted_user = user_service.delete('delete_id')
    # make sure all parameters are passed to repo
    assert mock_user_repo.delete_id == 'delete_id'
    # make sure user_service return the result correctly
    assert deleted_user == mock_user


def test_user_service_is_authorized_for_user_who_has_required_permission_should_be_true():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = User(
        id='mock_user_id',
        username='',
        email='',
        phone_number='',
        permissions=['random_permission', 'required_permission', 'other_random_permission'],
        role_ids=[],
        active=True,
        full_name='',
        created_by='mock_user_id',
        updated_by='mock_user_id'
    )
    is_authorized = user_service.is_authorized(user, 'required_permission')
    # make sure user_service return the result correctly
    assert is_authorized


def test_user_service_is_authorized_for_user_who_has_root_permission_should_be_true():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = User(
        id='mock_user_id',
        username='',
        email='',
        phone_number='',
        permissions=['random_permission', 'root_permission', 'other_random_permission'],
        role_ids=[],
        active=True,
        full_name='',
        created_by='mock_user_id',
        updated_by='mock_user_id'
    )
    is_authorized = user_service.is_authorized(user, 'required_permission')
    # make sure user_service return the result correctly
    assert is_authorized


def test_user_service_is_authorized_for_user_who_has_wrong_permission_should_be_false():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = User(
        id='mock_user_id',
        username='',
        email='',
        phone_number='',
        permissions=['random_permission', 'other_random_permission'],
        role_ids=[],
        active=True,
        full_name='',
        created_by='mock_user_id',
        updated_by='mock_user_id'
    )
    is_authorized = user_service.is_authorized(user, 'required_permission')
    # make sure user_service return the result correctly
    assert not is_authorized


def test_user_service_is_authorized_for_user_who_has_role_permission_should_be_true():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = User(
        id='mock_user_id',
        username='',
        email='',
        phone_number='',
        permissions=[],
        role_ids=['invalid_role_id_1', 'mock_role_id', 'invalid_role_id_2'],
        active=True,
        full_name='',
        created_by='mock_user_id',
        updated_by='mock_user_id'
    )
    is_authorized = user_service.is_authorized(user, 'role_permission')
    # make sure all parameters are passed to role service
    assert len(mock_role_repo.find_ids) == 2 # invalid_role_id_2 should not be checked
    # make sure user_service return the result correctly
    assert is_authorized


def test_user_service_is_authorized_for_user_who_has_no_role_permission_should_be_false():
    mock_mb = LocalMessageBus()
    mock_rpc = LocalRPC()
    mock_user_repo = MockUserRepo()
    mock_role_repo = MockRoleRepo()
    mock_role_service = RoleService(mock_mb, mock_rpc, mock_role_repo)
    user_service = DefaultUserService(mock_mb, mock_rpc, mock_user_repo, mock_role_service, 'guest_username', 'root_permission')
    user = User(
        id='mock_user_id',
        username='',
        email='',
        phone_number='',
        permissions=[],
        role_ids=['invalid_role_id_1', 'mock_role_id', 'invalid_role_id_2'],
        active=True,
        full_name='',
        created_by='mock_user_id',
        updated_by='mock_user_id'
    )
    is_authorized = user_service.is_authorized(user, 'inexist_permission')
    # make sure all parameters are passed to role service
    assert len(mock_role_repo.find_ids) == 3
    # make sure user_service return the result correctly
    assert not is_authorized

