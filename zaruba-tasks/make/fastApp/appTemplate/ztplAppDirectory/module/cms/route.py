from module.cms.content import register_content_api_route, register_content_ui_route
from module.cms.content_type import register_content_type_api_route, register_content_type_ui_route
from typing import Mapping, List, Any, Optional
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from schema.menu_context import MenuContext
from schema.user import User
from schema.auth_type import AuthType

import traceback
import sys


################################################
# -- ⚙️ API
################################################
# Note: 🤖 Don't delete the following statement
def register_cms_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    register_content_api_route(app, mb, rpc, auth_service)
    register_content_type_api_route(app, mb, rpc, auth_service)

    print('Register cms api route handler', file=sys.stderr)


################################################
# -- 👓 User Interface
################################################
# Note: 🤖 Don't delete the following statement
def register_cms_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):
    
    # CMS menu
    menu_service.add_menu(name='cms', title='CMS', url='#', auth_type=AuthType.ANYONE)

    register_content_ui_route(app, mb, rpc, menu_service, page_template)
    register_content_type_ui_route(app, mb, rpc, menu_service, page_template)

    print('Register cms UI route handler', file=sys.stderr)