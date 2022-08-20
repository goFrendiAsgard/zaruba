from auth.accountRpc import register_account_rpc
from auth.roleRpc import register_role_rpc
from auth.tokenRpc import register_token_rpc
from auth.userRpc import register_user_rpc
from auth.accountService import AccountService
from auth.userService import UserService
from auth.roleService import RoleService
from auth.tokenService import TokenService
from typing import Mapping, List, Any
from helpers.transport import RPC

import traceback
import sys

def register_auth_rpc_handler(rpc: RPC, role_service: RoleService, user_service: UserService, token_service: TokenService, account_service: AccountService):
    register_account_rpc(rpc, account_service)
    register_token_rpc(rpc, token_service)
    register_role_rpc(rpc, role_service)
    register_user_rpc(rpc, user_service)

    print('Register auth RPC handler')
