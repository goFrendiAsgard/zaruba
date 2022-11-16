from typing import Mapping, List, Any
from core import AuthService
from transport import AppMessageBus, AppRPC

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_ztpl_app_module_name_event_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    print('Register ztplAppModuleName event handler', file=sys.stderr)
