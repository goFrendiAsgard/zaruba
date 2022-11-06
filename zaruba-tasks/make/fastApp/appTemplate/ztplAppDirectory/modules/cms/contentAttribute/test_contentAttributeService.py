from modules.cms.contentAttribute.test_contentAttributeService_util import create_content_attribute_data, insert_content_attribute_data, init_test_content_attribute_service_components
from modules.auth.user.test_defaultUserService_util import AUTHORIZED_ACTIVE_USER


def test_content_attribute_service_crud_find_by_id_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    existing_content_attribute = insert_content_attribute_data(content_attribute_repo)
    # test find by id (existing)
    fetched_content_attribute = content_attribute_service.find_by_id(existing_content_attribute.id, AUTHORIZED_ACTIVE_USER)
    assert fetched_content_attribute is not None
    assert fetched_content_attribute.id == existing_content_attribute.id
    assert fetched_content_attribute.content_id == 'contentAttribute'
    assert fetched_content_attribute.created_by == 'original_user'
    assert fetched_content_attribute.updated_by == 'original_user'


def test_content_attribute_service_crud_find_by_id_non_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    insert_content_attribute_data(content_attribute_repo)
    # test find by id (non existing)
    is_error = False
    try:
        content_attribute_service.find_by_id('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_content_attribute_service_crud_find_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    existing_content_attribute = insert_content_attribute_data(content_attribute_repo)
    # test find (existing)
    fetched_content_attribute_result = content_attribute_service.find(keyword='contentAttribute', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert fetched_content_attribute_result.count == 1
    fetched_content_attribute = fetched_content_attribute_result.rows[0]
    assert fetched_content_attribute is not None
    assert fetched_content_attribute.id == existing_content_attribute.id
    assert fetched_content_attribute.content_id == 'contentAttribute'
    assert fetched_content_attribute.created_by == 'original_user'
    assert fetched_content_attribute.updated_by == 'original_user'


def test_content_attribute_service_crud_find_non_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    insert_content_attribute_data(content_attribute_repo)
    # test find (non existing)
    non_existing_content_attribute_result = content_attribute_service.find(keyword='invalid-keyword', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert non_existing_content_attribute_result.count == 0


def test_content_attribute_service_crud_find_pagination():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    for index in range(7):
        insert_content_attribute_data(content_attribute_repo, index)
    # test find (page 1)
    fetched_content_attribute_result = content_attribute_service.find(keyword='contentAttribute', limit=3, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_content_attribute_result.rows) == 3
    assert fetched_content_attribute_result.count == 7
    # test find (page 2)
    fetched_content_attribute_result = content_attribute_service.find(keyword='contentAttribute', limit=3, offset=3, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_content_attribute_result.rows) == 3
    assert fetched_content_attribute_result.count == 7
    # test find (page 3)
    fetched_content_attribute_result = content_attribute_service.find(keyword='contentAttribute', limit=3, offset=6, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_content_attribute_result.rows) == 1
    assert fetched_content_attribute_result.count == 7


def test_content_attribute_service_crud_insert():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare insert
    inserted_content_attribute_data = create_content_attribute_data()
    inserted_content_attribute_data.content_id = 'contentAttribute'
    inserted_content_attribute_data.created_by = 'original_user'
    inserted_content_attribute_data.updated_by = 'original_user'
    # test insert
    inserted_content_attribute = content_attribute_service.insert(inserted_content_attribute_data, AUTHORIZED_ACTIVE_USER)
    assert inserted_content_attribute is not None
    assert inserted_content_attribute.id != '' 
    assert inserted_content_attribute.content_id == 'contentAttribute'
    assert inserted_content_attribute.created_by == AUTHORIZED_ACTIVE_USER.id
    assert inserted_content_attribute.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert content_attribute_repo.count(keyword='') == 1


def test_content_attribute_service_crud_update_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    existing_content_attribute = insert_content_attribute_data(content_attribute_repo)
    # test update (existing)
    updated_content_attribute_data = create_content_attribute_data()
    updated_content_attribute_data.content_id = 'updated'
    updated_content_attribute_data.updated_by = 'editor'
    updated_content_attribute = content_attribute_service.update(existing_content_attribute.id, updated_content_attribute_data, AUTHORIZED_ACTIVE_USER)
    assert updated_content_attribute is not None
    assert updated_content_attribute.id == existing_content_attribute.id
    assert updated_content_attribute.content_id == 'updated'
    assert updated_content_attribute.created_by == 'original_user'
    assert updated_content_attribute.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert content_attribute_repo.count(keyword='') == 1


def test_content_attribute_service_crud_update_non_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    insert_content_attribute_data(content_attribute_repo)
    # test update (non existing)
    updated_content_attribute_data = create_content_attribute_data()
    updated_content_attribute_data.content_id = 'updated'
    updated_content_attribute_data.updated_by = 'editor'
    is_error = False
    try:
        content_attribute_service.update('invalid-id', updated_content_attribute_data, AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert content_attribute_repo.count(keyword='') == 1


def test_content_attribute_service_crud_delete_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    existing_content_attribute = insert_content_attribute_data(content_attribute_repo)
    # test find by id (existing)
    deleted_content_attribute = content_attribute_service.delete(existing_content_attribute.id, AUTHORIZED_ACTIVE_USER)
    assert deleted_content_attribute is not None
    assert deleted_content_attribute.id == existing_content_attribute.id
    assert deleted_content_attribute.content_id == 'contentAttribute'
    assert deleted_content_attribute.created_by == 'original_user'
    assert deleted_content_attribute.updated_by == 'original_user'
    assert content_attribute_repo.count(keyword='') == 0


def test_content_attribute_service_crud_delete_non_existing():
    content_attribute_service, content_attribute_repo, _, _ = init_test_content_attribute_service_components()
    # prepare repo
    insert_content_attribute_data(content_attribute_repo)
    # test find by id (non existing)
    is_error = False
    try:
        content_attribute_service.delete('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert content_attribute_repo.count(keyword='') == 1
