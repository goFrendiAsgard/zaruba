from fastapi.security import OAuth2PasswordBearer
from helper.config import get_boolean_env
from sqlalchemy import create_engine
from schema.user import UserData
from schema.content_type import ContentTypeData
from core import (
    register_session_api_route, register_session_ui_route, register_session_rpc,
    DefaultAuthRule, DefaultUserFetcher,
    AuthService, MenuService, SessionService, JWTTokenService,
)
from module.log import (
    register_log_api_route, register_log_ui_route, register_log_event_handler, register_log_rpc_handler,
    ActivityService, DBActivityRepo
)
from module.auth import (
    register_auth_api_route, register_auth_ui_route, register_auth_event_handler, register_auth_rpc_handler,
    DefaultUserService, UserSeederService, RoleService,
    DBRoleRepo, DBUserRepo
)
from module.cms import (
    register_cms_api_route, register_cms_ui_route, register_cms_event_handler, register_cms_rpc_handler,
    ContentTypeService, ContentService, ContentTypeSeederService,
    DBContentTypeRepo, DBContentRepo
)
from config import (
    # feature flags
    enable_api, enable_auth_module, enable_cms_module, enable_log_module,
    enable_event_handler, enable_route_handler, enable_rpc_handler, enable_ui, seed_root_user,
    # factories
    create_app, create_message_bus, create_rpc, create_page_template,
    # db
    db_create_all, db_url,
    # messagebus + rpc
    message_bus_type, rpc_type,
    # url
    create_access_token_url_path, create_oauth_access_token_url_path, renew_access_token_url_path,
    # auth
    root_initial_email, root_initial_fullname, root_initial_password, 
    root_initial_phone_number, root_username, root_permission, access_token_algorithm,
    access_token_expire, access_token_secret_key,
    # cms
    seed_default_content_type, default_content_type_name,
    # activity
    activity_events
)

import os

################################################
# -- 🚌 Message bus and RPC initialization
################################################
mb = create_message_bus(message_bus_type, activity_events)
rpc = create_rpc(rpc_type)

################################################
# -- 🛢️ Database engine initialization
################################################
engine = create_engine(db_url, echo=True)

################################################
# -- 🔒 Auth service initialization
################################################
oauth2_scheme = OAuth2PasswordBearer(tokenUrl = create_oauth_access_token_url_path, auto_error = False)
auth_rule = DefaultAuthRule(rpc)
user_fetcher = DefaultUserFetcher(rpc, oauth2_scheme)
auth_service = AuthService(auth_rule, user_fetcher, root_permission)

################################################
# -- 👓 User Interface initialization
################################################
menu_service = MenuService(rpc, auth_service)
page_template = create_page_template()

################################################
# -- ⚛️ FastAPI initialization
################################################
app = create_app(mb, rpc, menu_service, page_template)
# session API
if enable_route_handler:
    register_session_api_route(
        app, mb, rpc, auth_service,
        create_oauth_access_token_url_path, create_access_token_url_path, renew_access_token_url_path
    )
# session page
if enable_route_handler and enable_ui:
    register_session_ui_route(app, mb, rpc, menu_service, page_template, create_access_token_url_path)

################################################
# -- ✍️ Log module
################################################
# Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
if enable_log_module:
    activity_repo = DBActivityRepo(engine=engine, create_all=db_create_all)
    activity_service = ActivityService(mb, rpc, auth_service, activity_repo)
    # API route
    if enable_route_handler and enable_api:
        register_log_api_route(app, mb, rpc, auth_service)
    # UI route
    if enable_route_handler and enable_ui:
        register_log_ui_route(app, mb, rpc, menu_service, page_template)
    # handle event
    if enable_event_handler:
        register_log_event_handler(mb, rpc, auth_service, activity_service)
    # serve RPC
    if enable_rpc_handler:
        # Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
        register_log_rpc_handler(mb, rpc, auth_service, activity_service)

################################################
# -- 🔒 Auth module
################################################
# Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
if enable_auth_module:
    role_repo = DBRoleRepo(engine=engine, create_all=db_create_all)
    user_repo = DBUserRepo(engine=engine, create_all=db_create_all)
    role_service = RoleService(mb, rpc, role_repo)
    user_service = DefaultUserService(mb, rpc, user_repo, role_service, root_permission=root_permission)
    if seed_root_user:
        user_seeder_service = UserSeederService(auth_service, user_service)
        user_seeder_service.seed(UserData(
            username = root_username,
            email = root_initial_email,
            phone_number = root_initial_phone_number,
            password = root_initial_password,
            active = True,
            permissions = [root_permission],
            full_name = root_initial_fullname
        ))
    token_service = JWTTokenService(
        user_service = user_service,
        access_token_secret_key = access_token_secret_key,
        access_token_algorithm = access_token_algorithm,
        access_token_expire = access_token_expire
    )
    session_service = SessionService(user_service, token_service)
    # API route
    if enable_route_handler and enable_api:
        register_auth_api_route(app, mb, rpc, auth_service)
    # UI route
    if enable_route_handler and enable_ui:
        register_auth_ui_route(app, mb, rpc, menu_service, page_template, create_access_token_url_path)
    # handle event
    if enable_event_handler:
        register_auth_event_handler(mb, rpc, auth_service)
    # serve RPC
    if enable_rpc_handler:
        register_auth_rpc_handler(mb, rpc, auth_service, role_service, user_service)
        register_session_rpc(mb, rpc, auth_service, session_service)

################################################
# -- 📰 CMS module
################################################
# Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
if enable_cms_module:
    content_type_repo = DBContentTypeRepo(engine=engine, create_all=db_create_all)
    content_type_service = ContentTypeService(mb, rpc, content_type_repo)
    content_repo = DBContentRepo(engine=engine, create_all=db_create_all)
    content_service = ContentService(mb, rpc, content_repo, content_type_repo)
    # Seed default content type
    if seed_default_content_type:
        content_type_seeder_service = ContentTypeSeederService(auth_service, content_type_service)
        content_type_seeder_service.seed(ContentTypeData(name=default_content_type_name))
    # API route
    if enable_route_handler and enable_api:
        register_cms_api_route(app, mb, rpc, auth_service)
    # UI route
    if enable_route_handler and enable_ui:
        register_cms_ui_route(app, mb, rpc, menu_service, page_template)
    # handle event
    if enable_event_handler:
        register_cms_event_handler(mb, rpc, auth_service)
    # serve RPC
    if enable_rpc_handler:
        # Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
        register_cms_rpc_handler(mb, rpc, auth_service, content_type_service, content_service)
