from modules.cms.contentAttribute import ContentAttributeService, register_content_attribute_rpc
from modules.cms.content import ContentService, register_content_rpc
from modules.cms.contentType import ContentTypeService, register_content_type_rpc
from typing import Mapping, List, Any
from core import AuthService
from transport import AppMessageBus, AppRPC

import traceback
import sys

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
def register_cms_rpc_handler(mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService, content_type_service: ContentTypeService, content_service: ContentService, content_attribute_service: ContentAttributeService):

    register_content_attribute_rpc(mb, rpc, auth_service, content_attribute_service)

    register_content_rpc(mb, rpc, auth_service, content_service)

    register_content_type_rpc(mb, rpc, auth_service, content_type_service)

    print('Register cms RPC handler', file=sys.stderr)
