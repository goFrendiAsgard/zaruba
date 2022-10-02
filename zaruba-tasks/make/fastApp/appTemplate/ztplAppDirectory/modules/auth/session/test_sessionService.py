from typing import Optional
from modules.auth.session.sessionService import SessionService
from modules.auth.user.userService import UserService
from modules.auth.token.tokenService import TokenService
from schemas.user import User, UserData, UserWithoutPassword, UserResult

################################################
# -- ⚙️ Mock data and objects
################################################

mock_authenticated_user = User(
    id="mock_existing_user_id",
    username='mock_authenticated_user',
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
    email='',
    phone_number='',
    permissions=[],
    role_ids=[],
    active=True,
    full_name='',
    created_by='mock_user_id'
)

mock_authenticated_token = 'mock_authenticated_token'
mock_authenticated_password = 'mock_authenticated_password'


class MockUserService(UserService):
    def __init__(self):
        self.create_access_token_identity: Optional[str] = None
        self.create_access_token_password: Optional[str] = None
    
    def get_guest(self):
        return mock_guest_user
    
    def find(self, keyword: str, limit: int, offset: int) -> UserResult:
        return UserResult(count=1, rows=[mock_guest_user])

    def find_by_id(self, id: str) -> Optional[User]:
        return mock_guest_user

    def find_by_username(self, username: str) -> Optional[UserWithoutPassword]:
        return mock_guest_user

    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[User]:
        self.create_access_token_identity = identity
        self.create_access_token_password = password
        if identity == mock_authenticated_user.username and password == mock_authenticated_password:
            return mock_authenticated_user
        return None

    def insert(self, user_data: UserData) -> Optional[UserWithoutPassword]:
        return mock_guest_user

    def update(self, id: str, user_data: UserData) -> Optional[User]:
        return mock_guest_user

    def delete(self, id: str) -> Optional[User]:
        return mock_guest_user

    def is_authorized(self, user: User, permission: str) -> bool:
        return True


class MockTokenService(TokenService):
    def __init__(self):
        self.create_user_token_user: Optional[User] = None
        self.get_user_by_token_token: Optional[str] = None

    def create_user_token(self, user: User) -> str:
        self.create_user_token_user = user
        if user.id == mock_authenticated_user.id:
            return mock_authenticated_token
        return 'random_token'

    def get_user_by_token(self, token: str) -> Optional[User]:
        self.get_user_by_token_token = token
        if token == mock_authenticated_token:
            return mock_authenticated_user
        return None


################################################
# -- 🧪 Test
################################################

def test_session_service_create_access_token_for_authenticated_user():
    mock_user_service = MockUserService()
    mock_token_service = MockTokenService()
    session_service = SessionService(mock_user_service, mock_token_service)
    token = session_service.create_access_token(mock_authenticated_user.username, mock_authenticated_password)
    # make sure all parameters are passed to user and token service
    assert mock_user_service.create_access_token_identity == mock_authenticated_user.username
    assert mock_user_service.create_access_token_password == mock_authenticated_password
    assert mock_token_service.create_user_token_user == mock_authenticated_user
    # make sure token service return correct value
    assert token == mock_authenticated_token


def test_session_service_create_access_token_for_unauthenticated_user():
    mock_user_service = MockUserService()
    mock_token_service = MockTokenService()
    session_service = SessionService(mock_user_service, mock_token_service)
    is_error_happened = False
    try:
        session_service.create_access_token('invalid_identity', 'invalid_password')
    except:
        is_error_happened = True
    # make sure all parameters are passed to user and token service
    assert mock_user_service.create_access_token_identity == 'invalid_identity'
    assert mock_user_service.create_access_token_password == 'invalid_password'
    # make sure token service throw error
    assert is_error_happened


def test_session_service_refresh_access_token():
    mock_user_service = MockUserService()
    mock_token_service = MockTokenService()
    session_service = SessionService(mock_user_service, mock_token_service)
    token = session_service.refresh_access_token(mock_authenticated_token)
    # make sure all parameters are passed to user service
    assert mock_token_service.get_user_by_token_token == mock_authenticated_token
    assert mock_token_service.create_user_token_user is not None
    # make sure token service return correct value
    assert token == mock_authenticated_token

