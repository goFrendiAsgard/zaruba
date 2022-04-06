from auth.roleRpc import register_role_rpc
from auth.userRpc import register_user_rpc
from auth.userService import UserService
from auth.roleService import RoleService
from auth.tokenService import TokenService
from typing import Mapping, List, Any
from helpers.transport import RPC

import traceback

def register_auth_rpc_handler(rpc: RPC, role_service: RoleService, user_service: UserService, token_service: TokenService):
    register_role_rpc(rpc, role_service)
    register_user_rpc(rpc, user_service, token_service)

    print('Register auth RPC handler')
