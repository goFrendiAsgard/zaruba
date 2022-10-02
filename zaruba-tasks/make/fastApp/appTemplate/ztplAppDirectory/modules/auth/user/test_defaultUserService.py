from typing import Optional, List
from modules.auth.user.userService import DefaultUserService
from modules.auth.role.roleService import RoleService
from modules.auth.user.repos.dbUserRepo import DBUserRepo
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from schemas.user import UserData
from schemas.role import RoleData
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine

################################################
# -- ⚙️ Helpers
################################################

def create_user_data():
    dummy_user_data = UserData(
        username='',
        role_ids=[],
        permissions=[],
    )
    return dummy_user_data


def create_role_data():
    dummy_role_data = RoleData(
        name='',
        permissions=[],
    )
    return dummy_role_data


################################################
# -- 🧪 Test
################################################

def test_user_service():
    engine = create_engine('sqlite://', echo=True)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')

    # get guest
    guest_user = user_service.get_guest()
    assert guest_user.id == 'guest'
    assert guest_user.username == 'guest_username'
    assert len(guest_user.permissions) == 0
    assert len(guest_user.role_ids) == 0

    # prepare insert
    inserted_user_data = create_user_data()
    inserted_user_data.username = 'original'
    inserted_user_data.password = 'original_password'
    inserted_user_data.created_by = 'original_user'
    inserted_user_data.updated_by = 'original_user'
    # test insert
    inserted_user = user_service.insert(inserted_user_data)
    assert inserted_user is not None
    assert inserted_user.id != '' 
    assert inserted_user.username == 'original'
    assert inserted_user.created_by == 'original_user'
    assert inserted_user.updated_by == 'original_user'

    # test find by id (existing, after insert)
    existing_user = user_service.find_by_id(inserted_user.id)
    assert existing_user is not None
    assert existing_user.id == inserted_user.id
    assert existing_user.username == inserted_user.username
    assert existing_user.created_by == inserted_user.created_by
    assert existing_user.updated_by == inserted_user.updated_by

    # test find by username (existing, after insert)
    existing_user = user_service.find_by_username('original')
    assert existing_user is not None
    assert existing_user.id == inserted_user.id
    assert existing_user.username == inserted_user.username
    assert existing_user.created_by == inserted_user.created_by
    assert existing_user.updated_by == inserted_user.updated_by

    # test find by identity and password (existing, after insert)
    existing_user = user_service.find_by_identity_and_password('original', 'original_password')
    assert existing_user is not None
    assert existing_user.id == inserted_user.id
    assert existing_user.username == inserted_user.username
    assert existing_user.created_by == inserted_user.created_by
    assert existing_user.updated_by == inserted_user.updated_by

    # test find by id (non existing)
    non_existing_user = user_service.find_by_id('invalid_id')
    assert non_existing_user is None

    # test find by username (non existing)
    non_existing_user = user_service.find_by_username('invalid_username')
    assert non_existing_user is None

    # test find by identity and password (existing, invalid password)
    non_existing_user = user_service.find_by_identity_and_password('original', 'invalid_password')
    assert non_existing_user is None

    # test find by identity and password (non existing)
    non_existing_user = user_service.find_by_identity_and_password('invalid_username', 'invalid_password')
    assert non_existing_user is None

    # prepare update (existing)
    updated_user_data = create_user_data()
    updated_user_data.username = 'updated'
    updated_user_data.updated_by = 'editor'
    # test update (existing)
    updated_user = user_service.update(inserted_user.id, updated_user_data)
    assert updated_user is not None
    assert updated_user.id == inserted_user.id
    assert updated_user.username == 'updated'
    assert updated_user.created_by == 'original_user'
    assert updated_user.updated_by == 'editor'

    # test update (non existing)
    non_existing_user = user_service.update('invalid_id', updated_user_data)
    assert non_existing_user is None

    # test find by id (existing, after insert)
    existing_user = user_service.find_by_id(updated_user.id)
    assert existing_user is not None
    assert existing_user.id == inserted_user.id
    assert existing_user.username == 'updated'
    assert existing_user.created_by == 'original_user'
    assert existing_user.updated_by == 'editor'

    # test find (before delete, correct keyword)
    existing_result = user_service.find(keyword='updated', limit=10, offset=0)
    assert existing_result.count == 1
    assert len(existing_result.rows) == 1
    assert existing_result.rows[0].id == inserted_user.id

    # test find (before delete, incorrect keyword)
    non_existing_result = user_service.find(keyword='incorrect', limit=10, offset=0)
    assert non_existing_result.count == 0
    assert len(non_existing_result.rows) == 0

    # test delete existing
    deleted_user = user_service.delete(inserted_user.id)
    assert deleted_user is not None
    assert deleted_user.id == inserted_user.id
    assert deleted_user.username == 'updated'
    assert deleted_user.created_by == 'original_user'
    assert deleted_user.updated_by == 'editor'

    # test delete (non existing)
    non_existing_user = user_service.delete('invalid_id')
    assert non_existing_user is None

    # test find (after delete, correct keyword)
    non_existing_result = user_service.find(keyword='updated', limit=10, offset=0)
    assert non_existing_result.count == 0
    assert len(non_existing_result.rows) == 0

   
def test_user_service_authorization():
    engine = create_engine('sqlite://', echo=True)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    user_repo = DBUserRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, 'guest_username', 'root')

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
