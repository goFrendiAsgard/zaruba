from typing import Mapping, List, Any
from core import AuthService
from modules.log.activity import ActivityService, register_activity_rpc
from transport import AppMessageBus, AppRPC

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_rpc_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, activity_service: ActivityService):

    register_activity_rpc(mb, rpc, auth_service, activity_service)

    print('Register log RPC handler', file=sys.stderr)
