from typing import Any, List, Mapping
from auth.authModel import AuthModel
from auth.userModel import UserModel
from helpers.transport import MessageBus, RPC
from fastapi import Depends, FastAPI, HTTPException
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm
from schemas.user import User, UserData

import traceback

def register_user_route(app: FastAPI, mb: MessageBus, rpc: RPC, auth_model: AuthModel, user_model: UserModel):

    @app.post("/token")
    async def login(form_data: OAuth2PasswordRequestForm = Depends()):
        user = user_model.find_by_password(identity=form_data.username, password=form_data.password)
        if not user:
            raise HTTPException(status_code=400, detail="Incorrect username or password")
        access_token = user_model.create_token(user)
        return {"access_token": access_token, "token_type": "bearer"}

    @app.get('/users/', response_model=List[User])
    def find_user(current_user = Depends(auth_model.current_user_has_any_role(['user:*', 'user:read'])), keyword: str='', limit: int=100, offset: int=0) -> List[User]:
        results = []
        try:
            results = rpc.call('find_user', keyword, limit, offset)
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        return [User.parse_obj(result) for result in results]

    @app.get('/users/{id}', response_model=User)
    def find_user_by_id(id: str) -> User:
        result = None
        try:
            result = rpc.call('find_user_by_id', id)
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)

    @app.post('/users/', response_model=User)
    def insert_user(data: UserData) -> User:
        result = None
        try:
            result = rpc.call('insert_user', data.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)

    @app.put('/users/{id}', response_model=User)
    def update_user(id: str, data: UserData) -> User:
        result = None
        try:
            result = rpc.call('update_user', id, data.dict())
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)

    @app.delete('/users/{id}')
    def delete_user(id: str) -> User:
        result = None
        try:
            result = rpc.call('delete_user', id)
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')
        if result is None:
            raise HTTPException(status_code=404, detail='Not Found')
        return User.parse_obj(result)

    print('Handle HTTP routes for auth.User')