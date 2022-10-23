from typing import Mapping, List, Any
from helpers.transport import MessageBus, RPC

import traceback
import sys

def register_auth_event_handler(mb: MessageBus, rpc: RPC):

    print('Register auth event handler', file=sys.stderr)
