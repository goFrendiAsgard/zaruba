from core.security.rule.test_default_auth_rule_util import (
    init_test_default_auth_rule_components
)
from module.auth.user.test_default_user_service_util import (
    AUTHORIZED_ACTIVE_USER, AUTHORIZED_INACTIVE_USER,
    UNAUTHORIZED_ACTIVE_USER, UNAUTHORIZED_INACTIVE_USER
)
from schema.auth_type import AuthType


def test_default_auth_rule_anyone():
    auth_rule, _ = init_test_default_auth_rule_components()
    assert auth_rule.check_user_access(
        None, AuthType.ANYONE
    )
    assert auth_rule.check_user_access(
        UNAUTHORIZED_ACTIVE_USER, AuthType.ANYONE
    )
    assert auth_rule.check_user_access(
        UNAUTHORIZED_INACTIVE_USER, AuthType.ANYONE
    )
    assert auth_rule.check_user_access(
        AUTHORIZED_ACTIVE_USER, AuthType.ANYONE
    )
    assert auth_rule.check_user_access(
        AUTHORIZED_INACTIVE_USER, AuthType.ANYONE
    )


def test_default_auth_rule_visitor():
    auth_rule, _ = init_test_default_auth_rule_components()
    assert auth_rule.check_user_access(
        None, AuthType.VISITOR
    )
    assert not auth_rule.check_user_access(
        UNAUTHORIZED_ACTIVE_USER, AuthType.VISITOR
    )
    assert auth_rule.check_user_access(
        UNAUTHORIZED_INACTIVE_USER, AuthType.VISITOR
    )
    assert not auth_rule.check_user_access(
        AUTHORIZED_ACTIVE_USER, AuthType.VISITOR
    )
    assert auth_rule.check_user_access(
        AUTHORIZED_INACTIVE_USER, AuthType.VISITOR
    )


def test_default_auth_rule_user():
    auth_rule, _ = init_test_default_auth_rule_components()
    assert not auth_rule.check_user_access(
        None, AuthType.USER
    )
    assert auth_rule.check_user_access(
        UNAUTHORIZED_ACTIVE_USER, AuthType.USER
    )
    assert not auth_rule.check_user_access(
        UNAUTHORIZED_INACTIVE_USER, AuthType.USER
    )
    assert auth_rule.check_user_access(
        AUTHORIZED_ACTIVE_USER, AuthType.USER
    )
    assert not auth_rule.check_user_access(
        AUTHORIZED_INACTIVE_USER, AuthType.USER
    )


def test_default_auth_rule_has_permission():
    auth_rule, _ = init_test_default_auth_rule_components()
    assert not auth_rule.check_user_access(
        None, AuthType.HAS_PERMISSION, 'permission'
    )
    assert not auth_rule.check_user_access(
        UNAUTHORIZED_ACTIVE_USER, AuthType.HAS_PERMISSION, 'permission'
    )
    assert not auth_rule.check_user_access(
        UNAUTHORIZED_INACTIVE_USER, AuthType.HAS_PERMISSION, 'permission'
    )
    assert auth_rule.check_user_access(
        AUTHORIZED_ACTIVE_USER, AuthType.HAS_PERMISSION, 'permission'
    )
    assert not auth_rule.check_user_access(
        AUTHORIZED_INACTIVE_USER, AuthType.HAS_PERMISSION, 'permission'
    )
