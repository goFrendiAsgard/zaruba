from modules.cms.content.test_contentService_util import create_content_data, insert_content_data, init_test_content_service_components
from modules.auth.user.test_defaultUserService_util import AUTHORIZED_ACTIVE_USER


def test_content_service_crud_find_by_id_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    existing_content = insert_content_data(content_repo, content_type.id)
    # test find by id (existing)
    fetched_content = content_service.find_by_id(existing_content.id, AUTHORIZED_ACTIVE_USER)
    assert fetched_content is not None
    assert fetched_content.id == existing_content.id
    assert fetched_content.title == 'content'
    assert fetched_content.created_by == 'original_user'
    assert fetched_content.updated_by == 'original_user'


def test_content_service_crud_find_by_id_non_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    insert_content_data(content_repo, content_type.id)
    # test find by id (non existing)
    is_error = False
    try:
        content_service.find_by_id('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_content_service_crud_find_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    existing_content = insert_content_data(content_repo, content_type.id)
    # test find (existing)
    fetched_content_result = content_service.find(keyword='content', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert fetched_content_result.count == 1
    fetched_content = fetched_content_result.rows[0]
    assert fetched_content is not None
    assert fetched_content.id == existing_content.id
    assert fetched_content.title == 'content'
    assert fetched_content.created_by == 'original_user'
    assert fetched_content.updated_by == 'original_user'


def test_content_service_crud_find_non_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    insert_content_data(content_repo, content_type.id)
    # test find (non existing)
    non_existing_content_result = content_service.find(keyword='invalid-keyword', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert non_existing_content_result.count == 0


def test_content_service_crud_find_pagination():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    for index in range(7):
        insert_content_data(content_repo, content_type.id, index)
    # test find (page 1)
    fetched_content_result = content_service.find(keyword='content', limit=3, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_content_result.rows) == 3
    assert fetched_content_result.count == 7
    # test find (page 2)
    fetched_content_result = content_service.find(keyword='content', limit=3, offset=3, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_content_result.rows) == 3
    assert fetched_content_result.count == 7
    # test find (page 3)
    fetched_content_result = content_service.find(keyword='content', limit=3, offset=6, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_content_result.rows) == 1
    assert fetched_content_result.count == 7


def test_content_service_crud_insert():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare insert
    inserted_content_data = create_content_data(content_type.id)
    inserted_content_data.title = 'content'
    inserted_content_data.created_by = 'original_user'
    inserted_content_data.updated_by = 'original_user'
    # test insert
    inserted_content = content_service.insert(inserted_content_data, AUTHORIZED_ACTIVE_USER)
    assert inserted_content is not None
    assert inserted_content.id != '' 
    assert inserted_content.title == 'content'
    assert inserted_content.created_by == AUTHORIZED_ACTIVE_USER.id
    assert inserted_content.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert content_repo.count(keyword='') == 1


def test_content_service_crud_update_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    existing_content = insert_content_data(content_repo, content_type.id)
    # test update (existing)
    updated_content_data = create_content_data(content_type.id)
    updated_content_data.title = 'updated'
    updated_content_data.updated_by = 'editor'
    updated_content = content_service.update(existing_content.id, updated_content_data, AUTHORIZED_ACTIVE_USER)
    assert updated_content is not None
    assert updated_content.id == existing_content.id
    assert updated_content.title == 'updated'
    assert updated_content.created_by == 'original_user'
    assert updated_content.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert content_repo.count(keyword='') == 1


def test_content_service_crud_update_non_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    insert_content_data(content_repo, content_type.id)
    # test update (non existing)
    updated_content_data = create_content_data(content_type.id)
    updated_content_data.title = 'updated'
    updated_content_data.updated_by = 'editor'
    is_error = False
    try:
        content_service.update('invalid-id', updated_content_data, AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert content_repo.count(keyword='') == 1


def test_content_service_crud_delete_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    existing_content = insert_content_data(content_repo, content_type.id)
    # test find by id (existing)
    deleted_content = content_service.delete(existing_content.id, AUTHORIZED_ACTIVE_USER)
    assert deleted_content is not None
    assert deleted_content.id == existing_content.id
    assert deleted_content.title == 'content'
    assert deleted_content.created_by == 'original_user'
    assert deleted_content.updated_by == 'original_user'
    assert content_repo.count(keyword='') == 0


def test_content_service_crud_delete_non_existing():
    content_service, _, content_repo, content_type, _, _ = init_test_content_service_components()
    # prepare repo
    insert_content_data(content_repo, content_type.id)
    # test find by id (non existing)
    is_error = False
    try:
        content_service.delete('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert content_repo.count(keyword='') == 1
