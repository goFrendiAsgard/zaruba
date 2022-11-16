from module.cms.content_type_seeder.test_content_type_seeder_service_util import ARTICLE_CONTENT_TYPE_DATA, init_test_content_type_seeder_service_components

def test_content_type_seeder_service_with_existing_user():
    content_type_seeder_service, content_type_service, content_type_repo, _, _ = init_test_content_type_seeder_service_components()
    # Init existing user
    content_type_repo.insert(ARTICLE_CONTENT_TYPE_DATA)
    # This should yield no error
    content_type_seeder_service.seed(ARTICLE_CONTENT_TYPE_DATA)
    assert content_type_service.find(keyword='', limit=10, offset=0).count == 1


def test_content_type_seeder_service_with_non_existing_user():
    content_type_seeder_service, content_type_service, _, _, _ = init_test_content_type_seeder_service_components()
    # This should yield no error
    content_type_seeder_service.seed(ARTICLE_CONTENT_TYPE_DATA)
    assert content_type_service.find(keyword='', limit=10, offset=0).count == 1