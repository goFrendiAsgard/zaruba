from typing import Any, List, Mapping
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, HTTPException
from fastapi.security import OAuth2
from auth.authService import AuthService
from schemas.role import Role, RoleData

import traceback

def register_role_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_service: AuthService):

    @app.get('/api/v1/roles/', response_model=List[Role])
    def find_role(keyword: str='', limit: int=100, offset: int=0, current_user = Depends(auth_service.has_any_permissions( 'role:read'))) -> List[Role]:
        results = []
        try:
            results = rpc.call('find_role', keyword, limit, offset, current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return [Role.parse_obj(result) for result in results]

    @app.get('/api/v1/roles/{id}', response_model=Role)
    def find_role_by_id(id: str, current_user = Depends(auth_service.has_any_permissions( 'role:read'))) -> Role:
        result = None
        try:
            result = rpc.call('find_role_by_id', id, current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return Role.parse_obj(result)

    @app.post('/api/v1/roles/', response_model=Role)
    def insert_role(role_data: RoleData, current_user = Depends(auth_service.has_any_permissions( 'role:create'))) -> Role:
        result = None
        try:
            result = rpc.call('insert_role', role_data.dict(), current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return Role.parse_obj(result)

    @app.put('/api/v1/roles/{id}', response_model=Role)
    def update_role(id: str, role_data: RoleData, current_user = Depends(auth_service.has_any_permissions( 'role:update'))) -> Role:
        result = None
        try:
            result = rpc.call('update_role', id, role_data.dict(), current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return Role.parse_obj(result)

    @app.delete('/api/v1/roles/{id}')
    def delete_role(id: str, current_user = Depends(auth_service.has_any_permissions( 'role:delete'))) -> Role:
        result = None
        try:
            result = rpc.call('delete_role', id, current_user.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return Role.parse_obj(result)

    print('Handle HTTP routes for auth.Role')