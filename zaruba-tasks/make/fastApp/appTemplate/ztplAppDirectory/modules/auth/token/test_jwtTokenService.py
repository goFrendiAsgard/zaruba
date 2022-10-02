from typing import Optional
from modules.auth.user.userService import UserService
from modules.auth.token.tokenService import JWTTokenService
from schemas.user import User, UserData, UserWithoutPassword, UserResult

################################################
# -- ⚙️ Mock data and objects
################################################

mock_non_existing_user = User(
    id="mock_non_existing_user_id",
    username='mock_non_existing_user',
    email='',
    phone_number='',
    permissions=[],
    role_ids=['mock_role_id'],
    active=True,
    full_name='',
    created_by='mock_user_id'
)

mock_existing_user = User(
    id="mock_existing_user_id",
    username='mock_existing_user',
    email='',
    phone_number='',
    permissions=[],
    role_ids=['mock_role_id'],
    active=True,
    full_name='',
    created_by='mock_user_id'
)

mock_guest_user = User(
    id="mock_guest_user_id",
    username="guest_username",
    active=True
)


class MockUserService(UserService):
    def __init__(self):
        self.find_id: Optional[str] = None
    
    def get_guest(self):
        return mock_guest_user
    
    def find(self, keyword: str, limit: int, offset: int) -> UserResult:
        return UserResult(count=1, rows=[mock_existing_user])

    def find_by_id(self, id: str) -> Optional[User]:
        self.find_id = id
        if id == mock_existing_user.id:
            return mock_existing_user
        return None

    def find_by_username(self, username: str) -> Optional[UserWithoutPassword]:
        return mock_existing_user

    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[User]:
        return mock_existing_user

    def insert(self, user_data: UserData) -> Optional[UserWithoutPassword]:
        return mock_existing_user

    def update(self, id: str, user_data: UserData) -> Optional[User]:
        return mock_existing_user

    def delete(self, id: str) -> Optional[User]:
        return mock_existing_user

    def is_authorized(self, user: User, permission: str) -> bool:
        return True


################################################
# -- 🧪 Test
################################################

def test_jwt_token_service_get_user_by_token_empty():
    mock_user_service = MockUserService()
    token_service = JWTTokenService(mock_user_service, 'secret', 'HS256', 1800)
    user = token_service.get_user_by_token('')
    # make sure token service return correct value
    assert user is None
    

def test_jwt_token_service_get_user_by_token_none():
    mock_user_service = MockUserService()
    token_service = JWTTokenService(mock_user_service, 'secret', 'HS256', 1800)
    user = token_service.get_user_by_token(None)
    # make sure token service return correct value
    assert user is None


def test_jwt_token_service_get_user_by_invalid_token():
    mock_user_service = MockUserService()
    token_service = JWTTokenService(mock_user_service, 'secret', 'HS256', 1800)
    user = token_service.get_user_by_token('invalid token')
    # make sure token service return correct value
    assert user is None


def test_jwt_token_service_get_user_by_token_with_existing_user():
    mock_user_service = MockUserService()
    token_service = JWTTokenService(mock_user_service, 'secret', 'HS256', 1800)
    # Creating token
    token = token_service.create_user_token(mock_existing_user)
    # make sure token service return correct value
    assert token is not None
    assert token != ''
    # Get user
    user = token_service.get_user_by_token(token)
    # make sure all parameters are passed to user service
    assert mock_user_service.find_id == mock_existing_user.id
    # make sure token service return correct value
    assert user == mock_existing_user


def test_jwt_token_service_get_user_by_token_with_non_existing_user():
    mock_user_service = MockUserService()
    token_service = JWTTokenService(mock_user_service, 'secret', 'HS256', 1800)
    # Creating token
    token = token_service.create_user_token(mock_non_existing_user)
    # make sure token service return correct value
    assert token is not None
    assert token != ''
    # Get user
    user = token_service.get_user_by_token(token)
    # make sure all parameters are passed to user service
    assert mock_user_service.find_id == mock_non_existing_user.id
    # make sure token service return correct value
    assert user is None

