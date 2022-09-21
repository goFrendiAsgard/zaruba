from typing import Callable, Optional
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from helpers.transport import RPC
from starlette.requests import Request
from schemas.user import User
from modules.auth.auth.authService import AuthService


class NoAuthService(AuthService):

    def __init__(self, rpc: RPC):
        self.rpc = rpc

    def _get_guest_user(self) -> User:
        guest_user_data = self.rpc.call('get_guest_user')
        return User.parse_obj(guest_user_data)

    def everyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._get_guest_user()

    def is_unauthenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        if throw_error:
            raise HTTPException(
                status_code=status.HTTP_403_FORBIDDEN,
                detail='Forbidden'
            )
        return None

    def is_authenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._get_guest_user()

    def is_authorized(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._get_guest_user()

