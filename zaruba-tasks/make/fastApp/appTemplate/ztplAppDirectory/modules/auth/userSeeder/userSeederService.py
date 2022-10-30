from typing import Optional
from fastapi import HTTPException
from core import AuthService
from modules.auth.user.userService import UserService
from schemas.user import UserData, User

class UserSeederService():

    def __init__(self, auth_service: AuthService, user_service: UserService):
        self.auth_service = auth_service
        self.user_service = user_service


    def seed(self, user_data: UserData):
        system_user = self.auth_service.get_system_user()
        try:
            self.user_service.find_by_username(user_data.username, system_user)
        except HTTPException as error:
            if error.status_code == 404:
                self.user_service.insert(user_data, system_user)