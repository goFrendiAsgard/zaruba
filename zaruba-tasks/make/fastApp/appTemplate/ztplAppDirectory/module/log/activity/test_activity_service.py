from module.log.activity.test_activity_service_util import create_activity_data, insert_activity_data, init_test_activity_service_components
from module.auth.user.test_default_user_service_util import AUTHORIZED_ACTIVE_USER


def test_activity_service_crud_find_by_id_existing():
    activity_service, _, activity_repo, _, _ = init_test_activity_service_components()
    # prepare repo
    existing_activity = insert_activity_data(activity_repo)
    # test find by id (existing)
    fetched_activity = activity_service.find_by_id(existing_activity.id, AUTHORIZED_ACTIVE_USER)
    assert fetched_activity is not None
    assert fetched_activity.id == existing_activity.id
    assert fetched_activity.activity == 'activity'
    assert fetched_activity.created_by == 'original_user'
    assert fetched_activity.updated_by == 'original_user'


def test_activity_service_crud_find_by_id_non_existing():
    activity_service, _, activity_repo, _, _ = init_test_activity_service_components()
    # prepare repo
    insert_activity_data(activity_repo)
    # test find by id (non existing)
    is_error = False
    try:
        activity_service.find_by_id('invalid-id', AUTHORIZED_ACTIVE_USER)
    except Exception:
        is_error = True
    assert is_error


def test_activity_service_crud_find_existing():
    activity_service, _, activity_repo, _, _ = init_test_activity_service_components()
    # prepare repo
    existing_activity = insert_activity_data(activity_repo)
    # test find (existing)
    fetched_activity_result = activity_service.find(keyword='activity', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert fetched_activity_result.count == 1
    fetched_activity = fetched_activity_result.rows[0]
    assert fetched_activity is not None
    assert fetched_activity.id == existing_activity.id
    assert fetched_activity.activity == 'activity'
    assert fetched_activity.created_by == 'original_user'
    assert fetched_activity.updated_by == 'original_user'


def test_activity_service_crud_find_non_existing():
    activity_service, _, activity_repo, _, _ = init_test_activity_service_components()
    # prepare repo
    insert_activity_data(activity_repo)
    # test find (non existing)
    non_existing_activity_result = activity_service.find(keyword='invalid-keyword', limit=100, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert non_existing_activity_result.count == 0


def test_activity_service_crud_find_pagination():
    activity_service, _, activity_repo, _, _ = init_test_activity_service_components()
    # prepare repo
    for index in range(7):
        insert_activity_data(activity_repo, index)
    # test find (page 1)
    fetched_activity_result = activity_service.find(keyword='activity', limit=3, offset=0, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_activity_result.rows) == 3
    assert fetched_activity_result.count == 7
    # test find (page 2)
    fetched_activity_result = activity_service.find(keyword='activity', limit=3, offset=3, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_activity_result.rows) == 3
    assert fetched_activity_result.count == 7
    # test find (page 3)
    fetched_activity_result = activity_service.find(keyword='activity', limit=3, offset=6, current_user=AUTHORIZED_ACTIVE_USER)
    assert len(fetched_activity_result.rows) == 1
    assert fetched_activity_result.count == 7


def test_activity_service_crud_insert():
    activity_service, auth_service, activity_repo, _, _ = init_test_activity_service_components()
    system_user = auth_service.get_system_user()
    # prepare insert
    inserted_activity_data = create_activity_data()
    inserted_activity_data.activity = 'activity'
    inserted_activity_data.created_by = 'original_user'
    inserted_activity_data.updated_by = 'original_user'
    # test insert
    inserted_activity = activity_service.insert(inserted_activity_data)
    assert inserted_activity is not None
    assert inserted_activity.id != '' 
    assert inserted_activity.activity == 'activity'
    assert inserted_activity.created_by == system_user.id
    assert inserted_activity.updated_by == system_user.id
    assert activity_repo.count(keyword='') == 1
