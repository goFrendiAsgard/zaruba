from helpers.transport import RPC
from modules.auth.session.sessionService import SessionService

def register_session_rpc(rpc: RPC, session_service: SessionService):

    @rpc.handle('create_access_token')
    def create_access_token(identity: str, password: str) -> str:
        return session_service.create_access_token(identity, password)

    @rpc.handle('refresh_access_token')
    def refresh_access_token(token: str) -> str:
        return session_service.refresh_access_token(token)

    print('Handle RPC for auth.Account')