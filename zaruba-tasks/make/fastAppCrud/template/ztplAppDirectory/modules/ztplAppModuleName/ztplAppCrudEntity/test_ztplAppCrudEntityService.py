from modules.ztplAppModuleName.ztplAppCrudEntity.test_ztplAppCrudEntityService_util import create_ztpl_app_crud_entity_data, insert_ztpl_app_crud_entity_data, init_test_ztpl_app_crud_entity_service_components
from modules.auth.user.test_defaultUserService_util import AUTHORIZED_ACTIVE_USER


def test_ztpl_app_crud_entity_service_crud_find_by_id_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (existing)
    fetched_ztpl_app_crud_entity = ztpl_app_crud_entity_service.find_by_id(existing_ztpl_app_crud_entity.id, AUTHORIZED_ACTIVE_USER)
    assert fetched_ztpl_app_crud_entity is not None
    assert fetched_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert fetched_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert fetched_ztpl_app_crud_entity.created_by == 'original_user'
    assert fetched_ztpl_app_crud_entity.updated_by == 'original_user'


def test_ztpl_app_crud_entity_service_crud_find_by_id_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (non existing)
    is_error = False
    try:
        ztpl_app_crud_entity_service.find_by_id('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_ztpl_app_crud_entity_service_crud_find_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find (existing)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert fetched_ztpl_app_crud_entity_result.count == 1
    fetched_ztpl_app_crud_entity = fetched_ztpl_app_crud_entity_result.rows[0]
    assert fetched_ztpl_app_crud_entity is not None
    assert fetched_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert fetched_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert fetched_ztpl_app_crud_entity.created_by == 'original_user'
    assert fetched_ztpl_app_crud_entity.updated_by == 'original_user'


def test_ztpl_app_crud_entity_service_crud_find_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find (non existing)
    non_existing_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='invalid-keyword', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert non_existing_ztpl_app_crud_entity_result.count == 0


def test_ztpl_app_crud_entity_service_crud_find_pagination():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    for index in range(7):
        insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo, index)
    # test find (page 1)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=3, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_ztpl_app_crud_entity_result.rows) == 3
    assert fetched_ztpl_app_crud_entity_result.count == 7
    # test find (page 2)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=3, offset=3, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_ztpl_app_crud_entity_result.rows) == 3
    assert fetched_ztpl_app_crud_entity_result.count == 7
    # test find (page 3)
    fetched_ztpl_app_crud_entity_result = ztpl_app_crud_entity_service.find(keyword='ztplAppCrudEntity', limit=3, offset=6, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_ztpl_app_crud_entity_result.rows) == 1
    assert fetched_ztpl_app_crud_entity_result.count == 7


def test_ztpl_app_crud_entity_service_crud_insert():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare insert
    inserted_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    inserted_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'ztplAppCrudEntity'
    inserted_ztpl_app_crud_entity_data.created_by = 'original_user'
    inserted_ztpl_app_crud_entity_data.updated_by = 'original_user'
    # test insert
    inserted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.insert(inserted_ztpl_app_crud_entity_data, AUTHORIZED_ACTIVE_USER)
    assert inserted_ztpl_app_crud_entity is not None
    assert inserted_ztpl_app_crud_entity.id != '' 
    assert inserted_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert inserted_ztpl_app_crud_entity.created_by == AUTHORIZED_ACTIVE_USER.id
    assert inserted_ztpl_app_crud_entity.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1


def test_ztpl_app_crud_entity_service_crud_update_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test update (existing)
    updated_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    updated_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'updated'
    updated_ztpl_app_crud_entity_data.updated_by = 'editor'
    updated_ztpl_app_crud_entity = ztpl_app_crud_entity_service.update(existing_ztpl_app_crud_entity.id, updated_ztpl_app_crud_entity_data, AUTHORIZED_ACTIVE_USER)
    assert updated_ztpl_app_crud_entity is not None
    assert updated_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert updated_ztpl_app_crud_entity.ztplAppCrudFirstField == 'updated'
    assert updated_ztpl_app_crud_entity.created_by == 'original_user'
    assert updated_ztpl_app_crud_entity.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1


def test_ztpl_app_crud_entity_service_crud_update_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test update (non existing)
    updated_ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    updated_ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'updated'
    updated_ztpl_app_crud_entity_data.updated_by = 'editor'
    is_error = False
    try:
        ztpl_app_crud_entity_service.update('invalid-id', updated_ztpl_app_crud_entity_data, AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1


def test_ztpl_app_crud_entity_service_crud_delete_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    existing_ztpl_app_crud_entity = insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (existing)
    deleted_ztpl_app_crud_entity = ztpl_app_crud_entity_service.delete(existing_ztpl_app_crud_entity.id, AUTHORIZED_ACTIVE_USER)
    assert deleted_ztpl_app_crud_entity is not None
    assert deleted_ztpl_app_crud_entity.id == existing_ztpl_app_crud_entity.id
    assert deleted_ztpl_app_crud_entity.ztplAppCrudFirstField == 'ztplAppCrudEntity'
    assert deleted_ztpl_app_crud_entity.created_by == 'original_user'
    assert deleted_ztpl_app_crud_entity.updated_by == 'original_user'
    assert ztpl_app_crud_entity_repo.count(keyword='') == 0


def test_ztpl_app_crud_entity_service_crud_delete_non_existing():
    ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, _, _ = init_test_ztpl_app_crud_entity_service_components()
    # prepare repo
    insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo)
    # test find by id (non existing)
    is_error = False
    try:
        ztpl_app_crud_entity_service.delete('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert ztpl_app_crud_entity_repo.count(keyword='') == 1
