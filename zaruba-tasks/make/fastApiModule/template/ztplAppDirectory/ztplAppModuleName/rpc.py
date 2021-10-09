from typing import Mapping, List, Any
from helpers.transport import RPC

import traceback

def register_rpc_handler(rpc: RPC):
    print('register ztplAppModuleName RPC handler')
