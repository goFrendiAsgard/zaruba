from typing import Callable, Optional
from starlette.requests import Request
from schemas.user import User

import abc

class AuthService(abc.ABC):

    @abc.abstractmethod
    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        '''
        Return boolean, representing whether current_user should pass the authentication/authorization or not.
        '''
        pass

    @abc.abstractmethod
    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Always return current_user or throw error if something goes wrong.
        '''
        pass

    @abc.abstractmethod
    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Only return current user if current user is considered authenticated.
        Otherwise this will return None or throw error, depending on `throw_error` parameter
        '''
        pass

    @abc.abstractmethod
    def is_visitor(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Only return current user if current user is not authenticated.
        Otherwise this will return None or throw error, depending on `throw_error` parameter
        '''
        pass

    @abc.abstractmethod
    def has_permission(self, permission_name: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Only return current user if current user is authenticated and has required permission.
        Otherwise this will return None or throw error, depending on `throw_error` parameter
        '''
        pass
