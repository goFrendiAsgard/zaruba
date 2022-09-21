from modules.auth.role.roleRpc import register_role_rpc
from modules.auth.role.roleService import RoleService
from modules.auth.session.sessionRpc import register_session_rpc
from modules.auth.session.sessionService import SessionService
from modules.auth.token.tokenRpc import register_token_rpc
from modules.auth.token.tokenService import TokenService
from modules.auth.user.userRpc import register_user_rpc
from modules.auth.user.userService import UserService
from helpers.transport import RPC

def register_auth_rpc_handler(rpc: RPC, role_service: RoleService, user_service: UserService, token_service: TokenService, session_service: SessionService):
    register_session_rpc(rpc, session_service)
    register_token_rpc(rpc, token_service)
    register_role_rpc(rpc, role_service)
    register_user_rpc(rpc, user_service)

    print('Register auth RPC handler')
