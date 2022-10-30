from typing import Optional, Callable
from starlette.requests import Request
from schemas.user import User

import abc

class UserFetcher(abc.ABC):

    @abc.abstractmethod
    def get_user_fetcher(self, throw_error: bool = True) -> Callable[[Request], Optional[User]]:
        '''
        Return a function to fetch current user based on HTTP request.
        To be used with fastapi.Depends
        '''
        pass