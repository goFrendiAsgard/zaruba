from modules.auth.role import RoleService, register_role_rpc
from modules.auth.user import UserService, register_user_rpc
from helpers.transport import RPC, MessageBus

import sys

def register_auth_rpc_handler(mb: MessageBus, rpc: RPC, role_service: RoleService, user_service: UserService):
    register_role_rpc(mb, rpc, role_service)
    register_user_rpc(mb, rpc, user_service)

    print('Register auth RPC handler', file=sys.stderr)
