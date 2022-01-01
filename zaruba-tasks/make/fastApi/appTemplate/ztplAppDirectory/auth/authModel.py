from typing import Callable, List
from fastapi.security import OAuth2PasswordBearer, OAuth2
from fastapi import Depends, FastAPI, HTTPException, status
from starlette.requests import Request
from auth.userModel import UserModel
from schemas.user import User

class AuthModel():

    def __init__(self, user_model: UserModel, oauth2_scheme: OAuth2, root_permission: str):
        self.user_model = user_model
        self.oauth2_scheme = oauth2_scheme
        self.root_permission = root_permission

    def raise_unauthorized_exception(self, detail: str):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )

    def everyone(self) -> Callable[[Request], User]:
        async def verify_everyone(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                return self.user_model.get_guest_user()
            current_user = self.user_model.find_by_token(token)
            if not current_user:
                return self.user_model.get_guest_user()
            return current_user
        return verify_everyone 

    def is_authenticated(self) -> Callable[[Request], User]:
        async def verify_is_authenticated(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                self.raise_unauthorized_exception('Not authenticated')
            current_user = self.user_model.find_by_token(token)
            if not current_user:
                self.raise_unauthorized_exception('Not authenticated')
            return current_user
        return verify_is_authenticated

    def has_any_permissions(self, *permissions: str) -> Callable[[Request], User]:
        async def verify_has_any_permission(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                self.raise_unauthorized_exception('Not authenticated')
            current_user = self.user_model.find_by_token(token)
            if not current_user:
                self.raise_unauthorized_exception('Not authenticated')
            if len(permissions) == 0:
                return current_user
            if current_user.has_permission(self.root_permission):
                return current_user
            for permission in permissions:
                if current_user.has_permission(permission):
                    return current_user
            self.raise_unauthorized_exception('Unauthorized')
        return verify_has_any_permission