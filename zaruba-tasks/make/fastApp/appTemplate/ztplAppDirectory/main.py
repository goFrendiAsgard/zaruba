from fastapi.security import OAuth2PasswordBearer
from sqlalchemy import create_engine
from schemas.user import UserData
from modules.auth import (
    register_auth_api_route, register_auth_ui_route, register_auth_event_handler, register_auth_rpc_handler,
    SessionService, TokenOAuth2AuthService, JWTTokenService, DefaultUserService, UserSeederService, RoleService,
    DBRoleRepo, DBUserRepo
)
from configs import (
    # feature flags
    enable_api, enable_auth_module, enable_event_handler, enable_route_handler, enable_rpc_handler,
    enable_ui, seed_root_user,
    # factories
    create_app, create_menu_service, create_message_bus, create_rpc, create_page_template,
    # db
    db_create_all, db_url,
    # messagebus + rpc
    message_bus_type, rpc_type,
    # url
    create_access_token_url_path, create_oauth_access_token_url_path, renew_access_token_url_path,
    # auth
    guest_username, root_initial_email, root_initial_fullname, root_initial_password, 
    root_initial_phone_number, root_username, root_permission, access_token_algorithm,
    access_token_expire, access_token_secret_key
)

import os

################################################
# -- 🚌 Message bus and RPC initialization
################################################
mb = create_message_bus(message_bus_type)
rpc = create_rpc(rpc_type)

################################################
# -- 🛢️ Database engine initialization
################################################
engine = create_engine(db_url, echo=True)

################################################
# -- 🔒 Auth service initialization
################################################
oauth2_scheme = OAuth2PasswordBearer(tokenUrl = create_oauth_access_token_url_path, auto_error = False)
auth_service = TokenOAuth2AuthService(rpc, oauth2_scheme)

################################################
# -- 👓 User Interface initialization
################################################
menu_service = create_menu_service(rpc, auth_service)
page_template = create_page_template()

################################################
# -- ⚛️ FastAPI initialization
################################################
app = create_app(mb, rpc, page_template)

################################################
# -- 🔒 Auth module
################################################
if enable_auth_module:
    role_repo = DBRoleRepo(engine=engine, create_all=db_create_all)
    user_repo = DBUserRepo(engine=engine, create_all=db_create_all)
    role_service = RoleService(role_repo)
    user_service = DefaultUserService(user_repo, role_service, guest_username, root_permission=root_permission)
    if seed_root_user:
        user_seeder_service = UserSeederService(user_service)
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
        register_auth_api_route(app, mb, rpc, auth_service, create_oauth_access_token_url_path, create_access_token_url_path, renew_access_token_url_path)
    # UI route
    if enable_route_handler and enable_ui:
        register_auth_ui_route(app, mb, rpc, menu_service, page_template, create_access_token_url_path)
    # handle event
    if enable_event_handler:
        register_auth_event_handler(mb)
    # serve RPC
    if enable_rpc_handler:
        register_auth_rpc_handler(rpc, role_service, user_service, token_service, session_service)
