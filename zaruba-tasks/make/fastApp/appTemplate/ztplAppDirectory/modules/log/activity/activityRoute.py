from typing import Any, List, Mapping, Optional
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.security import OAuth2
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from schemas.activity import Activity, ActivityData, ActivityResult
from schemas.menuContext import MenuContext
from schemas.user import User

import traceback
import sys

################################################
# -- ⚙️ API
################################################
def register_activity_api_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService):

    @app.get('/api/v1/activities/', response_model=ActivityResult)
    async def find_activities(keyword: str='', limit: int=100, offset: int=0, current_user: Optional[User] = Depends(auth_service.has_permission('api:activity:read'))) -> ActivityResult:
        '''
        Serving API to find activities by keyword.
        '''
        result = {}
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            result = rpc.call('find_activity', keyword, limit, offset, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return ActivityResult.parse_obj(result)


    @app.get('/api/v1/activities/{id}', response_model=Activity)
    async def find_activity_by_id(id: str, current_user: Optional[User] = Depends(auth_service.has_permission('api:activity:read'))) -> Activity:
        '''
        Serving API to find activity by id.
        '''
        activity = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            activity = rpc.call('find_activity_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return Activity.parse_obj(activity)


    @app.post('/api/v1/activities/', response_model=Activity)
    async def insert_activity(activity_data: ActivityData, current_user: Optional[User] = Depends(auth_service.has_permission('api:activity:create'))) -> Activity:
        '''
        Serving API to insert new activity.
        '''
        activity = None
        try:
            if not current_user:
                current_user = User.parse_obj(auth_service.get_guest_user())
            activity = rpc.call('insert_activity', activity_data.dict(), current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except:
            print(traceback.format_exc(), file=sys.stderr) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return Activity.parse_obj(activity)


################################################
# -- 👓 User Interface
################################################
def register_activity_ui_route(app: FastAPI, mb: MessageBus, rpc: RPC, menu_service: MenuService, page_template: Jinja2Templates):

    @app.get('/log/activities', response_class=HTMLResponse)
    async def manage_activity(request: Request, context: MenuContext = Depends(menu_service.has_access('log:activities'))):
        '''
        Serving user interface for managing activity.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'api_path': '/api/vi/ztp_app_crud_entities',
            'content_path': 'modules/log/crud/activities.html',
            'request': request, 
            'context': context
        }, status_code=200)
