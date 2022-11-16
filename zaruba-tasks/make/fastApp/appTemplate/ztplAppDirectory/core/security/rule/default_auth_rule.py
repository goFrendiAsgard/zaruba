from typing import Optional
from schema.user import User
from core.security.rule.auth_rule import AuthRule
from schema.auth_type import AuthType
from helper.transport.rpc import RPC

class DefaultAuthRule(AuthRule):

    def __init__(self, rpc: RPC):
        self.rpc = rpc


    def check_user_access(self, current_user: Optional[User], auth_type: int, permission_name: Optional[str] = None) -> bool:
        if auth_type == AuthType.ANYONE:
            return True
        if auth_type == AuthType.VISITOR:
            return current_user is None or not current_user.active
        if auth_type == AuthType.USER:
            return current_user is not None and current_user.active
        if auth_type == AuthType.HAS_PERMISSION:
            if current_user is None or not current_user.active:
                return False
            current_user_data = current_user.dict()
            return self.rpc.call('is_user_authorized', current_user_data, permission_name)
        return False
