from typing import Callable, List, Optional
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
    def is_authorized(self, *permissions: str) -> Callable[[Request], User]:
        pass

    @abc.abstractmethod
    def is_user_authorized(self, user: Optional[User], permissions: List[str]) -> bool:
        pass

class NoAuthService(AuthService):

    def __init__(self, user_service: UserService, root_permission: str = 'root'):
        self.user_service: str = user_service

    def _always_authorized(self, Request) -> User:
        return self.user_service.get_guest_user()

    def everyone(self) -> Callable[[Request], User]:
        return self._always_authorized

    def is_authenticated(self) -> Callable[[Request], User]:
        return self._always_authorized

    def is_authorized(self, *permissions: str) -> Callable[[Request], User]:
        return self._always_authorized

    def is_user_authorized(self, user: Optional[User], permissions: List[str]) -> bool:
        return True

class TokenOAuth2AuthService(AuthService):

    def __init__(self, user_service: UserService, role_service: RoleService, token_service: TokenService, oauth2_scheme: OAuth2, root_permission: str = 'root'):
        self.user_service = user_service
        self.role_service = role_service
        self.token_service = token_service
        self.oauth2_scheme = oauth2_scheme
        self.root_permission = root_permission

    def _raise_unauthorized_exception(self, detail: str):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )

    def everyone(self) -> Callable[[Request], User]:
        async def verify_everyone(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                return self.user_service.get_guest_user()
            current_user = self.token_service.get_user_by_token(token)
            if not current_user or not current_user.active:
                return self.user_service.get_guest_user()
            return current_user
        return verify_everyone 

    def is_authenticated(self) -> Callable[[Request], User]:
        async def verify_is_authenticated(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                self._raise_unauthorized_exception('Not authenticated')
            current_user = self.token_service.get_user_by_token(token)
            if not current_user:
                self._raise_unauthorized_exception('Not authenticated')
            if not current_user.active:
                self._raise_unauthorized_exception('User deactivated')
            return current_user
        return verify_is_authenticated

    def is_authorized(self, *permissions: str) -> Callable[[Request], User]:
        async def verify_has_any_permission(token = Depends(self.oauth2_scheme)) -> User:
            if token is None:
                self._raise_unauthorized_exception('Not authenticated')
            current_user = self.token_service.get_user_by_token(token)
            if not current_user:
                self._raise_unauthorized_exception('Not authenticated')
            if not current_user.active:
                self._raise_unauthorized_exception('User deactivated')
            if self.is_user_authorized(current_user, permissions):
                return current_user
            self._raise_unauthorized_exception('Unauthorized')
        return verify_has_any_permission
    
    def is_user_authorized(self, user: Optional[User], permissions: List[str]) -> bool:
        # no permission
        if len(permissions) == 0:
            return True
        # no user
        if user is None:
            return False
        # user has root permission
        if user.has_permission(self.root_permission):
            return True
        # user has any required permission
        for permission in permissions:
            if user.has_permission(permission):
                return True
        # user has any role that has any required permission
        role_ids = user.role_ids
        for role_id in role_ids:
            role = self.role_service.find_by_id(role_id)
            if role.has_permission(permission):
                return True