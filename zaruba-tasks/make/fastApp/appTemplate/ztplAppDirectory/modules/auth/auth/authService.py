from typing import Callable, Optional
from starlette.requests import Request
from schemas.user import User

import abc

class AuthService(abc.ABC):

    @abc.abstractmethod
    def everyone(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass

    @abc.abstractmethod
    def is_authenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass

    @abc.abstractmethod
    def is_unauthenticated(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass

    @abc.abstractmethod
    def is_authorized(self, permission: str, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        pass
