from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schemas.user import User
from auth.authService import AuthService
from ui.menuService import MenuService
from helpers.transport import MessageBus, RPC

import traceback

def register_ztpl_app_module_name_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService, menu_service: MenuService, templates: Jinja2Templates, enable_ui: bool):
    # NOTE: follow [this](https://fastapi.tiangolo.com/tutorial/security/first-steps/#how-it-looks) guide for authorization

    ################################################
    # -- 👓 User Interface
    ################################################
    if enable_ui:
        @app.get('/ztpl-app-module-name', response_class=HTMLResponse)
        async def user_interface(request: Request, context = Depends(menu_service.validate('ztplAppModuleName'))):
            return templates.TemplateResponse(
                'default_page.html', 
                context={
                    'request': request, 
                    'context': context
                }, 
                status_code=200
            )

    print('Register ztplAppModuleName route handler')

