from typing import Callable, List
from fastapi.security import OAuth2PasswordBearer, OAuth2
from fastapi import Depends, FastAPI, HTTPException, status
from starlette.requests import Request
from auth.roleModel import RoleModel
from auth.userModel import UserModel
from auth.tokenModel import TokenModel
from schemas.user import User
import abc

class AuthModel(abc.ABC):

    @abc.abstractmethod
    def everyone(self) -> Callable[[Request], User]:
        pass

    @abc.abstractmethod
    def is_authenticated(self) -> Callable[[Request], User]:
        pass

    @abc.abstractmethod
    def has_any_permissions(self, *permissions: str) -> Callable[[Request], User]:
        pass

class NoAuthModel(AuthModel):

    def __init__(self, user_model: UserModel):
        self.user_model = user_model

    def always_authorized(self, Request) -> User:
        return self.user_model.get_guest_user()

    def everyone(self) -> Callable[[Request], User]:
        return self.always_authorized

    def is_authenticated(self) -> Callable[[Request], User]:
        return self.always_authorized

    def has_any_permissions(self, *permissions: str) -> Callable[[Request], User]:
        return self.always_authorized

class TokenOAuth2AuthModel(AuthModel):

    def __init__(self, user_model: UserModel, role_model: RoleModel, token_model: TokenModel, oauth2_scheme: OAuth2, root_permission: str):
        self.user_model = user_model
        self.role_model = role_model
        self.token_model = token_model
        self.oauth2_scheme = oauth2_scheme
        self.root_permission = root_permission

    def raise_unauthorized_exception(self, detail: str):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )

    def everyone(self) -> Callable[[Request], User]:
        async def verify_everyone(request: Request) -> User:
            try:
                token = self.oauth2_scheme(request)
                if token is None:
                    return self.user_model.get_guest_user()
                current_user = self.token_model.get_user_by_token(token)
                if not current_user:
                    current_user = self.user_model.get_guest_user()
                if not current_user.active:
                    self.raise_unauthorized_exception('User deactivated')
                return current_user 
            except:
                return self.user_model.get_guest_user()
        return verify_everyone 

    def is_authenticated(self) -> Callable[[Request], User]:
        async def verify_is_authenticated(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                self.raise_unauthorized_exception('Not authenticated')
            current_user = self.token_model.get_user_by_token(token)
            if not current_user:
                self.raise_unauthorized_exception('Not authenticated')
            if not current_user.active:
                self.raise_unauthorized_exception('User deactivated')
            return current_user
        return verify_is_authenticated

    def has_any_permissions(self, *permissions: str) -> Callable[[Request], User]:
        async def verify_has_any_permission(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                self.raise_unauthorized_exception('Not authenticated')
            current_user = self.token_model.get_user_by_token(token)
            if not current_user:
                self.raise_unauthorized_exception('Not authenticated')
            if not current_user.active:
                self.raise_unauthorized_exception('User deactivated')
            if len(permissions) == 0:
                return current_user
            if current_user.has_permission(self.root_permission):
                return current_user
            for permission in permissions:
                if current_user.has_permission(permission):
                    return current_user
            current_user_role_ids = current_user.role_ids
            for current_user_role_id in current_user_role_ids:
                current_user_role = self.role_model.find_by_id(current_user_role_id)
                if current_user_role.has_permission(permission):
                    return current_user
            self.raise_unauthorized_exception('Unauthorized')
        return verify_has_any_permission