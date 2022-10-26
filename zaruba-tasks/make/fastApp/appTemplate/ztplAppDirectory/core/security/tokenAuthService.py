from typing import Callable, Optional
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from helpers.transport import RPC
from starlette.requests import Request
from schemas.user import User
from core.security.authService import AuthService

import sys
import traceback


class TokenAuthService(AuthService):

    def __init__(self, rpc: RPC, oauth2_scheme: OAuth2):
        self.rpc = rpc
        self.oauth2_scheme = oauth2_scheme


    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_everyone(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            current_user = self._get_user(bearer_token, app_access_token, throw_error=False)
            if not current_user or not current_user.active:
                return None
            return current_user
        return verify_everyone 


    def is_visitor(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_unauthenticated(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            current_user = self._get_user(bearer_token, app_access_token, throw_error=False)
            if not current_user or not current_user.active:
                return None
            return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
        return verify_is_unauthenticated


    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_authenticated(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            return self._get_authenticated_user(bearer_token, app_access_token, throw_error)
        return verify_is_authenticated


    def has_permission(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_authorized(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            authenticated_user = self._get_authenticated_user(bearer_token, app_access_token, throw_error)
            if not authenticated_user or not authenticated_user.active:
                return self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Not authenticated')
            if self._is_user_authorized(authenticated_user, permission):
                return authenticated_user
            self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Unauthorized')
        return verify_is_authorized

    
    def _get_user(self, bearer_token: Optional[str], app_access_token: Optional[str], throw_error: bool = True) -> Optional[User]:
        if bearer_token is None and app_access_token is None:
            return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
        token = bearer_token if bearer_token is not None else app_access_token
        return self._get_user_by_token(token)
 

    def _get_authenticated_user(self, bearer_token: Optional[str], app_access_token: Optional[str], throw_error: bool = True) -> Optional[User]:
        authenticated_user = self._get_user(bearer_token, app_access_token, throw_error)
        if not authenticated_user or not authenticated_user.active:
            return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
        return authenticated_user

   
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
            user_data = self.rpc.call('get_user_by_access_token', token)
            return None if user_data is None else User.parse_obj(user_data)
        except:
            print('Error while fetching user with token {token}'.format(token=token), file=sys.stderr)
            print(traceback.format_exc(), file=sys.stderr)
            return None


    def _is_user_authorized(self, user: User, permission: str) -> bool:
        user_data = user.dict()
        return self.rpc.call('is_user_authorized', user_data, permission)
