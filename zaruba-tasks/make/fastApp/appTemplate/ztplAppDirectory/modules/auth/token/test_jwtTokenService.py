from modules.auth.user.test_util import create_user
from modules.auth.token.test_util import init_test_jwt_token_service_components, ROOT_USER_DATA

from sqlalchemy import create_engine


def test_jwt_token_service_with_empty_token():
    token_service, _, _, _, _, _, _ = init_test_jwt_token_service_components()
    user = token_service.get_user_by_token('')
    # make sure token service return correct value
    assert user is None
    

def test_jwt_token_service_with_null_token():
    token_service, _, _, _, _, _, _ = init_test_jwt_token_service_components()
    user = token_service.get_user_by_token(None)
    # make sure token service return correct value
    assert user is None


def test_jwt_token_service_with_invalid_token():
    token_service, _, _, _, _, _, _ = init_test_jwt_token_service_components()
    user = token_service.get_user_by_token('invalid token')
    # make sure token service return correct value
    assert user is None


def test_jwt_token_service_with_existing_user():
    token_service, _, _, _, user_repo, _, _ = init_test_jwt_token_service_components()
    # Init existing user
    root_user = user_repo.insert(ROOT_USER_DATA)
    # Creating token
    token = token_service.create_user_token(root_user)
    # make sure token service return correct value
    assert token is not None
    assert token != ''
    # Get user
    user = token_service.get_user_by_token(token)
    # make sure token service return correct value
    assert user == root_user


def test_jwt_token_service_with_non_existing_user():
    token_service, _, _, _, user_repo, _, _ = init_test_jwt_token_service_components()
    # Init existing user
    user_repo.insert(ROOT_USER_DATA)
    # Init non existing user
    inexist_user = create_user()
    inexist_user.id='inexist'
    inexist_user.username = 'inexist'
    inexist_user.email = 'inexist@innistrad.com'
    inexist_user.phone_number = '+6213456784'
    inexist_user.password = 'root'
    # Creating token
    token = token_service.create_user_token(inexist_user)
    # make sure token service return correct value
    assert token is not None
    assert token != ''
    # Get user
    user = token_service.get_user_by_token(token)
    # make sure token service return correct value
    assert user is None

