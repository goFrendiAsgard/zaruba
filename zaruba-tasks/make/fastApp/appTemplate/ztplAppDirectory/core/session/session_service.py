from typing import Optional
from module.auth.user.user_service import UserService
from core.token.token_service import TokenService
from schema.user import User

class SessionService():

    def __init__(self, user_service: UserService, token_service: TokenService):
        self.user_service = user_service
        self.token_service = token_service

    def create_cred_token(self, identity: str, password: str, current_user: Optional[User] = None) -> str:
        authenticated_user = self.user_service.find_by_identity_and_password(identity, password, current_user)
        if not authenticated_user:
            raise Exception('Incorrect identity or password')
        return self.token_service.create_cred_token(authenticated_user, current_user)

    def get_user_by_token(self, token: str, current_user: Optional[User] = None) -> Optional[User]:
        return self.token_service.get_user_by_cred_token(token, current_user)

    def renew_cred_token(self, token: str, current_user: Optional[User] = None) -> str:
        user = self.token_service.get_user_by_cred_token(token, current_user)
        return self.token_service.create_cred_token(user, current_user)