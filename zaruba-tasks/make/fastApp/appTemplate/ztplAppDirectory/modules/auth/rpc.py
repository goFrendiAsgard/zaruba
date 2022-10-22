from modules.auth.role.roleRpc import register_role_rpc
from modules.auth.role.roleService import RoleService
from modules.auth.session.sessionRpc import register_session_rpc
from modules.auth.session.sessionService import SessionService
from modules.auth.user.userRpc import register_user_rpc
from modules.auth.user.userService import UserService
from helpers.transport import RPC, MessageBus

def register_auth_rpc_handler(mb: MessageBus, rpc: RPC, role_service: RoleService, user_service: UserService, session_service: SessionService):
    register_session_rpc(mb, rpc, session_service)
    register_role_rpc(mb, rpc, role_service)
    register_user_rpc(mb, rpc, user_service)

    print('Register auth RPC handler')
