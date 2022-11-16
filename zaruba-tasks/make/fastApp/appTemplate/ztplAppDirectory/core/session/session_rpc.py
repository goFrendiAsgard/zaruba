from typing import Any, Optional, Mapping
from transport import AppMessageBus, AppRPC
from core.session.session_service import SessionService
from core.security.service.auth_service import AuthService
from schema.user import User

import sys

def register_session_rpc(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, session_service: SessionService):

    @rpc.handle('create_access_token')
    def create_access_token(identity: str, password: str, current_user_data: Optional[Mapping[str, Any]] = None) -> str:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        return session_service.create_access_token(identity, password, current_user)


    @rpc.handle('renew_access_token')
    def renew_access_token(token: str, current_user_data: Optional[Mapping[str, Any]] = None) -> str:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        return session_service.renew_access_token(token, current_user)


    @rpc.handle('get_user_by_access_token')
    def get_user_by_access_token(token: str, current_user_data: Optional[Mapping[str, Any]] = None) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        user = session_service.get_user_by_token(token, current_user)
        return None if user is None else user.dict()


    print('Handle RPC for auth.Session', file=sys.stderr)