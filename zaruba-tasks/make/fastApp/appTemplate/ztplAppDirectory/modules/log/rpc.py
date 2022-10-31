from modules.log.activity import ActivityService, register_activity_rpc
from typing import Mapping, List, Any
from helpers.transport import RPC, MessageBus

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_rpc_handler(mb: MessageBus, rpc: RPC, activity_service: ActivityService):

    register_activity_rpc(mb, rpc, activity_service)

    print('Register log RPC handler', file=sys.stderr)
