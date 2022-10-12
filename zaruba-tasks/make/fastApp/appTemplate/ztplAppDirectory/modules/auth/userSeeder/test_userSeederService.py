from modules.auth.userSeeder.test_util import ROOT_USER_DATA, init_test_user_seeder_service_components


def test_user_seeder_service_with_existing_user():
    user_seeder_service, _, user_service, _, user_repo, _, _ = init_test_user_seeder_service_components()
    # Init existing user
    user_repo.insert(ROOT_USER_DATA)
    # This should yield no error
    user_seeder_service.seed(ROOT_USER_DATA)
    assert user_service.find(keyword='', limit=10, offset=0).count == 1


def test_user_seeder_service_with_non_existing_user():
    user_seeder_service, _, user_service, _, _, _, _ = init_test_user_seeder_service_components()
    # This should yield no error
    user_seeder_service.seed(ROOT_USER_DATA)
    assert user_service.find(keyword='', limit=10, offset=0).count == 1