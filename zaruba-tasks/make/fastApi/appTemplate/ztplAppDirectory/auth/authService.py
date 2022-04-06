from typing import Callable, List
from fastapi.security import OAuth2PasswordBearer, OAuth2
from fastapi import Depends, FastAPI, HTTPException, status
from starlette.requests import Request
from auth.roleService import RoleService
from auth.userService import UserService
from auth.tokenService import TokenService
from schemas.user import User
import abc

class AuthService(abc.ABC):

    @abc.abstractmethod
    def everyone(self) -> Callable[[Request], User]:
        pass

    @abc.abstractmethod
    def is_authenticated(self) -> Callable[[Request], User]:
        pass

    @abc.abstractmethod
    def has_any_permissions(self, *permissions: str) -> Callable[[Request], User]:
        pass

class NoAuthService(AuthService):

    def __init__(self, user_service: UserService):
        self.user_service = user_service

    def always_authorized(self, Request) -> User:
        return self.user_service.get_guest_user()

    def everyone(self) -> Callable[[Request], User]:
        return self.always_authorized

    def is_authenticated(self) -> Callable[[Request], User]:
        return self.always_authorized

    def has_any_permissions(self, *permissions: str) -> Callable[[Request], User]:
        return self.always_authorized

class TokenOAuth2AuthService(AuthService):

    def __init__(self, user_service: UserService, role_service: RoleService, token_service: TokenService, oauth2_scheme: OAuth2, root_permission: str):
        self.user_service = user_service
        self.role_service = role_service
        self.token_service = token_service
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
                    return self.user_service.get_guest_user()
                current_user = self.token_service.get_user_by_token(token)
                if not current_user:
                    current_user = self.user_service.get_guest_user()
                if not current_user.active:
                    self.raise_unauthorized_exception('User deactivated')
                return current_user 
            except:
                return self.user_service.get_guest_user()
        return verify_everyone 

    def is_authenticated(self) -> Callable[[Request], User]:
        async def verify_is_authenticated(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                self.raise_unauthorized_exception('Not authenticated')
            current_user = self.token_service.get_user_by_token(token)
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
            current_user = self.token_service.get_user_by_token(token)
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
                current_user_role = self.role_service.find_by_id(current_user_role_id)
                if current_user_role.has_permission(permission):
                    return current_user
            self.raise_unauthorized_exception('Unauthorized')
        return verify_has_any_permission