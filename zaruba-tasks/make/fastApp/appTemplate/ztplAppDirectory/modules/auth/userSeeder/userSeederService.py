from typing import Optional
from fastapi import HTTPException
from modules.auth.user.userService import UserService
from schemas.user import UserData, User

class UserSeederService():

    def __init__(self, user_service: UserService):
        self.user_service = user_service


    def seed(self, user_data: UserData, current_user: Optional[User] = None):
        try:
            self.user_service.find_by_username(user_data.username, current_user)
        except HTTPException as error:
            if error.status_code == 404:
                if current_user is None:
                    current_user = self.user_service.get_system_user()
                self.user_service.insert(user_data, current_user)