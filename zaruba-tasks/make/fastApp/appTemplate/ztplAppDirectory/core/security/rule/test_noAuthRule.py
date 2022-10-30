from core.security.rule.noAuthRule import NoAuthRule
from modules.auth.user.test_defaultUserService_util import AUTHORIZED_ACTIVE_USER, AUTHORIZED_INACTIVE_USER, UNAUTHORIZED_ACTIVE_USER, UNAUTHORIZED_INACTIVE_USER
from schemas.authType import AuthType


def test_no_auth_rule_anyone():
    auth_rule = NoAuthRule()
    assert auth_rule.check_user_access(None, AuthType.ANYONE) == True
    assert auth_rule.check_user_access(UNAUTHORIZED_ACTIVE_USER, AuthType.ANYONE) == True
    assert auth_rule.check_user_access(UNAUTHORIZED_INACTIVE_USER, AuthType.ANYONE) == True
    assert auth_rule.check_user_access(AUTHORIZED_ACTIVE_USER, AuthType.ANYONE) == True
    assert auth_rule.check_user_access(AUTHORIZED_INACTIVE_USER, AuthType.ANYONE) == True


def test_no_auth_rule_visitor():
    auth_rule = NoAuthRule()
    assert auth_rule.check_user_access(None, AuthType.VISITOR) == False
    assert auth_rule.check_user_access(UNAUTHORIZED_ACTIVE_USER, AuthType.VISITOR) == False
    assert auth_rule.check_user_access(UNAUTHORIZED_INACTIVE_USER, AuthType.VISITOR) == False
    assert auth_rule.check_user_access(AUTHORIZED_ACTIVE_USER, AuthType.VISITOR) == False
    assert auth_rule.check_user_access(AUTHORIZED_INACTIVE_USER, AuthType.VISITOR) == False


def test_no_auth_rule_user():
    auth_rule = NoAuthRule()
    assert auth_rule.check_user_access(None, AuthType.USER) == False
    assert auth_rule.check_user_access(UNAUTHORIZED_ACTIVE_USER, AuthType.USER) == False
    assert auth_rule.check_user_access(UNAUTHORIZED_INACTIVE_USER, AuthType.USER) == False
    assert auth_rule.check_user_access(AUTHORIZED_ACTIVE_USER, AuthType.USER) == False
    assert auth_rule.check_user_access(AUTHORIZED_INACTIVE_USER, AuthType.USER) == False


def test_no_auth_rule_has_permission():
    auth_rule = NoAuthRule()
    assert auth_rule.check_user_access(None, AuthType.HAS_PERMISSION, 'permission') == True
    assert auth_rule.check_user_access(UNAUTHORIZED_ACTIVE_USER, AuthType.HAS_PERMISSION, 'permission') == True
    assert auth_rule.check_user_access(UNAUTHORIZED_INACTIVE_USER, AuthType.HAS_PERMISSION, 'permission') == True
    assert auth_rule.check_user_access(AUTHORIZED_ACTIVE_USER, AuthType.HAS_PERMISSION, 'permission') == True
    assert auth_rule.check_user_access(AUTHORIZED_INACTIVE_USER, AuthType.HAS_PERMISSION, 'permission') == True
    
    