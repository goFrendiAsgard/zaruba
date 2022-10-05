from typing import Optional, Tuple
from modules.ztplAppModuleName.ztplAppCrudEntity.ztplAppCrudEntityService import ZtplAppCrudEntityService
from modules.ztplAppModuleName.ztplAppCrudEntity.repos.dbZtplAppCrudEntityRepo import DBZtplAppCrudEntityRepo
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
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


def init_test_ztpl_app_crud_entity_components() -> Tuple[ZtplAppCrudEntityService, DBZtplAppCrudEntityRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    ztpl_app_crud_entity_repo = DBZtplAppCrudEntityRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mb, rpc, ztpl_app_crud_entity_repo)
    return ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, mb, rpc


def insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo: DBZtplAppCrudEntityRepo, index: Optional[int] = None) -> ZtplAppCrudEntity:
    ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'ztplAppCrudEntity' if index is None else 'ztplAppCrudEntity-{index}'.format(index=index)
    ztpl_app_crud_entity_data.created_by = 'original_user'
    ztpl_app_crud_entity_data.updated_by = 'original_user'
    return ztpl_app_crud_entity_repo.insert(ztpl_app_crud_entity_data)


################################################
# -- 🧪 Test
################################################

def test_ztpl_app_crud_entity_service_crud_find_by_id_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (existing)
    fetched_ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(existing_ztpl_app_crud_entity.id)
    assert fetched_ztpl_app_crud_entity is not None
    assert fetched_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert fetched_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert fetched_ztpl_app_crud_entity.created_by == 'original_user'
    assert fetched_ztpl_app_crud_entity.updated_by == 'original_user'


def test_ztpl_app_crud_entity_service_crud_find_by_id_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (non existing)
    non_existing_ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id('invalid-id')
    assert non_existing_ztpl_app_crud_entity is None


def test_ztpl_app_crud_entity_service_crud_find_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find (existing)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=100, offset=0)
    assert fetched_ztpl_app_crud_entity_result.count == 1
    fetched_ztpl_app_crud_entity = fetched_ztpl_app_crud_entity_result.rows[0]
    assert fetched_ztpl_app_crud_entity is not None
    assert fetched_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert fetched_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert fetched_ztpl_app_crud_entity.created_by == 'original_user'
    assert fetched_ztpl_app_crud_entity.updated_by == 'original_user'


def test_ztpl_app_crud_entity_service_crud_find_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find (non existing)
    non_existing_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='invalid-keyword', limit=100, offset=0)
    assert non_existing_ztpl_app_crud_entity_result.count == 0


def test_ztpl_app_crud_entity_service_crud_find_pagination():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    for index in range(7):
        insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo, index)
    # test find (page 1)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=3, offset=0)
    assert len(fetched_ztpl_app_crud_entity_result.rows) == 3
    assert fetched_ztpl_app_crud_entity_result.count == 7
    # test find (page 2)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=3, offset=3)
    assert len(fetched_ztpl_app_crud_entity_result.rows) == 3
    assert fetched_ztpl_app_crud_entity_result.count == 7
    # test find (page 3)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=3, offset=6)
    assert len(fetched_ztpl_app_crud_entity_result.rows) == 1
    assert fetched_ztpl_app_crud_entity_result.count == 7


def test_ztpl_app_crud_entity_service_crud_insert():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare insert
    inserted_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    inserted_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'ztplAppCrudEntity'
    inserted_ztpl_app_crud_entity_data.created_by = 'original_user'
    inserted_ztpl_app_crud_entity_data.updated_by = 'original_user'
    # test insert
    inserted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(inserted_ztpl_app_crud_entity_data)
    assert inserted_ztpl_app_crud_entity is not None
    assert inserted_ztpl_app_crud_entity.id != '' 
    assert inserted_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert inserted_ztpl_app_crud_entity.created_by == 'original_user'
    assert inserted_ztpl_app_crud_entity.updated_by == 'original_user'
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1


def test_ztpl_app_crud_entity_service_crud_update_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test update (existing)
    updated_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    updated_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'updated'
    updated_ztpl_app_crud_entity_data.updated_by = 'editor'
    updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update(existing_ztpl_app_crud_entity.id, updated_ztpl_app_crud_entity_data)
    assert updated_ztpl_app_crud_entity is not None
    assert updated_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert updated_ztpl_app_crud_entity.ztplAppCrudFirstField == 'updated'
    assert updated_ztpl_app_crud_entity.created_by == 'original_user'
    assert updated_ztpl_app_crud_entity.updated_by == 'editor'
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1


def test_ztpl_app_crud_entity_service_crud_update_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test update (non existing)
    updated_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    updated_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'updated'
    updated_ztpl_app_crud_entity_data.updated_by = 'editor'
    updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update('invalid-id', updated_ztpl_app_crud_entity_data)
    assert updated_ztpl_app_crud_entity == None
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1


def test_ztpl_app_crud_entity_service_crud_delete_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (existing)
    deleted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete(existing_ztpl_app_crud_entity.id)
    assert deleted_ztpl_app_crud_entity is not None
    assert deleted_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert deleted_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert deleted_ztpl_app_crud_entity.created_by == 'original_user'
    assert deleted_ztpl_app_crud_entity.updated_by == 'original_user'
    assert ztpl_app_crud_entity_repo.count(keyword='') == 0


def test_ztpl_app_crud_entity_service_crud_delete_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (non existing)
    deleted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete('invalid-id')
    assert deleted_ztpl_app_crud_entity is None
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1
