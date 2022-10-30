from typing import Mapping, List, Any
from helpers.transport import RPC, MessageBus

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_log_event_handler(mb: MessageBus, rpc: RPC):

    print('Register log event handler', file=sys.stderr)
