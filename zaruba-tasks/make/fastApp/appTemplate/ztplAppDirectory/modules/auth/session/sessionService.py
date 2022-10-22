from typing import Optional
from modules.auth.user.userService import UserService
from modules.auth.token.tokenService import TokenService
from schemas.user import User

class SessionService():

    def __init__(self, user_service: UserService, token_service: TokenService):
        self.user_service = user_service
        self.token_service = token_service

    def create_access_token(self, identity: str, password: str) -> str:
        authenticated_user = self.user_service.find_by_identity_and_password(identity, password)
        if not authenticated_user:
            raise Exception('Incorrect identity or password')
        return self.token_service.create_access_token(authenticated_user)

    def get_user_by_token(self, token: str) -> Optional[User]:
        return self.token_service.get_user_by_token(token)

    def refresh_access_token(self, token: str) -> str:
        user = self.token_service.get_user_by_token(token)
        return self.token_service.create_access_token(user)