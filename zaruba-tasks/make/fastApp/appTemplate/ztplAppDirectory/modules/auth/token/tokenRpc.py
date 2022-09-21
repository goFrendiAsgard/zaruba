from typing import Any, Optional, Mapping
from helpers.transport import RPC
from schemas.user import User
from modules.auth.token.tokenService import TokenService

def register_token_rpc(rpc: RPC, token_service: TokenService):

    @rpc.handle('create_user_token')
    def create_user_token(user_data: Mapping[str, Any]) -> str:
        user = User.parse_obj(user_data)
        return token_service.create_user_token(user)

    @rpc.handle('get_user_by_token')
    def get_user_by_token(token: str) -> Optional[Mapping[str, Any]]:
        user = token_service.get_user_by_token(token)
        return None if user is None else user.dict()

    print('Handle RPC for auth.Token')