from typing import Optional
from schema.user import User

import abc


class TokenService(abc.ABC):

    @abc.abstractmethod
    def create_cred_token(
        self,
        user: User,
        current_user: Optional[User]
    ) -> str:
        pass

    @abc.abstractmethod
    def get_user_by_cred_token(
        self,
        token: str,
        current_user: Optional[User]
    ) -> Optional[User]:
        pass
