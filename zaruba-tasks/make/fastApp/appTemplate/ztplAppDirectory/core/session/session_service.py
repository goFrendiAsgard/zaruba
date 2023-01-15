from typing import Optional
from module.auth.user.user_service import UserService
from core.token.token_service import TokenService
from schema.user import User


class SessionService():
    '''
    Service to handle user session:

    You can use SessionService to:
    - Create credential token for a valid user.
    - Renew credential token.
    - Get user based on credential token.
    '''

    def __init__(self, user_service: UserService, token_service: TokenService):
        self.user_service = user_service
        self.token_service = token_service

    def create_cred_token(
        self,
        identity: str,
        password: str,
        current_user: Optional[User] = None
    ) -> str:
        '''
        Return credential token if identity and password is valid
        '''
        authenticated_user = self.user_service.find_by_identity_and_password(
            identity, password, current_user
        )
        if not authenticated_user:
            raise Exception('Incorrect identity or password')
        return self.token_service.create_cred_token(
            authenticated_user, current_user
        )

    def get_user_by_cred_token(
        self,
        cred_token: str,
        current_user: Optional[User] = None
    ) -> Optional[User]:
        '''
        Return a user by credential token
        '''
        return self.token_service.get_user_by_cred_token(
            cred_token, current_user
        )

    def renew_cred_token(
        self,
        cred_token: str,
        current_user: Optional[User] = None
    ) -> str:
        '''
        Return a new credential token based on old credential token.
        '''
        user = self.token_service.get_user_by_cred_token(
            cred_token, current_user
        )
        return self.token_service.create_cred_token(user, current_user)
