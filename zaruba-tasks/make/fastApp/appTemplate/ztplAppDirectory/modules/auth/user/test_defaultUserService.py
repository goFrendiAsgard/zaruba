from typing import Optional, List, Tuple
from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from schemas.user import User, UserData
from schemas.role import Role, RoleData
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine

################################################
# -- ⚙️ Helpers
################################################

def create_user_data():
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_user_data = UserData(
        username='',
        email='',
        phone_number='',
        permissions=[],
        role_ids=[],
        active=True,
        full_name='',
        created_by=''
    )
    return dummy_user_data


def create_role_data():
    dummy_role_data = RoleData(
        name='',
        permissions=[],
        created_by=''
    )
    return dummy_role_data


def init_test_user_service_components() -> Tuple[DefaultUserService, RoleService, DBUserRepo, DBRoleRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')
    return user_service, role_service, user_repo, role_repo, mb, rpc


def init_user_data(user_repo: DBUserRepo, index: Optional[int] = None, permissions: List[str] = [], role_ids: List[str] = [], password: str = '', active: bool = True) -> User:
    user_data = create_user_data()
    user_data.username = 'user' if index is None else 'user-{index}'.format(index=index)
    user_data.email = '{username}@innistrad.com'.format(username=user_data.username)
    user_data.password = password
    user_data.phone_number = '+628123456789' if index is None else '+6281234567890{index}'.format(index=index)
    user_data.permissions = permissions
    user_data.role_ids = role_ids
    user_data.active = active
    user_data.full_name = '{username} Nguyen'.format(username=user_data.username)
    user_data.created_by = 'original_user'
    user_data.updated_by = 'original_user'
    return user_repo.insert(user_data)


################################################
# -- 🧪 Test
################################################

def test_user_service_get_guest_user():
    user_service, _, _, _, _, _ = init_test_user_service_components()
    # test get guest
    guest_user = user_service.get_guest()
    assert guest_user.id == 'guest'
    assert guest_user.username == 'guest_username'
    assert len(guest_user.permissions) == 0
    assert len(guest_user.role_ids) == 0


def test_user_service_crud_find_by_id_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test find by id (existing)
    fetched_user = user_service.find_by_id(existing_user.id)
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_id_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find by id (non existing)
    non_existing_user = user_service.find_by_id('invalid-id')
    assert non_existing_user is None


def test_user_service_crud_find_by_username_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test find by id (existing)
    fetched_user = user_service.find_by_username(existing_user.username)
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_username_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find by id (non existing)
    non_existing_user = user_service.find_by_username('invalid-user')
    assert non_existing_user is None


def test_user_service_crud_find_by_username_and_password_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find by id (existing)
    fetched_user = user_service.find_by_identity_and_password(existing_user.username, 'password')
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_email_and_password_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find by id (existing)
    fetched_user = user_service.find_by_identity_and_password(existing_user.email, 'password')
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_identity_and_password_invalid_password():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find by id (non existing)
    non_existing_user = user_service.find_by_identity_and_password(existing_user.username, 'invalid-password')
    assert non_existing_user is None


def test_user_service_crud_find_by_identity_and_password_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    init_user_data(user_repo, password='password')
    # test find by id (non existing)
    non_existing_user = user_service.find_by_identity_and_password('invalid-user', 'invalid-password')
    assert non_existing_user is None


def test_user_service_crud_find_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find (existing)
    fetched_user_result = user_service.find(keyword='user', limit=100, offset=0)
    assert fetched_user_result.count == 1
    fetched_user = fetched_user_result.rows[0]
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find (non existing)
    non_existing_user_result = user_service.find(keyword='invalid-keyword', limit=100, offset=0)
    assert non_existing_user_result.count == 0


def test_user_service_crud_find_pagination():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    for index in range(7):
        init_user_data(user_repo, index)
    # test find (page 1)
    fetched_user_result = user_service.find(keyword='user', limit=3, offset=0)
    assert len(fetched_user_result.rows) == 3
    assert fetched_user_result.count == 7
    # test find (page 2)
    fetched_user_result = user_service.find(keyword='user', limit=3, offset=3)
    assert len(fetched_user_result.rows) == 3
    assert fetched_user_result.count == 7
    # test find (page 3)
    fetched_user_result = user_service.find(keyword='user', limit=3, offset=6)
    assert len(fetched_user_result.rows) == 1
    assert fetched_user_result.count == 7


def test_user_service_crud_insert():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare insert
    inserted_user_data = create_user_data()
    inserted_user_data.username = 'user'
    inserted_user_data.created_by = 'original_user'
    inserted_user_data.updated_by = 'original_user'
    # test insert
    inserted_user = user_service.insert(inserted_user_data)
    assert inserted_user is not None
    assert inserted_user.id != '' 
    assert inserted_user.username == 'user'
    assert inserted_user.created_by == 'original_user'
    assert inserted_user.updated_by == 'original_user'
    assert user_repo.count(keyword='') == 1


def test_user_service_crud_update_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test update (existing)
    updated_user_data = create_user_data()
    updated_user_data.username = 'updated'
    updated_user_data.updated_by = 'editor'
    updated_user = user_service.update(existing_user.id, updated_user_data)
    assert updated_user is not None
    assert updated_user.id == existing_user.id
    assert updated_user.username == 'updated'
    assert updated_user.created_by == 'original_user'
    assert updated_user.updated_by == 'editor'
    assert user_repo.count(keyword='') == 1


def test_user_service_crud_update_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test update (non existing)
    updated_user_data = create_user_data()
    updated_user_data.username = 'updated'
    updated_user_data.updated_by = 'editor'
    updated_user = user_service.update('invalid-id', updated_user_data)
    assert updated_user == None
    assert user_repo.count(keyword='') == 1


def test_user_service_crud_delete_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test find by id (existing)
    deleted_user = user_service.delete(existing_user.id)
    assert deleted_user is not None
    assert deleted_user.id == existing_user.id
    assert deleted_user.username == 'user'
    assert deleted_user.created_by == 'original_user'
    assert deleted_user.updated_by == 'original_user'
    assert user_repo.count(keyword='') == 0


def test_user_service_crud_delete_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find by id (non existing)
    deleted_user = user_service.delete('invalid-id')
    assert deleted_user is None
    assert user_repo.count(keyword='') == 1


def test_user_service_authorization():
    user_service, _, user_repo, role_repo, _, _ = init_test_user_service_components()
    
    authorized_role_data = create_role_data()
    authorized_role_data.name='authorized_role'
    authorized_role_data.permissions=['unauthorized-1', 'authorized_permission', 'unauthorized-2']
    authorized_role = role_repo.insert(authorized_role_data)

    unauthorized_role_data = create_role_data()
    unauthorized_role_data.name='unauthorized_role'
    unauthorized_role_data.permissions=['unauthorized-1', 'unauthorized-2']
    unauthorized_role = role_repo.insert(unauthorized_role_data)

    root_user_data = create_user_data()
    root_user_data.username = 'root'
    root_user_data.email = 'root@innistrad.com'
    root_user_data.phone_number = '+6213456781'
    root_user_data.permissions = ['root']
    root_user = user_repo.insert(root_user_data)

    directly_authorized_user_data = create_user_data()
    directly_authorized_user_data.username = 'directly_authorized'
    directly_authorized_user_data.email = 'directly_authorized@innistrad.com'
    directly_authorized_user_data.phone_number = '+6213456782'
    directly_authorized_user_data.permissions = ['unauthorized-1', 'authorized_permission', 'unauthorized-2']
    directly_authorized_user = user_repo.insert(directly_authorized_user_data)

    role_authorized_user_data = create_user_data()
    role_authorized_user_data.username = 'role_authorized'
    role_authorized_user_data.email = 'role_authorized@innistrad.com'
    role_authorized_user_data.phone_number = '+6213456783'
    role_authorized_user_data.role_ids = [authorized_role.id, unauthorized_role.id]
    role_authorized_user = user_repo.insert(role_authorized_user_data)

    unauthorized_user_data = create_user_data()
    unauthorized_user_data.username = 'unauthorized'
    unauthorized_user_data.email = 'unauthorized@innistrad.com'
    unauthorized_user_data.phone_number = '+6213456784'
    unauthorized_user_data.role_ids = [unauthorized_role.id, 'invalid-role-id']
    unauthorized_user = user_repo.insert(unauthorized_user_data)

    # test authorizations
    assert user_service.is_authorized(root_user, 'authorized_permission')
    assert user_service.is_authorized(directly_authorized_user, 'authorized_permission')
    assert user_service.is_authorized(role_authorized_user, 'authorized_permission')
    assert not user_service.is_authorized(unauthorized_user, 'authorized_permission')
