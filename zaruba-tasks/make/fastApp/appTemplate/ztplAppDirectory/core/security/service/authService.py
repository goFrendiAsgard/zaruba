from typing import Callable, Optional
from fastapi import Depends, HTTPException, status
from starlette.requests import Request
from schemas.user import User
from core.security.rule.authRule import AuthRule
from core.security.middleware.userFetcher import UserFetcher
from schemas.authType import AuthType

import datetime


class AuthService():

    def __init__(self, auth_rule: AuthRule, user_fetcher: UserFetcher, root_permission: str):
        self.auth_rule = auth_rule
        self.user_fetcher = user_fetcher
        self.earliest_date = datetime.datetime(1970, 1, 1, 0, 0, 0, 0, datetime.timezone.utc)
        self.root_permission = root_permission


    def get_guest_user(self) -> User:
        return User(
            id = 'guest',
            username = 'guest', 
            active = True,
            updated_at = self.earliest_date,
            created_at = self.earliest_date,
        )


    def get_system_user(self) -> User:
        return User(
            id = 'system',
            username = 'system', 
            active = True,
            permissions=[self.root_permission],
            updated_at = self.earliest_date,
            created_at = self.earliest_date,
        )


    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        '''
        Return boolean, representing whether current_user should pass the authentication/authorization or not.
        '''
        return self.auth_rule.check_user_access(current_user, auth_type, permission_name)

    
    def get_user_fetcher(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Return a function to fetch current user based on HTTP request.
        To be used with fastapi.Depends
        '''
        return self.user_fetcher.get_user_fetcher(throw_error)


    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Always return current_user or throw error if something goes wrong.
        '''
        async def check_anyone(current_user = Depends(self.get_user_fetcher(throw_error))) -> Optional[User]:
            if self.check_user_access(current_user, AuthType.ANYONE):
                return current_user if current_user is not None and current_user.active else None
            return self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Forbidden')
        return check_anyone


    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Only return current user if current user is considered authenticated.
        Otherwise this will return None or throw error, depending on `throw_error` parameter
        '''
        async def check_is_user(current_user = Depends(self.get_user_fetcher(throw_error))) -> Optional[User]:
            if self.check_user_access(current_user, AuthType.USER):
                return current_user if current_user is not None and current_user.active else None
            return self._raise_error_or_return_none(throw_error, status.HTTP_401_UNAUTHORIZED, 'Not authenticated')
        return check_is_user


    def is_visitor(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Only return current user if current user is not authenticated.
        Otherwise this will return None or throw error, depending on `throw_error` parameter
        '''
        async def check_is_visitor(current_user = Depends(self.get_user_fetcher(throw_error))) -> Optional[User]:
            if self.check_user_access(current_user, AuthType.VISITOR):
                return current_user if current_user is not None and current_user.active else None
            return self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Forbidden')
        return check_is_visitor


    def has_permission(self, permission_name: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Only return current user if current user is authenticated and has required permission.
        Otherwise this will return None or throw error, depending on `throw_error` parameter
        '''
        async def check_has_permission(current_user = Depends(self.get_user_fetcher(throw_error))) -> Optional[User]:
            if self.check_user_access(current_user, AuthType.HAS_PERMISSION, permission_name):
                return current_user if current_user is not None and current_user.active else None
            return self._raise_error_or_return_none(throw_error, status.HTTP_403_FORBIDDEN, 'Forbidden')
        return check_has_permission


    def _raise_error_or_return_none(self, throw_error: bool, status_code: int, detail: str) -> None:
        if not throw_error:
            return None
        raise HTTPException(
            status_code=status_code,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )

