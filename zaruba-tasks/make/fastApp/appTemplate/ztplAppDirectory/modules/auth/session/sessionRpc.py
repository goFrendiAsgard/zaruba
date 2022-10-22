from typing import Any, Optional, Mapping
from helpers.transport import RPC, MessageBus
from modules.auth.session.sessionService import SessionService
from schemas.user import User

def register_session_rpc(mb: MessageBus, rpc: RPC, session_service: SessionService):

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


    print('Handle RPC for auth.Session')