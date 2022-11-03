from typing import Mapping, List, Any
from transport import AppMessageBus, AppRPC
from core.security.service.authService import AuthService

import traceback
import sys

def register_auth_event_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    print('Register auth event handler', file=sys.stderr)
