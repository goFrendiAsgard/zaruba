from typing import Tuple
from helpers.transport import RPC, LocalRPC
from core.security.rule.defaultAuthRule import DefaultAuthRule
from modules.auth.user.test_util import AUTHORIZED_ACTIVE_USER, AUTHORIZED_INACTIVE_USER
from schemas.user import User, UserData

rpc = LocalRPC()

@rpc.handle('is_user_authorized')
def is_user_authorized(user_data: UserData, permission: str) -> bool:
    user = User.parse_obj(user_data)
    return user.id in [AUTHORIZED_ACTIVE_USER.id, AUTHORIZED_INACTIVE_USER]


def init_test_default_auth_rule_components() -> Tuple[DefaultAuthRule, RPC]:
    default_auth_rule = DefaultAuthRule(rpc)
    return default_auth_rule, rpc
    

