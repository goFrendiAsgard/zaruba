from typing import Optional
from schemas.user import User
from core.security.rule.authRule import AuthRule
from schemas.authType import AuthType

class NoAuthRule(AuthRule):

    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        if auth_type == AuthType.ANYONE:
            return True
        if auth_type == AuthType.VISITOR:
            return False
        if auth_type == AuthType.USER:
            return False
        if auth_type == AuthType.HAS_PERMISSION:
            return True
        return False

