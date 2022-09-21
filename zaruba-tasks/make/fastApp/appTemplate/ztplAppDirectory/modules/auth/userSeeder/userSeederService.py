from modules.auth.user.userService import UserService
from schemas.user import UserData

class UserSeederService():

    def __init__(self, user_service: UserService):
        self.user_service = user_service


    def seed(self, user_data: UserData):
        existing_root_user = self.user_service.find_by_username(user_data.username)
        if existing_root_user:
            return
        self.user_service.insert(user_data)