from typing import Callable, Optional
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from helpers.transport import RPC
from starlette.requests import Request
from schemas.user import User
from modules.auth.auth.authService import AuthService

import traceback


class TokenOAuth2AuthService(AuthService):

    def __init__(self, rpc: RPC, oauth2_scheme: OAuth2):
        self.rpc = rpc
        self.oauth2_scheme = oauth2_scheme

    def _raise_error_or_return_none(self, throw_error: bool, status_code: int, detail: str) -> None:
        if not throw_error:
            return None
        raise HTTPException(
            status_code=status_code,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )

    def _get_user_by_token(self, token: str) -> Optional[User]:
        try:
            user_data = self.rpc.call('get_user_by_token', token)
            return None if user_data is None else User.parse_obj(user_data)
        except:
            print(traceback.format_exc)
            return None

    def _is_user_authorized(self, user: User, permission: str) -> bool:
        user_data = user.dict()
        return self.rpc.call('is_user_authorized', user_data, permission)

    def everyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_everyone(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            if bearer_token is None and app_access_token is None:
                return None
            token = bearer_token if bearer_token is not None else app_access_token
            return self._get_user_by_token(token)
        return verify_everyone 

    def is_unauthenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_unauthenticated(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            if bearer_token is None and app_access_token is None:
                return None
            token = bearer_token if bearer_token is not None else app_access_token
            current_user = self._get_user_by_token(token)
            if not current_user or not current_user.active:
                return None
            return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
        return verify_is_unauthenticated

    def is_authenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_authenticated(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            if bearer_token is None and app_access_token is None:
                return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
            token = bearer_token if bearer_token is not None else app_access_token
            current_user = self._get_user_by_token(token)
            if not current_user or not current_user.active:
                return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
            return current_user
        return verify_is_authenticated

    def is_authorized(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_authorized(current_user = Depends(self.is_authenticated(throw_error=throw_error))) -> Optional[User]:
            if not current_user:
                return self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Not authenticated')
            if self._is_user_authorized(current_user, permission):
                return current_user
            self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Unauthorized')
        return verify_is_authorized
   