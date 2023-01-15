from typing import Optional
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schema.content import Content, ContentData, ContentResult
from schema.menu_context import MenuContext
from schema.user import User
from schema.auth_type import AuthType

import logging


################################################
# -- ⚙️ API
################################################
def register_content_api_route(
    app: FastAPI,
    mb: AppMessageBus,
    rpc: AppRPC, auth_service: AuthService
):

    @app.get(
        '/api/v1/contents/',
        response_model=ContentResult
    )
    async def find_contents(
        keyword: str = '',
        limit: int = 100,
        offset: int = 0,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:content:read')
        )
    ) -> ContentResult:
        '''
        Serving API to find contents by keyword.
        '''
        result = {}
        try:
            current_user = _get_user_or_guest(current_user)
            result = rpc.call(
                'find_content', keyword, limit, offset, current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return ContentResult.parse_obj(result)

    @app.get(
        '/api/v1/contents/{id}',
        response_model=Content
    )
    async def find_content_by_id(
        id: str,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:content:read')
        )
    ) -> Content:
        '''
        Serving API to find content by id.
        '''
        content = None
        try:
            current_user = _get_user_or_guest(current_user)
            content = rpc.call('find_content_by_id', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return Content.parse_obj(content)

    @app.post(
        '/api/v1/contents/',
        response_model=Content
    )
    async def insert_content(
        content_data: ContentData,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:content:create')
        )
    ) -> Content:
        '''
        Serving API to insert new content.
        '''
        content = None
        try:
            current_user = _get_user_or_guest(current_user)
            content = rpc.call(
                'insert_content', content_data.dict(), current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return Content.parse_obj(content)

    @app.put(
        '/api/v1/contents/{id}',
        response_model=Content
    )
    async def update_content(
        id: str,
        content_data: ContentData,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:content:update')
        )
    ) -> Content:
        '''
        Serving API to update content by id.
        '''
        content = None
        try:
            current_user = _get_user_or_guest(current_user)
            content = rpc.call(
                'update_content', id, content_data.dict(), current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return Content.parse_obj(content)

    @app.delete(
        '/api/v1/contents/{id}',
        response_model=Content
    )
    async def delete_content(
        id: str,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:content:delete')
        )
    ) -> Content:
        '''
        Serving API to delete content by id.
        '''
        content = None
        try:
            current_user = _get_user_or_guest(current_user)
            content = rpc.call('delete_content', id, current_user.dict())
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return Content.parse_obj(content)

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

    logging.info('Register cms.content API route handler')


################################################
# -- 👓 User Interface
################################################
def register_content_ui_route(
    app: FastAPI,
    mb: AppMessageBus,
    rpc: AppRPC,
    menu_service: MenuService,
    page_template: Jinja2Templates
):

    # Content CRUD page
    menu_service.add_menu(
        name='cms:contents',
        title='Contents',
        url='/cms/contents',
        auth_type=AuthType.HAS_PERMISSION,
        permission_name='ui:cms:content',
        parent_name='cms'
    )

    @app.get(
        '/cms/contents',
        response_class=HTMLResponse
    )
    async def manage_content(
        request: Request,
        context: MenuContext = Depends(
            menu_service.has_access('cms:contents')
        )
    ):
        '''
        Serving user interface for managing content.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path': 'modules/cms/crud/contents.html',
            'request': request,
            'context': context
        }, status_code=200)

    logging.info('Register cms.content UI route handler')
