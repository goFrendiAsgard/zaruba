from typing import Mapping, List, Any
from core import AuthService
from helpers.transport import RPC, MessageBus

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_ztpl_app_module_name_rpc_handler(mb: MessageBus, rpc: RPC, auth_service: AuthService):

    print('Register ztplAppModuleName RPC handler', file=sys.stderr)
