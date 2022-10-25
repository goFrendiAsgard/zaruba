from typing import Callable, Optional
from starlette.requests import Request
from schemas.user import User

import abc

class AuthService(abc.ABC):

    @abc.abstractmethod
    def anyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass

    @abc.abstractmethod
    def is_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass

    @abc.abstractmethod
    def is_not_user(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass

    @abc.abstractmethod
    def has_permission(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass
