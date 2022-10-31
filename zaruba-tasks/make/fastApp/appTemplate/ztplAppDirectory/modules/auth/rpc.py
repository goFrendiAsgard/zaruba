from modules.auth.role import RoleService, register_role_rpc
from modules.auth.user import UserService, register_user_rpc
from helpers.transport import RPC, MessageBus
from core.security.service.authService import AuthService

import sys

def register_auth_rpc_handler(mb: MessageBus, rpc: RPC, auth_service: AuthService, role_service: RoleService, user_service: UserService):
    register_role_rpc(mb, rpc, auth_service, role_service)
    register_user_rpc(mb, rpc, auth_service, user_service)

    print('Register auth RPC handler', file=sys.stderr)
