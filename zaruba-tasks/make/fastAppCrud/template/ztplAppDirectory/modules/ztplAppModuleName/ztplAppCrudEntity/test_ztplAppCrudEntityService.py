from typing import Optional, List
from modules.ztplAppModuleName.ztplAppCrudEntity.ztplAppCrudEntityService import ZtplAppCrudEntityService
from modules.ztplAppModuleName.ztplAppCrudEntity.repos.dbZtplAppCrudEntityRepo import DBZtplAppCrudEntityRepo
from schemas.ztplAppCrudEntity import ZtplAppCrudEntityData, ZtplAppCrudEntityData
from helpers.transport import LocalRPC, LocalMessageBus

from sqlalchemy import create_engine

################################################
# -- ⚙️ Helpers
################################################

def create_ztpl_app_crud_entity_data():
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_ztpl_app_crud_entity_data = ZtplAppCrudEntityData(
        created_by=''
    )
    return dummy_ztpl_app_crud_entity_data


################################################
# -- 🧪 Test
################################################

def test_ztpl_app_crud_entity_service_crud():
    engine = create_engine('sqlite://', echo=False)
    ztpl_app_crud_entity_repo = DBZtplAppCrudEntityRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mb, rpc, ztpl_app_crud_entity_repo)

    # prepare insert
    inserted_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    inserted_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'original'
    inserted_ztpl_app_crud_entity_data.created_by = 'original_user'
    inserted_ztpl_app_crud_entity_data.updated_by = 'original_user'
    # test insert
    inserted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(inserted_ztpl_app_crud_entity_data)
    assert inserted_ztpl_app_crud_entity is not None
    assert inserted_ztpl_app_crud_entity.id != '' 
    assert inserted_ztpl_app_crud_entity.ztplAppCrudFirstField == 'original'
    assert inserted_ztpl_app_crud_entity.created_by == 'original_user'
    assert inserted_ztpl_app_crud_entity.updated_by == 'original_user'

    # test find by id (existing, after insert)
    existing_ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(inserted_ztpl_app_crud_entity.id)
    assert existing_ztpl_app_crud_entity is not None
    assert existing_ztpl_app_crud_entity.id == inserted_ztpl_app_crud_entity.id
    assert existing_ztpl_app_crud_entity.ztplAppCrudFirstField == inserted_ztpl_app_crud_entity.ztplAppCrudFirstField
    assert existing_ztpl_app_crud_entity.created_by == inserted_ztpl_app_crud_entity.created_by
    assert existing_ztpl_app_crud_entity.updated_by == inserted_ztpl_app_crud_entity.updated_by

    # test find by id (non existing)
    non_existing_ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id('invalid_id')
    assert non_existing_ztpl_app_crud_entity is None

    # prepare update (existing)
    updated_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    updated_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'updated'
    updated_ztpl_app_crud_entity_data.updated_by = 'editor'
    # test update (existing)
    updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update(inserted_ztpl_app_crud_entity.id, updated_ztpl_app_crud_entity_data)
    assert updated_ztpl_app_crud_entity is not None
    assert updated_ztpl_app_crud_entity.id == inserted_ztpl_app_crud_entity.id
    assert updated_ztpl_app_crud_entity.ztplAppCrudFirstField == 'updated'
    assert updated_ztpl_app_crud_entity.created_by == 'original_user'
    assert updated_ztpl_app_crud_entity.updated_by == 'editor'

    # test update (non existing)
    non_existing_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update('invalid_id', updated_ztpl_app_crud_entity_data)
    assert non_existing_ztpl_app_crud_entity is None

    # test find by id (existing, after insert)
    existing_ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(updated_ztpl_app_crud_entity.id)
    assert existing_ztpl_app_crud_entity is not None
    assert existing_ztpl_app_crud_entity.id == inserted_ztpl_app_crud_entity.id
    assert existing_ztpl_app_crud_entity.ztplAppCrudFirstField == 'updated'
    assert existing_ztpl_app_crud_entity.created_by == 'original_user'
    assert existing_ztpl_app_crud_entity.updated_by == 'editor'

    # test find (before delete, correct keyword)
    existing_result = ztpl_app_crud_entity_service.find(keyword='updated', limit=10, offset=0)
    assert existing_result.count == 1
    assert len(existing_result.rows) == 1
    assert existing_result.rows[0].id == inserted_ztpl_app_crud_entity.id

    # test find (before delete, incorrect keyword)
    non_existing_result = ztpl_app_crud_entity_service.find(keyword='incorrect', limit=10, offset=0)
    assert non_existing_result.count == 0
    assert len(non_existing_result.rows) == 0

    # test delete existing
    deleted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete(inserted_ztpl_app_crud_entity.id)
    assert deleted_ztpl_app_crud_entity is not None
    assert deleted_ztpl_app_crud_entity.id == inserted_ztpl_app_crud_entity.id
    assert deleted_ztpl_app_crud_entity.ztplAppCrudFirstField == 'updated'
    assert deleted_ztpl_app_crud_entity.created_by == 'original_user'
    assert deleted_ztpl_app_crud_entity.updated_by == 'editor'

    # test delete (non existing)
    non_existing_ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete('invalid_id')
    assert non_existing_ztpl_app_crud_entity is None

    # test find (after delete, no keyword)
    non_existing_result = ztpl_app_crud_entity_service.find(keyword='', limit=10, offset=0)
    assert non_existing_result.count == 0
    assert len(non_existing_result.rows) == 0
    