from typing import Callable, Optional
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from helpers.transport import RPC
from core.security.authService import AuthService
from starlette.requests import Request
from schemas.user import User


class NoAuthService(AuthService):

    def __init__(self, rpc: RPC):
        self.rpc = rpc


    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._allow(throw_error)


    def is_visitor(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._restrict(throw_error)


    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._restrict(throw_error)


    def has_permission(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        return self._allow(throw_error)

    
    def _allow(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def allow(request: Request) -> Optional[User]:
            guest_user_data = self.rpc.call('get_guest_user')
            return User.parse_obj(guest_user_data)
        return allow


    def _restrict(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        def restrict(request: Request) -> Optional[User]:
            if throw_error:
                raise HTTPException(
                    status_code=status.HTTP_403_FORBIDDEN,
                    detail='Forbidden'
                )
            return None
        return restrict

