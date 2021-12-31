from repos.user import UserRepo
from auth.userRpc import register_user_rpc
from typing import Mapping, List, Any
from helpers.transport import RPC

import traceback

def register_auth_rpc_handler(rpc: RPC, user_repo: UserRepo):

    register_user_rpc(rpc, user_repo)

    print('Register auth RPC handler')
