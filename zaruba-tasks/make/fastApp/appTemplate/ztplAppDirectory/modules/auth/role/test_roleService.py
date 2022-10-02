from typing import Optional, List
from modules.auth.role.roleService import RoleService
from modules.auth.role.repos.dbRoleRepo import DBRoleRepo
from schemas.role import RoleData
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine

################################################
# -- ⚙️ Helpers
################################################

def create_role_data():
    dummy_role_data = RoleData(
        name='',
        permissions=[],
    )
    return dummy_role_data


################################################
# -- 🧪 Test
################################################

def test_role_service():
    engine = create_engine('sqlite://', echo=True)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    role_service = RoleService(mb, rpc, role_repo)

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

    # test find by id (existing, after insert)
    existing_role = role_service.find_by_id(inserted_role.id)
    assert existing_role is not None
    assert existing_role.id == inserted_role.id
    assert existing_role.name == inserted_role.name
    assert existing_role.created_by == inserted_role.created_by
    assert existing_role.updated_by == inserted_role.updated_by

    # test find by name (existing, after insert)
    existing_role = role_service.find_by_name('original')
    assert existing_role is not None
    assert existing_role.id == inserted_role.id
    assert existing_role.name == inserted_role.name
    assert existing_role.created_by == inserted_role.created_by
    assert existing_role.updated_by == inserted_role.updated_by

    # test find by id (non existing)
    non_existing_role = role_service.find_by_id('invalid_id')
    assert non_existing_role is None

    # test find by name (non existing)
    non_existing_role = role_service.find_by_name('invalid_name')
    assert non_existing_role is None

    # prepare update (existing)
    updated_role_data = create_role_data()
    updated_role_data.name = 'updated'
    updated_role_data.updated_by = 'editor'
    # test update (existing)
    updated_role = role_service.update(inserted_role.id, updated_role_data)
    assert updated_role is not None
    assert updated_role.id == inserted_role.id
    assert updated_role.name == 'updated'
    assert updated_role.created_by == 'original_user'
    assert updated_role.updated_by == 'editor'

    # test update (non existing)
    non_existing_role = role_service.update('invalid_id', updated_role_data)
    assert non_existing_role is None

    # test find by id (existing, after insert)
    existing_role = role_service.find_by_id(updated_role.id)
    assert existing_role is not None
    assert existing_role.id == inserted_role.id
    assert existing_role.name == 'updated'
    assert existing_role.created_by == 'original_user'
    assert existing_role.updated_by == 'editor'

    # test find (before delete, correct keyword)
    existing_result = role_service.find(keyword='updated', limit=10, offset=0)
    assert existing_result.count == 1
    assert len(existing_result.rows) == 1
    assert existing_result.rows[0].id == inserted_role.id

    # test find (before delete, incorrect keyword)
    non_existing_result = role_service.find(keyword='incorrect', limit=10, offset=0)
    assert non_existing_result.count == 0
    assert len(non_existing_result.rows) == 0

    # test delete existing
    deleted_role = role_service.delete(inserted_role.id)
    assert deleted_role is not None
    assert deleted_role.id == inserted_role.id
    assert deleted_role.name == 'updated'
    assert deleted_role.created_by == 'original_user'
    assert deleted_role.updated_by == 'editor'

    # test delete (non existing)
    non_existing_role = role_service.delete('invalid_id')
    assert non_existing_role is None

    # test find (after delete, correct keyword)
    non_existing_result = role_service.find(keyword='updated', limit=10, offset=0)
    assert non_existing_result.count == 0
    assert len(non_existing_result.rows) == 0
    