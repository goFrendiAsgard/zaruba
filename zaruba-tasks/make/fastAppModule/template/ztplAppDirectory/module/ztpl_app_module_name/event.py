from typing import Mapping, List, Any
from core import AuthService
from transport import AppMessageBus, AppRPC

import logging

# Note: 🤖 Don't delete the following statement
def register_ztpl_app_module_name_event_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    logging.info('Register ztplAppModuleName event handler')
