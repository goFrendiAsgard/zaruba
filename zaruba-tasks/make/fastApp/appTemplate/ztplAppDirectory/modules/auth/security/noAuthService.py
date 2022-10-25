from typing import Callable, Optional
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from helpers.transport import RPC
from starlette.requests import Request
from schemas.user import User
from modules.auth.security.authService import AuthService


class NoAuthService(AuthService):

    def __init__(self, rpc: RPC):
        self.rpc = rpc

    def _allow_everyone_as_guest(self, request: Request) -> Optional[User]:
        guest_user_data = self.rpc.call('get_guest_user')
        return User.parse_obj(guest_user_data)

    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._allow_everyone_as_guest

    def is_not_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def restrict(request: Request) -> Optional[User]:
            if throw_error:
                raise HTTPException(
                    status_code=status.HTTP_403_FORBIDDEN,
                    detail='Forbidden'
                )
            return None
        return restrict

    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._allow_everyone_as_guest

    def has_permission(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._allow_everyone_as_guest

