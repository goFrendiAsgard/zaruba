from typing import Mapping, List, Any
from helpers.transport import MessageBus

import traceback

def register_auth_event_handler(mb: MessageBus):

    print('Register auth event handler')
