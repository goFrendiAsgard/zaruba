from typing import List
from schemas.user import User, UserData
from auth.tokenModel import TokenModel
from repos.user import UserRepo

import datetime

class UserModel():

    def __init__(self, user_repo: UserRepo, tokenModel: TokenModel, guest_username: str):
        self.user_repo = user_repo
        self.token_model = tokenModel
        self.guest_username = guest_username
        self.earliest_date = datetime.datetime.min

    def get_guest_user(self) -> User:
        return User(
            id = 'guest',
            username = self.guest_username, 
            email = '',
            roles = '',
            active = True,
            password = '',
            full_name = '',
            updated_at = self.earliest_date,
            created_at = self.earliest_date,
        )

    def create_token(self, user:User) -> str:
        return self.token_model.create_token(user.username)

    def find_by_token(self, token: str) -> User:
        username = self.token_model.extract_from_token(token)
        if username is None:
            return None
        return self.find_by_username(username)

    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        return self.user_repo.find(keyword, limit, offset)

    def find_by_id(self, id: str) -> User:
        return self.user_repo.find_by_id(id)

    def find_by_username(self, username: str) -> User:
        return self.user_repo.find_by_username(username)

    def find_by_password(self, identity: str, password: str) -> User:
        return self.user_repo.find_by_password(identity, password)

    def insert(self, user_data: UserData) -> User:
        return self.user_repo.insert(user_data)

    def update(self, id: str, user_data: UserData) -> User:
        return self.user_repo.update(id, user_data)

    def delete(self, id: str) -> User:
        return self.user_repo.delete(id)