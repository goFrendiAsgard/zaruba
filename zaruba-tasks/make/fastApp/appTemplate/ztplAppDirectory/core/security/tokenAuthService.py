from typing import Callable, Optional
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from helpers.transport import RPC
from starlette.requests import Request
from schemas.user import User
from schemas.authType import AuthType
from core.security.authService import AuthService

import sys
import traceback


class TokenAuthService(AuthService):

    def __init__(self, rpc: RPC, oauth2_scheme: OAuth2):
        self.rpc = rpc
        self.oauth2_scheme = oauth2_scheme


    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        if auth_type == AuthType.ANYONE:
            return True
        if auth_type == AuthType.VISITOR:
            return current_user is None or not current_user.active
        if auth_type == AuthType.USER:
            return current_user is not None and current_user.active
        if auth_type == AuthType.HAS_PERMISSION:
            if current_user is None or not current_user.active:
                return False
            current_user_data = current_user.dict()
            return self.rpc.call('is_user_authorized', current_user_data, permission_name)
        return False


    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_everyone(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            current_user = self._get_user_by_token(bearer_token, app_access_token, throw_error=False)
            if self.check_user_access(current_user, AuthType.ANYONE):
                return current_user if current_user is not None and current_user.active else None
            return None
        return verify_everyone 


    def is_visitor(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_unauthenticated(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            current_user = self._get_user_by_token(bearer_token, app_access_token, throw_error=False)
            if self.check_user_access(current_user, AuthType.VISITOR):
                return current_user if current_user is not None and current_user.active else None
            return self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Forbidden')
        return verify_is_unauthenticated


    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_authenticated(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            current_user = self._get_user_by_token(bearer_token, app_access_token, throw_error=False)
            if self.check_user_access(current_user, AuthType.USER):
                return current_user if current_user is not None and current_user.active else None
            return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
        return verify_is_authenticated


    def has_permission(self, permission_name: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        async def verify_is_authorized(bearer_token = Depends(self.oauth2_scheme), app_access_token=Cookie(default=None)) -> Optional[User]:
            current_user = self._get_user_by_token(bearer_token, app_access_token, throw_error=False)
            if self.check_user_access(current_user, AuthType.HAS_PERMISSION, permission_name):
                return current_user if current_user is not None and current_user.active else None
            return self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Forbidden')
        return verify_is_authorized

    
    def _get_user_by_token(self, bearer_token: Optional[str], app_access_token: Optional[str], throw_error: bool = True) -> Optional[User]:
        if bearer_token is None and app_access_token is None:
            return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
        token = bearer_token if bearer_token is not None else app_access_token
        try:
            user_data = self.rpc.call('get_user_by_access_token', token)
            return None if user_data is None else User.parse_obj(user_data)
        except:
            print('Error while fetching user with token {token}'.format(token=token), file=sys.stderr)
            print(traceback.format_exc(), file=sys.stderr)
            return None


    def _raise_error_or_return_none(self, throw_error: bool, status_code: int, detail: str) -> None:
        if not throw_error:
            return None
        raise HTTPException(
            status_code=status_code,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )
