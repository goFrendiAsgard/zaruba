from modules.log.activity.repos.activityRepo import ActivityRepo
from modules.log.activity.activityRpc import register_activity_entity_rpc
from typing import Mapping, List, Any
from helpers.transport import RPC, MessageBus

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_rpc_handler(mb: MessageBus, rpc: RPC, activity_repo: ActivityRepo):

    register_activity_entity_rpc(mb, rpc, activity_repo)

    print('Register log RPC handler', file=sys.stderr)
