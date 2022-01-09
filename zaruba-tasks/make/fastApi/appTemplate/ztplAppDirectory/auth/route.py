from auth.roleRoute import register_role_route
from auth.userRoute import register_user_route
from typing import Mapping, List, Any
from pydantic import BaseModel
from fastapi import Depends, FastAPI, HTTPException
from fastapi.security import OAuth2PasswordRequestForm
from auth.authModel import AuthModel
from auth.userModel import UserModel
from helpers.transport import MessageBus, RPC

import traceback

class TokenResponse(BaseModel):
    access_token: str
    token_type: str


def register_auth_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, access_token_url: str, auth_model: AuthModel):

    @app.post(access_token_url, response_model=TokenResponse)
    async def login(form_data: OAuth2PasswordRequestForm = Depends()):
        try:
            access_token = rpc.call('get_user_token', form_data.username, form_data.password)
            return TokenResponse(access_token = access_token, token_type = 'bearer')
        except:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=400, detail='Incorrect identity or password')

    register_role_route(app, mb, rpc, auth_model)

    register_user_route(app, mb, rpc, auth_model)

    print('Register auth route handler')