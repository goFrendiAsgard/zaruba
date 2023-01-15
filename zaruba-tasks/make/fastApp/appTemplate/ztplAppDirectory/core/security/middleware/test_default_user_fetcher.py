from module.auth.user.test_default_user_service_util import (
    AUTHORIZED_ACTIVE_USER
)
from core.security.middleware.test_default_user_fetcher_util import (
    init_test_default_user_fetcher_components
)
import pytest


@pytest.mark.asyncio
async def test_default_user_fetcher_with_no_parameter():
    user_fetcher, _, _ = init_test_default_user_fetcher_components()
    fetch_user = user_fetcher.get_user_fetcher()
    user = await fetch_user()
    assert user is None


@pytest.mark.asyncio
async def test_default_user_fetcher_with_bearer_token():
    user_fetcher, _, _ = init_test_default_user_fetcher_components()
    fetch_user = user_fetcher.get_user_fetcher()
    user = await fetch_user(bearer_token='authorized_active')
    assert user == AUTHORIZED_ACTIVE_USER


@pytest.mark.asyncio
async def test_default_user_fetcher_with_cookie():
    user_fetcher, _, _ = init_test_default_user_fetcher_components()
    fetch_user = user_fetcher.get_user_fetcher()
    user = await fetch_user(
        bearer_token=None,
        app_cred_token='authorized_active'
    )
    assert user == AUTHORIZED_ACTIVE_USER


@pytest.mark.asyncio
async def test_default_user_fetcher_with_throw_error():
    user_fetcher, _, _ = init_test_default_user_fetcher_components()
    fetch_user = user_fetcher.get_user_fetcher(throw_error=True)
    is_error = False
    try:
        await fetch_user(bearer_token='error')
    except Exception:
        is_error = True
    assert is_error


@pytest.mark.asyncio
async def test_default_user_fetcher_without_throw_error():
    user_fetcher, _, _ = init_test_default_user_fetcher_components()
    fetch_user = user_fetcher.get_user_fetcher(throw_error=False)
    user = await fetch_user(bearer_token='error')
    assert user is None
