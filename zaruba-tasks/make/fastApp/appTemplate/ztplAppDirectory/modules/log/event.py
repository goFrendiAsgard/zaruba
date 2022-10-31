from typing import Mapping, List, Any
from core import AuthService
from helpers.transport import RPC, MessageBus
from modules.log.activity import ActivityService, register_activity_event

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_event_handler(mb: MessageBus, rpc: RPC, auth_service: AuthService, activity_service: ActivityService):

    register_activity_event(mb, rpc, auth_service, activity_service)

    print('Register log event handler', file=sys.stderr)
