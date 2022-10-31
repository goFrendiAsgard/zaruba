from typing import Mapping, List, Any
from helpers.transport import MessageBus, RPC
from core.security.service.authService import AuthService

import traceback
import sys

def register_auth_event_handler(mb: MessageBus, rpc: RPC, auth_service: AuthService):

    print('Register auth event handler', file=sys.stderr)
