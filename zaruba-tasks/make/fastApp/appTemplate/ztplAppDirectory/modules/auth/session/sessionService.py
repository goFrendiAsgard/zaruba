from modules.auth.user.userService import UserService
from modules.auth.token.tokenService import TokenService

class SessionService():

    def __init__(self, user_service: UserService, token_service: TokenService):
        self.user_service = user_service
        self.token_service = token_service

    def create_access_token(self, identity: str, password: str) -> str:
        authenticated_user = self.user_service.find_by_identity_and_password(identity, password)
        if not authenticated_user:
            raise Exception('Incorrect identity or password')
        return self.token_service.create_user_token(authenticated_user)

    def refresh_access_token(self, token: str) -> str:
        user = self.token_service.get_user_by_token(token)
        return self.token_service.create_user_token(user)