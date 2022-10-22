from typing import Any, Optional, Mapping
from helpers.transport import RPC, MessageBus
from modules.auth.session.sessionService import SessionService

def register_session_rpc(mb: MessageBus, rpc: RPC, session_service: SessionService):

    @rpc.handle('create_access_token')
    def create_access_token(identity: str, password: str) -> str:
        return session_service.create_access_token(identity, password)

    @rpc.handle('refresh_access_token')
    def refresh_access_token(token: str) -> str:
        return session_service.refresh_access_token(token)

    @rpc.handle('get_user_by_access_token')
    def get_user_by_access_token(token: str) -> Optional[Mapping[str, Any]]:
        user = session_service.get_user_by_token(token)
        return None if user is None else user.dict()

    print('Handle RPC for auth.Account')