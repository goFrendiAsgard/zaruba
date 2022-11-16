from typing import Optional
from schema.user import User

import abc

class AuthRule(abc.ABC):

    @abc.abstractmethod
    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        '''
        Return boolean, representing whether current_user should pass the authentication/authorization or not.
        '''
        pass
