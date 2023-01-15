from typing import Optional
from schema.user import User
from core.security.rule.auth_rule import AuthRule
from schema.auth_type import AuthType


class NoAuthRule(AuthRule):
    '''
    No authentication, all user are unauthenticated and assumed to have
    all permissions.

    - ANYONE: can be accessed by anyone.
    - VISITOR: can never be accessed
    - USER: can never be accessed
    - HAS_PERMISSION: can be accessed by anyone.
    '''

    def check_user_access(
        self,
        current_user: Optional[User],
        auth_type: int,
        permission_name: Optional[str] = None
    ) -> bool:
        '''
        Initiate NoAuthRule.

        - ANYONE: can be accessed by anyone.
        - VISITOR: can never be accessed
        - USER: can never be accessed
        - HAS_PERMISSION: can be accessed by anyone.
        '''
        if auth_type == AuthType.ANYONE:
            return True
        if auth_type == AuthType.VISITOR:
            return False
        if auth_type == AuthType.USER:
            return False
        if auth_type == AuthType.HAS_PERMISSION:
            return True
        return False
