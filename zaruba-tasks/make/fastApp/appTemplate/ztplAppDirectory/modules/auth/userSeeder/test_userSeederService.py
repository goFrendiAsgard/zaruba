from typing import Optional
from modules.auth.userSeeder.userSeederService import UserSeederService
from modules.auth.user.userService import UserService
from schemas.user import User, UserData, UserWithoutPassword, UserResult

################################################
# -- ⚙️ Mock data and objects
################################################

mock_new_user_data = UserData(
    username='mock_new_user',
    email='',
    phone_number='',
    permissions=[],
    role_ids=['mock_role_id'],
    active=True,
    full_name='',
    created_by='mock_user_id'
)

mock_existing_user_data = UserData(
    username='mock_existing_user',
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
        self.find_username: Optional[str] = None
        self.insert_user_data: Optional[UserData] = None
    
    def get_guest(self):
        return mock_guest_user
    
    def find(self, keyword: str, limit: int, offset: int) -> UserResult:
        return UserResult(count=1, rows=[mock_existing_user])

    def find_by_id(self, id: str) -> Optional[User]:
        return mock_existing_user

    def find_by_username(self, username: str) -> Optional[UserWithoutPassword]:
        self.find_username = username
        if username == mock_existing_user.username:
            return mock_existing_user
        return None

    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[User]:
        return mock_existing_user

    def insert(self, user_data: UserData) -> Optional[UserWithoutPassword]:
        self.insert_user_data = user_data
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

def test_user_seeder_service_seed_existing_user():
    mock_user_service = MockUserService()
    user_seeder_service = UserSeederService(mock_user_service)
    user_seeder_service.seed(mock_existing_user_data)
    # make sure all parameters are passed to user service
    assert mock_user_service.find_username == mock_existing_user_data.username
    assert mock_user_service.insert_user_data is None # not inserting new user, because user already exists
    

def test_user_seeder_service_seed_non_existing_user():
    mock_user_service = MockUserService()
    user_seeder_service = UserSeederService(mock_user_service)
    user_seeder_service.seed(mock_new_user_data)
    # make sure all parameters are passed to user service
    assert mock_user_service.find_username == mock_new_user_data.username
    assert mock_user_service.insert_user_data == mock_new_user_data
    
