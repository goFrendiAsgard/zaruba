from module.auth.role.test_role_service import create_role_data
from module.auth.user.test_default_user_service_util import create_user_data, init_test_default_user_service_components, init_user_data, AUTHORIZED_ACTIVE_USER


def test_user_service_crud_find_by_id_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test find by id (existing)
    fetched_user = user_service.find_by_id(existing_user.id, AUTHORIZED_ACTIVE_USER)
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_id_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find by id (non existing)
    is_error = False
    try:
        user_service.find_by_id('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_user_service_crud_find_by_username_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test find by id (existing)
    fetched_user = user_service.find_by_username(existing_user.username, AUTHORIZED_ACTIVE_USER)
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_username_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find by id (non existing)
    is_error = False
    try:
        user_service.find_by_username('invalid-user', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_user_service_crud_find_by_username_and_password_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find by id (existing)
    fetched_user = user_service.find_by_identity_and_password(existing_user.username, 'password', AUTHORIZED_ACTIVE_USER)
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_email_and_password_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find by id (existing)
    fetched_user = user_service.find_by_identity_and_password(existing_user.email, 'password', AUTHORIZED_ACTIVE_USER)
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_by_identity_and_password_invalid_password():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find by id (non existing)
    is_error = False
    try:
        user_service.find_by_identity_and_password(existing_user.username, 'invalid-password', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_user_service_crud_find_by_identity_and_password_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    init_user_data(user_repo, password='password')
    # test find by id (non existing)
    is_error = False
    try:
        user_service.find_by_identity_and_password('invalid-user', 'invalid-password', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error


def test_user_service_crud_find_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo, password='password')
    # test find (existing)
    fetched_user_result = user_service.find(keyword='user', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert fetched_user_result.count == 1
    fetched_user = fetched_user_result.rows[0]
    assert fetched_user is not None
    assert fetched_user.id == existing_user.id
    assert fetched_user.username == 'user'
    assert fetched_user.created_by == 'original_user'
    assert fetched_user.updated_by == 'original_user'


def test_user_service_crud_find_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find (non existing)
    non_existing_user_result = user_service.find(keyword='invalid-keyword', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert non_existing_user_result.count == 0


def test_user_service_crud_find_pagination():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    for index in range(7):
        init_user_data(user_repo, index)
    # test find (page 1)
    fetched_user_result = user_service.find(keyword='user', limit=3, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_user_result.rows) == 3
    assert fetched_user_result.count == 7
    # test find (page 2)
    fetched_user_result = user_service.find(keyword='user', limit=3, offset=3, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_user_result.rows) == 3
    assert fetched_user_result.count == 7
    # test find (page 3)
    fetched_user_result = user_service.find(keyword='user', limit=3, offset=6, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_user_result.rows) == 1
    assert fetched_user_result.count == 7


def test_user_service_crud_insert():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare insert
    inserted_user_data = create_user_data()
    inserted_user_data.username = 'user'
    # test insert
    inserted_user = user_service.insert(inserted_user_data, AUTHORIZED_ACTIVE_USER)
    assert inserted_user is not None
    assert inserted_user.id != '' 
    assert inserted_user.username == 'user'
    assert inserted_user.created_by == AUTHORIZED_ACTIVE_USER.id
    assert inserted_user.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert user_repo.count(keyword='') == 1


def test_user_service_crud_update_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test update (existing)
    updated_user_data = create_user_data()
    updated_user_data.username = 'updated'
    updated_user = user_service.update(existing_user.id, updated_user_data, AUTHORIZED_ACTIVE_USER)
    assert updated_user is not None
    assert updated_user.id == existing_user.id
    assert updated_user.username == 'updated'
    assert updated_user.created_by == 'original_user'
    assert updated_user.updated_by == AUTHORIZED_ACTIVE_USER.id
    assert user_repo.count(keyword='') == 1


def test_user_service_crud_update_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test update (non existing)
    updated_user_data = create_user_data()
    updated_user_data.username = 'updated'
    updated_user_data.updated_by = 'editor'
    is_error = False
    try:
        user_service.update('invalid-id', updated_user_data, AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert user_repo.count(keyword='') == 1


def test_user_service_crud_delete_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    existing_user = init_user_data(user_repo)
    # test find by id (existing)
    deleted_user = user_service.delete(existing_user.id, AUTHORIZED_ACTIVE_USER)
    assert deleted_user is not None
    assert deleted_user.id == existing_user.id
    assert deleted_user.username == 'user'
    assert deleted_user.created_by == 'original_user'
    assert deleted_user.updated_by == 'original_user'
    assert user_repo.count(keyword='') == 0


def test_user_service_crud_delete_non_existing():
    user_service, _, user_repo, _, _, _ = init_test_default_user_service_components()
    # prepare repo
    init_user_data(user_repo)
    # test find by id (non existing)
    is_error = False
    try:
        user_service.delete('invalid-id', AUTHORIZED_ACTIVE_USER)
    except:
        is_error = True
    assert is_error
    assert user_repo.count(keyword='') == 1


def test_user_service_authorization():
    user_service, _, user_repo, role_repo, _, _ = init_test_default_user_service_components()
    
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
