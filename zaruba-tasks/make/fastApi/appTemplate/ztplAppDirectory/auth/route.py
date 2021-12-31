from auth.userRoute import register_user_route
from typing import Mapping, List, Any
from fastapi import Depends, FastAPI, HTTPException
from auth.authModel import AuthModel
from auth.userModel import UserModel
from helpers.transport import MessageBus, RPC

import traceback

def register_auth_route_handler(app: FastAPI, mb: MessageBus, rpc: RPC, auth_model: AuthModel, user_model: UserModel):

    register_user_route(app, mb, rpc, auth_model, user_model)

    print('Register auth route handler')