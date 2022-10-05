from typing import Optional, Tuple
from modules.auth.role.roleService import RoleService
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from schemas.role import Role, RoleData
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine

################################################
# -- ⚙️ Helpers
################################################

def create_role_data():
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_role_data = RoleData(
        name='',
        permissions=[],
        created_by=''
    )
    return dummy_role_data


def init_test_role_components() -> Tuple[RoleService, DBRoleRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)
    return role_service, role_repo, mb, rpc


def insert_role_data(role_repo: DBRoleRepo, index: Optional[int] = None) -> Role:
    role_data = create_role_data()
    role_data.name = 'original' if index is None else 'original-{index}'.format(index=index)
    role_data.created_by = 'original_user'
    role_data.updated_by = 'original_user'
    return role_repo.insert(role_data)


################################################
# -- 🧪 Test
################################################

def test_role_service_crud_find_by_id_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find by id (existing)
    fetched_role = role_service.find_by_id(existing_role.id)
    assert fetched_role is not None
    assert fetched_role.id == existing_role.id
    assert fetched_role.name == 'original'
    assert fetched_role.created_by == 'original_user'
    assert fetched_role.updated_by == 'original_user'


def test_role_service_crud_find_by_id_non_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find by id (non existing)
    non_existing_role = role_service.find_by_id('invalid-id')
    assert non_existing_role is None


def test_role_service_crud_find_by_name_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find by id (existing)
    fetched_role = role_service.find_by_name('original')
    assert fetched_role is not None
    assert fetched_role.id == existing_role.id
    assert fetched_role.name == 'original'
    assert fetched_role.created_by == 'original_user'
    assert fetched_role.updated_by == 'original_user'


def test_role_service_crud_find_by_name_non_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find by id (non existing)
    non_existing_role = role_service.find_by_name('invalid-id')
    assert non_existing_role is None


def test_role_service_crud_find_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find (existing)
    fetched_role_result = role_service.find(keyword='original', limit=100, offset=0)
    assert fetched_role_result.count == 1
    fetched_role = fetched_role_result.rows[0]
    assert fetched_role is not None
    assert fetched_role.id == existing_role.id
    assert fetched_role.name == 'original'
    assert fetched_role.created_by == 'original_user'
    assert fetched_role.updated_by == 'original_user'


def test_role_service_crud_find_non_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find (non existing)
    non_existing_role_result = role_service.find(keyword='invalid-keyword', limit=100, offset=0)
    assert non_existing_role_result.count == 0


def test_role_service_crud_find_pagination():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    for index in range(7):
        insert_role_data(role_repo, index)
    # test find (page 1)
    fetched_role_result = role_service.find(keyword='original', limit=3, offset=0)
    assert len(fetched_role_result.rows) == 3
    assert fetched_role_result.count == 7
    # test find (page 2)
    fetched_role_result = role_service.find(keyword='original', limit=3, offset=3)
    assert len(fetched_role_result.rows) == 3
    assert fetched_role_result.count == 7
    # test find (page 3)
    fetched_role_result = role_service.find(keyword='original', limit=3, offset=6)
    assert len(fetched_role_result.rows) == 1
    assert fetched_role_result.count == 7


def test_role_service_crud_insert():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare insert
    inserted_role_data = create_role_data()
    inserted_role_data.name = 'original'
    inserted_role_data.created_by = 'original_user'
    inserted_role_data.updated_by = 'original_user'
    # test insert
    inserted_role = role_service.insert(inserted_role_data)
    assert inserted_role is not None
    assert inserted_role.id != '' 
    assert inserted_role.name == 'original'
    assert inserted_role.created_by == 'original_user'
    assert inserted_role.updated_by == 'original_user'
    assert role_repo.count(keyword='') == 1


def test_role_service_crud_update_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test update (existing)
    updated_role_data = create_role_data()
    updated_role_data.name = 'updated'
    updated_role_data.updated_by = 'editor'
    updated_role = role_service.update(existing_role.id, updated_role_data)
    assert updated_role is not None
    assert updated_role.id == existing_role.id
    assert updated_role.name == 'updated'
    assert updated_role.created_by == 'original_user'
    assert updated_role.updated_by == 'editor'
    assert role_repo.count(keyword='') == 1


def test_role_service_crud_update_non_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    insert_role_data(role_repo)
    # test update (non existing)
    updated_role_data = create_role_data()
    updated_role_data.name = 'updated'
    updated_role_data.updated_by = 'editor'
    updated_role = role_service.update('invalid-id', updated_role_data)
    assert updated_role == None
    assert role_repo.count(keyword='') == 1


def test_role_service_crud_delete_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find by id (existing)
    deleted_role = role_service.delete(existing_role.id)
    assert deleted_role is not None
    assert deleted_role.id == existing_role.id
    assert deleted_role.name == 'original'
    assert deleted_role.created_by == 'original_user'
    assert deleted_role.updated_by == 'original_user'
    assert role_repo.count(keyword='') == 0


def test_role_service_crud_delete_non_existing():
    role_service, role_repo, _, _ = init_test_role_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find by id (non existing)
    deleted_role = role_service.delete('invalid-id')
    assert deleted_role is None
    assert role_repo.count(keyword='') == 1
