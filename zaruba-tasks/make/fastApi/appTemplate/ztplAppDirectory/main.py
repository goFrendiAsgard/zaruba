# -- 📖 Common import
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.security import OAuth2PasswordBearer
from fastapi.templating import Jinja2Templates
from sqlalchemy import create_engine
from helpers.transport import RMQEventMap, KafkaEventMap, KafkaAvroEventMap, create_kafka_connection_parameters, create_kafka_avro_connection_parameters, create_rmq_connection_parameters
from helpers.app import get_abs_static_dir, create_message_bus, create_rpc, handle_app_shutdown, register_static_dir_route_handler, register_readiness_handler
from repos.dbUser import DBUserRepo
from repos.dbRole import DBRoleRepo
from auth import register_auth_route_handler, register_auth_event_handler, register_auth_rpc_handler, TokenOAuth2AuthService, JWTTokenService, DefaultUserService, UserSeederService, RoleService
from schemas.user import UserData
from ui import MenuService, create_menu_service, register_template_exception_handler

import os
import json

error_threshold = int(os.getenv('APP_ERROR_THRESHOLD', '10'))

################################################
# -- 🐇 Rabbitmq setting
################################################
rmq_connection_parameters = create_rmq_connection_parameters(
    host = os.getenv('APP_RABBITMQ_HOST', 'localhost'),
    user = os.getenv('APP_RABBITMQ_USER', ''),
    password = os.getenv('APP_RABBITMQ_PASS', ''),
    virtual_host = os.getenv('APP_RABBITMQ_VHOST', '/'),
    heartbeat=30
)
rmq_event_map = RMQEventMap({})

################################################
# -- 🪠 Kafka setting
################################################
kafka_connection_parameters = create_kafka_connection_parameters(
    bootstrap_servers = os.getenv('APP_KAFKA_BOOTSTRAP_SERVERS', 'localhost:29092'),
    sasl_mechanism = os.getenv('APP_KAFKA_SASL_MECHANISM', 'PLAIN'),
    sasl_plain_username = os.getenv('APP_KAFKA_SASL_PLAIN_USERNAME', ''),
    sasl_plain_password = os.getenv('APP_KAFKA_SASL_PLAIN_PASSWORD', '')
)
kafka_event_map = KafkaEventMap({})

################################################
# -- 🪠 Kafka avro setting
################################################
kafka_avro_connection_parameters = create_kafka_avro_connection_parameters(
    bootstrap_servers = os.getenv('APP_KAFKA_BOOTSTRAP_SERVERS', 'localhost:29092'),
    schema_registry = os.getenv('APP_KAFKA_SCHEMA_REGISTRY', 'http://localhost:8035'),
    sasl_mechanism = os.getenv('APP_KAFKA_SASL_MECHANISM', 'PLAIN'),
    sasl_plain_username = os.getenv('APP_KAFKA_SASL_PLAIN_USERNAME', ''),
    sasl_plain_password = os.getenv('APP_KAFKA_SASL_PLAIN_PASSWORD', '')
)
kafka_avro_event_map = KafkaAvroEventMap({})

################################################
# -- 🚌 Message bus and RPC initialization
################################################
mb_type = os.getenv('APP_MESSAGE_BUS_TYPE', 'local')
rpc_type = os.getenv('APP_RPC_TYPE', 'local')
mb = create_message_bus(mb_type, rmq_connection_parameters, rmq_event_map, kafka_connection_parameters, kafka_event_map, kafka_avro_connection_parameters, kafka_avro_event_map)
rpc = create_rpc(rpc_type, rmq_connection_parameters, rmq_event_map)

################################################
# -- 🛢️ Database engine initialization
################################################
db_url = os.getenv('APP_SQLALCHEMY_DATABASE_URL', 'sqlite:///database.db')
engine = create_engine(db_url, echo=True)
role_repo = DBRoleRepo(engine=engine, create_all=True)
user_repo = DBUserRepo(engine=engine, create_all=True)

################################################
# -- 👤 User initialization
################################################
role_service = RoleService(role_repo)
guest_username = os.getenv('APP_GUEST_USERNAME', 'guest')
root_permission = os.getenv('APP_ROOT_PERMISSION', 'root')
user_service = DefaultUserService(user_repo, guest_username)
user_seeder_service = UserSeederService(user_service)
user_seeder_service.seed(UserData(
    username = os.getenv('APP_ROOT_USERNAME', 'root'),
    email = os.getenv('APP_ROOT_INITIAL_EMAIL', 'root@root.com'),
    phone_number = os.getenv('APP_ROOT_INITIAL_PHONE_NUMBER', '+621234567890'),
    password = os.getenv('APP_ROOT_INITIAL_PASSWORD', 'toor'),
    active = True,
    permissions = [root_permission],
    full_name = os.getenv('APP_ROOT_INITIAL_FULL_NAME', 'root')
))
token_service = JWTTokenService(
    user_service = user_service,
    access_token_secret_key = os.getenv('APP_ACCESS_TOKEN_SECRET_KEY', '123'),
    access_token_algorithm = os.getenv('APP_ACCESS_TOKEN_ALGORITHM', 'HS256'),
    access_token_expire_minutes = int(os.getenv('APP_ACCESS_TOKEN_EXPIRE_MINUTES', '30'))
)
access_token_url = os.getenv('APP_ACCESS_TOKEN_URL', '/api/v1/token/')
oauth2_scheme = OAuth2PasswordBearer(tokenUrl = access_token_url, auto_error = False)
auth_service = TokenOAuth2AuthService(user_service, role_service, token_service, oauth2_scheme, root_permission)

################################################
# -- 👓 User Interface initialization
################################################
menu_service: MenuService = create_menu_service(auth_service)
templates = Jinja2Templates(directory='templates')

################################################
# -- ⚛️ FastAPI initialization
################################################
app = FastAPI(title='ztplAppName')
app.add_middleware(
    CORSMiddleware,
    allow_origins = json.loads(os.getenv('APP_CORS_ALLOW_ORIGINS', '["*"]')),
    allow_origin_regex = os.getenv('APP_CORS_ALLOW_ORIGIN_REGEX', ''),
    allow_methods = json.loads(os.getenv('APP_CORS_ALLOW_METHODS', '["*"]')),
    allow_headers = json.loads(os.getenv('APP_CORS_ALLOW_HEADERS', '["*"]')),
    allow_credentials = os.getenv('APP_CORS_ALLOW_CREDENTIALS', 'false') == 'true',
    expose_headers = os.getenv('APP_CORS_EXPOSE_HEADERS', 'false') == 'true',
    max_age = int(os.getenv('APP_CORS_MAX_AGE', '600')),
)

################################################
# -- 📖 Register core handlers
################################################
enable_route_handler = os.getenv('APP_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_ui = os.getenv('APP_ENABLE_UI', '1') != '0'
enable_event_handler = os.getenv('APP_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler = os.getenv('APP_ENABLE_RPC_HANDLER', '1') != '0'
static_url = os.getenv('APP_STATIC_URL', '/static')
static_dir = get_abs_static_dir(os.getenv('APP_STATIC_DIR', ''))
handle_app_shutdown(app, mb, rpc)
register_readiness_handler(app, mb, rpc, error_threshold)
register_static_dir_route_handler(app, static_url, static_dir, static_route_name='static')
register_template_exception_handler(app, templates)
if enable_route_handler:
    register_auth_route_handler(app, mb, rpc, auth_service, menu_service, templates, enable_ui, access_token_url)
if enable_event_handler:
    register_auth_event_handler(mb)
if enable_rpc_handler:
    register_auth_rpc_handler(rpc, role_service, user_service, token_service)
