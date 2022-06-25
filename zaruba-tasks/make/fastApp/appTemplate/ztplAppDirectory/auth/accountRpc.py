from typing import Any, List, Mapping
from helpers.transport import RPC
from schemas.user import UserData
from schemas.user import User
from auth.userService import UserService
from auth.tokenService import TokenService

def register_account_rpc(rpc: RPC, user_service: UserService, token_service: TokenService):

    @rpc.handle('create_access_token')
    def create_access_token(identity: str, password: str) -> str:
        authenticated_user = user_service.find_by_password(identity, password)
        if not authenticated_user:
            raise Exception('Incorrect identity or password')
        return token_service.create_user_token(authenticated_user)

    @rpc.handle('refresh_access_token')
    def refresh_access_token(token: str) -> str:
        user = token_service.get_user_by_token(token)
        return token_service.create_user_token(user)

    print('Handle RPC for auth.Account')