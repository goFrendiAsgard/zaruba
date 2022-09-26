from typing import Mapping, List, Any
from helpers.transport import RPC, MessageBus

import traceback
import sys

def register_ztpl_app_module_name_event_handler(mb: MessageBus, rpc: RPC):

    print('Register ztplAppModuleName event handler')
