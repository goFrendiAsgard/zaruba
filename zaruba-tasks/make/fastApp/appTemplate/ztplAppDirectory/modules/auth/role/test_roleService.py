from modules.auth.role.test_roleService_util import create_role_data, insert_role_data, init_test_role_service_components
from modules.auth.user.test_defaultUserService_util import AUTHORIZED_ACTIVE_USER


def test_role_service_crud_find_by_id_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find by id (existing)
    fetched_role = role_service.find_by_id(existing_role.id, AUTHORIZED_ACTIVE_USER)
    assert fetched_role is not None
    assert fetched_role.id == existing_role.id
    assert fetched_role.name == 'original'
    assert fetched_role.created_by == 'original_user'
    assert fetched_role.updated_by == 'original_user'


def test_role_service_crud_find_by_id_non_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find by id (non existing)
    is_error = False
    try:
        role_service.find_by_id('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_role_service_crud_find_by_name_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find by id (existing)
    fetched_role = role_service.find_by_name('original', AUTHORIZED_ACTIVE_USER)
    assert fetched_role is not None
    assert fetched_role.id == existing_role.id
    assert fetched_role.name == 'original'
    assert fetched_role.created_by == 'original_user'
    assert fetched_role.updated_by == 'original_user'


def test_role_service_crud_find_by_name_non_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find by id (non existing)
    is_error = False
    try:
        role_service.find_by_name('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_role_service_crud_find_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find (existing)
    fetched_role_result = role_service.find(keyword='original', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert fetched_role_result.count == 1
    fetched_role = fetched_role_result.rows[0]
    assert fetched_role is not None
    assert fetched_role.id == existing_role.id
    assert fetched_role.name == 'original'
    assert fetched_role.created_by == 'original_user'
    assert fetched_role.updated_by == 'original_user'


def test_role_service_crud_find_non_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find (non existing)
    non_existing_role_result = role_service.find(keyword='invalid-keyword', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert non_existing_role_result.count == 0


def test_role_service_crud_find_pagination():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    for index in range(7):
        insert_role_data(role_repo, index)
    # test find (page 1)
    fetched_role_result = role_service.find(keyword='original', limit=3, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_role_result.rows) == 3
    assert fetched_role_result.count == 7
    # test find (page 2)
    fetched_role_result = role_service.find(keyword='original', limit=3, offset=3, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_role_result.rows) == 3
    assert fetched_role_result.count == 7
    # test find (page 3)
    fetched_role_result = role_service.find(keyword='original', limit=3, offset=6, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_role_result.rows) == 1
    assert fetched_role_result.count == 7


def test_role_service_crud_insert():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare insert
    inserted_role_data = create_role_data()
    inserted_role_data.name = 'original'
    inserted_role_data.created_by = 'original_user'
    inserted_role_data.updated_by = 'original_user'
    # test insert
    inserted_role = role_service.insert(inserted_role_data, AUTHORIZED_ACTIVE_USER)
    assert inserted_role is not None
    assert inserted_role.id != '' 
    assert inserted_role.name == 'original'
    assert inserted_role.created_by == AUTHORIZED_ACTIVE_USER.id
    assert inserted_role.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert role_repo.count(keyword='') == 1


def test_role_service_crud_update_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test update (existing)
    updated_role_data = create_role_data()
    updated_role_data.name = 'updated'
    updated_role_data.updated_by = 'editor'
    updated_role = role_service.update(existing_role.id, updated_role_data, AUTHORIZED_ACTIVE_USER)
    assert updated_role is not None
    assert updated_role.id == existing_role.id
    assert updated_role.name == 'updated'
    assert updated_role.created_by == 'original_user'
    assert updated_role.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert role_repo.count(keyword='') == 1


def test_role_service_crud_update_non_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    insert_role_data(role_repo)
    # test update (non existing)
    updated_role_data = create_role_data()
    updated_role_data.name = 'updated'
    updated_role_data.updated_by = 'editor'
    is_error = False
    try:
        role_service.update('invalid-id', updated_role_data, AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert role_repo.count(keyword='') == 1


def test_role_service_crud_delete_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    existing_role = insert_role_data(role_repo)
    # test find by id (existing)
    deleted_role = role_service.delete(existing_role.id, AUTHORIZED_ACTIVE_USER)
    assert deleted_role is not None
    assert deleted_role.id == existing_role.id
    assert deleted_role.name == 'original'
    assert deleted_role.created_by == 'original_user'
    assert deleted_role.updated_by == 'original_user'
    assert role_repo.count(keyword='') == 0


def test_role_service_crud_delete_non_existing():
    role_service, role_repo, _, _ = init_test_role_service_components()
    # prepare repo
    insert_role_data(role_repo)
    # test find by id (non existing)
    is_error = False
    try:
        role_service.delete('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert role_repo.count(keyword='') == 1
