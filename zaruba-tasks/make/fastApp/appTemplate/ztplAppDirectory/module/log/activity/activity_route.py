from typing import Optional
from helper.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from core import AuthService, MenuService
from schema.activity import Activity, ActivityData, ActivityResult
from schema.menu_context import MenuContext
from schema.user import User
from schema.auth_type import AuthType

import logging


################################################
# -- ⚙️ API
################################################
def register_activity_api_route(
    app: FastAPI,
    mb: MessageBus,
    rpc: RPC, auth_service: AuthService
):

    @app.get(
        '/api/v1/activities',
        response_model=ActivityResult
    )
    async def find_activities(
        keyword: str = '',
        limit: int = 100,
        offset: int = 0,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:activity:read')
        )
    ) -> ActivityResult:
        '''
        Serving API to find activities by keyword.
        '''
        result = {}
        try:
            current_user = _get_user_or_guest(current_user)
            result = rpc.call(
                'find_activity', keyword, limit, offset, current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return ActivityResult.parse_obj(result)

    @app.get(
        '/api/v1/activities/{id}',
        response_model=Activity
    )
    async def find_activity_by_id(
        id: str,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:activity:read')
        )
    ) -> Activity:
        '''
        Serving API to find activity by id.
        '''
        activity = None
        try:
            current_user = _get_user_or_guest(current_user)
            activity = rpc.call('find_activity_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return Activity.parse_obj(activity)

    @app.post(
        '/api/v1/activities',
        response_model=Activity
    )
    async def insert_activity(
        activity_data: ActivityData,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:activity:create')
        )
    ) -> Activity:
        '''
        Serving API to insert new activity.
        '''
        activity = None
        try:
            current_user = _get_user_or_guest(current_user)
            activity = rpc.call(
                'insert_activity', activity_data.dict(), current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return Activity.parse_obj(activity)

    def _handle_non_http_exception():
        '''
        Handle non HTTPException and return a default HTTPException
        '''
        logging.error('Non HTTPException error', exc_info=True)
        raise HTTPException(
            status_code=500,
            detail='Internal server serror'
        )

    def _get_user_or_guest(user: Optional[User]) -> User:
        '''
        If user is not set, this function will return guest_user
        '''
        if user is None:
            return User.parse_obj(auth_service.get_guest_user())
        return user

    logging.info('Register log.activity API route handler')


################################################
# -- 👓 User Interface
################################################
def register_activity_ui_route(
    app: FastAPI, 
    mb: MessageBus, 
    rpc: RPC, 
    menu_service: MenuService, 
    page_template: Jinja2Templates
):

    # Activity list page
    menu_service.add_menu(
        name='log:activities',
        title='User Activities',
        url='/log/activities',
        auth_type=AuthType.HAS_PERMISSION,
        permission_name='ui:log:activity', parent_name='log'
    )

    @app.get(
        '/log/activities',
        response_class=HTMLResponse
    )
    async def manage_activity(
        request: Request,
        context: MenuContext = Depends(
            menu_service.has_access('log:activities')
        )
    ):
        '''
        Serving user interface for managing activity.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/log/crud/activities.html',
            'request': request,
            'context': context
        }, status_code=200)
