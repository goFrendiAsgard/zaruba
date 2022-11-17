from typing import Mapping, List, Any
from core import AuthService
from transport import AppMessageBus, AppRPC

import traceback
import sys

# Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
def register_ztpl_app_module_name_rpc_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    print('Register ztplAppModuleName RPC handler', file=sys.stderr)
