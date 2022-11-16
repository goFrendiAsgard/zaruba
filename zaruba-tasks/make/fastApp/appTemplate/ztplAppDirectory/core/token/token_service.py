from typing import Optional
from schema.user import User

import abc

class TokenService(abc.ABC):

    @abc.abstractmethod
    def create_access_token(self, user: User, current_user: Optional[User]) -> str:
        pass

    @abc.abstractmethod
    def get_user_by_token(self, token: str, current_user: Optional[User]) -> Optional[User]:
        pass
