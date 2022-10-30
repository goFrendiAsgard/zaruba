from typing import Callable, Optional
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from helpers.transport import RPC
from core.security.authService import AuthService
from starlette.requests import Request
from schemas.user import User
from schemas.authType import AuthType


class NoAuthService(AuthService):

    def __init__(self, rpc: RPC):
        self.rpc = rpc

    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        if auth_type == AuthType.ANYONE:
            return True
        if auth_type == AuthType.VISITOR:
            return False
        if auth_type == AuthType.USER:
            return False
        if auth_type == AuthType.HAS_PERMISSION:
            return True
        return False


    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._get_access_validator(AuthType.ANYONE, throw_error=throw_error)


    def is_visitor(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._get_access_validator(AuthType.VISITOR, throw_error=throw_error)


    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._get_access_validator(AuthType.USER, throw_error=throw_error)


    def has_permission(self, permission_name: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._get_access_validator(AuthType.HAS_PERMISSION, permission_name=permission_name, throw_error=throw_error)
    

    def _get_access_validator(self, auth_type: int, permission_name: Optional[str] = None, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def check_user_access(result: Request) -> Optional[User]:
            if self.check_user_access(current_user=None, auth_type=auth_type, permission_name=permission_name):
                guest_user_data = self.rpc.call('get_guest_user')
                return User.parse_obj(guest_user_data)
            if throw_error:
                raise HTTPException(
                    status_code=status.HTTP_403_FORBIDDEN,
                    detail='Forbidden'
                )
            return None
        return check_user_access
