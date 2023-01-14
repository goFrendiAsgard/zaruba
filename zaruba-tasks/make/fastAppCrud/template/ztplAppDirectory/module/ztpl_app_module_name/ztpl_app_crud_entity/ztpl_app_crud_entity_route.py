from typing import Any, List, Mapping, Optional
from core import AuthService, MenuService
from transport import AppMessageBus, AppRPC
from fastapi import Depends, FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from schema.ztpl_app_crud_entity import (
    ZtplAppCrudEntity, ZtplAppCrudEntityData, ZtplAppCrudEntityResult
)
from schema.menu_context import MenuContext
from schema.user import User
from schema.auth_type import AuthType

import logging


################################################
# -- ⚙️ API
################################################
def register_ztpl_app_crud_entity_api_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, auth_service: AuthService):

    @app.get(
        '/api/v1/ztpl-app-crud-entities',
        response_model=ZtplAppCrudEntityResult
    )
    async def find_ztpl_app_crud_entities(
        keyword: str = '',
        limit: int = 100,
        offset: int = 0,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:ztpl_app_crud_entity:read')
        )
    ) -> ZtplAppCrudEntityResult:
        '''
        Serving API to find ztplAppCrudEntities by keyword.
        '''
        result = {}
        try:
            current_user = _get_user_or_guest(current_user)
            result = rpc.call(
                'find_ztpl_app_crud_entity',
                keyword, limit, offset, current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return ZtplAppCrudEntityResult.parse_obj(result)

    @app.get(
        '/api/v1/ztpl-app-crud-entities/{id}',
        response_model=ZtplAppCrudEntity
    )
    async def find_ztpl_app_crud_entity_by_id(
        id: str,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:ztpl_app_crud_entity:read')
        )
    ) -> ZtplAppCrudEntity:
        '''
        Serving API to find ztplAppCrudEntity by id.
        '''
        ztpl_app_crud_entity = None
        try:
            current_user = _get_user_or_guest(current_user)
            ztpl_app_crud_entity = rpc.call(
                'find_ztpl_app_crud_entity_by_id',
                id, current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)

    @app.post(
        '/api/v1/ztpl-app-crud-entities',
        response_model=ZtplAppCrudEntity
    )
    async def insert_ztpl_app_crud_entity(
        ztpl_app_crud_entity_data: ZtplAppCrudEntityData,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:ztpl_app_crud_entity:create')
        )
    ) -> ZtplAppCrudEntity:
        '''
        Serving API to insert new ztplAppCrudEntity.
        '''
        ztpl_app_crud_entity = None
        try:
            current_user = _get_user_or_guest(current_user)
            ztpl_app_crud_entity = rpc.call(
                'insert_ztpl_app_crud_entity',
                ztpl_app_crud_entity_data.dict(), current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)

    @app.put(
        '/api/v1/ztpl-app-crud-entities/{id}',
        response_model=ZtplAppCrudEntity
    )
    async def update_ztpl_app_crud_entity(
        id: str,
        ztpl_app_crud_entity_data: ZtplAppCrudEntityData,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:ztpl_app_crud_entity:update')
        )
    ) -> ZtplAppCrudEntity:
        '''
        Serving API to update ztplAppCrudEntity by id.
        '''
        ztpl_app_crud_entity = None
        try:
            current_user = _get_user_or_guest(current_user)
            ztpl_app_crud_entity = rpc.call(
                'update_ztpl_app_crud_entity',
                id, ztpl_app_crud_entity_data.dict(), current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)

    @app.delete(
        '/api/v1/ztpl-app-crud-entities/{id}',
        response_model=ZtplAppCrudEntity
    )
    async def delete_ztpl_app_crud_entity(
        id: str,
        current_user: Optional[User] = Depends(
            auth_service.has_permission('api:ztpl_app_crud_entity:delete')
        )
    ) -> ZtplAppCrudEntity:
        '''
        Serving API to delete ztplAppCrudEntity by id.
        '''
        ztpl_app_crud_entity = None
        try:
            current_user = _get_user_or_guest(current_user)
            ztpl_app_crud_entity = rpc.call(
                'delete_ztpl_app_crud_entity',
                id, current_user.dict()
            )
        except HTTPException as http_exception:
            raise http_exception
        except Exception:
            _handle_non_http_exception()
        return ZtplAppCrudEntity.parse_obj(ztpl_app_crud_entity)

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

    logging.info(
        'Register ztplAppModuleName.ztpl_app_crud_entity API route handler'
    )


################################################
# -- 👓 User Interface
################################################
def register_ztpl_app_crud_entity_ui_route(app: FastAPI, mb: AppMessageBus, rpc: AppRPC, menu_service: MenuService, page_template: Jinja2Templates):

    # ZtplAppCrudEntity CRUD page
    menu_service.add_menu(
        name='ztplAppModuleName:ztplAppCrudEntities',
        title='ZtplAppCrudEntities',
        url='/ztpl-app-module-name/ztpl-app-crud-entities',
        auth_type=AuthType.HAS_PERMISSION,
        permission_name='ui:ztplAppModuleName:ztplAppCrudEntity',
        parent_name='ztplAppModuleName'
    )

    @app.get(
        '/ztpl-app-module-name/ztpl-app-crud-entities',
        response_class=HTMLResponse
    )
    async def manage_ztpl_app_crud_entity(
        request: Request,
        context: MenuContext = Depends(
            menu_service.has_access('ztplAppModuleName:ztplAppCrudEntities')
        )
    ):
        '''
        Serving user interface for managing ztplAppCrudEntity.
        '''
        return page_template.TemplateResponse('default_crud.html', context={
            'content_path':
                'modules/ztplAppModuleName/crud/ztpl_app_crud_entities.html',
            'request': request,
            'context': context
        }, status_code=200)

    logging.info(
        'Register ztplAppModuleName.ztpl_app_crud_entity UI route handler'
    )
