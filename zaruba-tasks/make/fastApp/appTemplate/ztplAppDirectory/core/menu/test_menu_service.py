from schema.auth_type import AuthType
from core.menu.test_menu_service_util import UNAUTHORIZED_ACTIVE_USER, AUTHORIZED_ACTIVE_USER, MenuTestCase, init_test_menu_service_components, init_test_menu_data, check_has_access

import pytest


def test_menu_service_add_menu_with_invalid_auth_type():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    is_error = False
    try:
        menu_service.add_menu(name='test', title='test', url='/test', auth_type= 500)
    except:
        is_error = True
    assert is_error


def test_menu_service_add_menu_with_non_existing_parent():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    is_error = False
    try:
        menu_service.add_menu(name='test', title='test', url='/test', auth_type= AuthType.ANYONE, parent_name='invalid')
    except:
        is_error = True
    assert is_error


def test_menu_service_get_accessible_non_existing_menu():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    # test get accessible menu for not existing menu
    is_error = False
    try:
        menu_service.get_accessible_menu('invalid', None)
    except:
        is_error = True
    assert is_error


@pytest.mark.asyncio
async def test_menu_service_no_user_get_accessible_everyone_menu():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('everyone-everyone', None)
    assert len(root.submenus) == 2
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=True, submenus_count=2, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=True),
            MenuTestCase(name='everyone-unauthenticated', is_highlighted=False),
        ]),
        MenuTestCase(name='unauthenticated', is_highlighted=False, submenus_count=2, children = [
            MenuTestCase(name='unauthenticated-everyone', is_highlighted=False),
            MenuTestCase(name='unauthenticated-unauthenticated', is_highlighted=False),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_no_user_get_accessible_unauthenticated_menu():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('unauthenticated-unauthenticated', None)
    assert len(root.submenus) == 2
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=False, submenus_count=2, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=False),
            MenuTestCase(name='everyone-unauthenticated', is_highlighted=False),
        ]),
        MenuTestCase(name='unauthenticated', is_highlighted=True, submenus_count=2, children = [
            MenuTestCase(name='unauthenticated-everyone', is_highlighted=False),
            MenuTestCase(name='unauthenticated-unauthenticated', is_highlighted=True),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_authenticated_user_get_accessible_authenticated_menu():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('authenticated-authenticated', UNAUTHORIZED_ACTIVE_USER)
    assert len(root.submenus) == 2
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=False, submenus_count=2, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=False),
            MenuTestCase(name='everyone-authenticated', is_highlighted=False),
        ]),
        MenuTestCase(name='authenticated', is_highlighted=True, submenus_count=2, children = [
            MenuTestCase(name='authenticated-everyone', is_highlighted=False),
            MenuTestCase(name='authenticated-authenticated', is_highlighted=True),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_authorized_user_get_accessible_authorized_menu():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    # test get accessible menu for authorized subbmenu
    root = menu_service.get_accessible_menu('authorized-authorized', AUTHORIZED_ACTIVE_USER)
    assert len(root.submenus) == 3
    test_cases = [
        MenuTestCase(name='everyone', is_highlighted=False, submenus_count=3, children = [
            MenuTestCase(name='everyone-everyone', is_highlighted=False),
            MenuTestCase(name='everyone-authenticated', is_highlighted=False),
            MenuTestCase(name='everyone-authorized', is_highlighted=False),
        ]),
        MenuTestCase(name='authenticated', is_highlighted=False, submenus_count=3, children = [
            MenuTestCase(name='authenticated-everyone', is_highlighted=False),
            MenuTestCase(name='authenticated-authenticated', is_highlighted=False),
            MenuTestCase(name='authenticated-authorized', is_highlighted=False),
        ]),
        MenuTestCase(name='authorized', is_highlighted=True, submenus_count=3, children = [
            MenuTestCase(name='authorized-everyone', is_highlighted=False),
            MenuTestCase(name='authorized-authenticated', is_highlighted=False),
            MenuTestCase(name='authorized-authorized', is_highlighted=True),
        ])
    ]
    for index, test_case in enumerate(test_cases):
        test_case.assert_menu(root.submenus[index])


@pytest.mark.asyncio
async def test_menu_service_no_user_authorize():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    await check_has_access(menu_service, None, {
        'everyone' : True,
        'unauthenticated' : True,
        'authenticated' : False,
        'authorized': False,
        'invalid': False
    })


@pytest.mark.asyncio
async def test_menu_service_authenticated_user_authorize():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    await check_has_access(menu_service, UNAUTHORIZED_ACTIVE_USER, {
        'everyone' : True,
        'unauthenticated' : False,
        'authenticated' : True,
        'authorized': False,
        'invalid': False
    })


@pytest.mark.asyncio
async def test_menu_service_authorized_user_authorize():
    menu_service, _ = init_test_menu_service_components()
    init_test_menu_data(menu_service)
    await check_has_access(menu_service, AUTHORIZED_ACTIVE_USER, {
        'everyone' : True,
        'unauthenticated' : False,
        'authenticated' : True,
        'authorized': True,
        'invalid': False
    })
