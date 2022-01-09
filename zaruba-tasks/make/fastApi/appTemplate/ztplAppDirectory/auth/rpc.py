from auth.roleRpc import register_role_rpc
from auth.userRpc import register_user_rpc
from auth.userModel import UserModel
from auth.roleModel import RoleModel
from auth.tokenModel import TokenModel
from typing import Mapping, List, Any
from helpers.transport import RPC

import traceback

def register_auth_rpc_handler(rpc: RPC, role_model: RoleModel, user_model: UserModel, token_model: TokenModel):
    register_role_rpc(rpc, role_model)
    register_user_rpc(rpc, user_model, token_model)

    print('Register auth RPC handler')
