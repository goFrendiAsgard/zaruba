from typing import Mapping, List, Any
from helpers.transport import MessageBus

import traceback

def register_ztpl_app_module_name_event_handler(mb: MessageBus):

    print('Register ztplAppModuleName event handler')
