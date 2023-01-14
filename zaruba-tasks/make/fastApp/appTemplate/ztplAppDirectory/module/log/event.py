from typing import Mapping, List, Any
from core import AuthService
from transport import AppMessageBus, AppRPC
from module.log.activity import ActivityService, register_activity_event

import traceback
import sys

# Note: 🤖 Don't delete the following statement
def register_log_event_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, activity_service: ActivityService):

    register_activity_event(mb, rpc, auth_service, activity_service)

    print('Register log event handler', file=sys.stderr)
