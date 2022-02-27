# -- 📖 Common import
from fastapi import FastAPI
from fastapi.security import OAuth2PasswordBearer
from sqlalchemy import create_engine
from helpers.transport import RMQEventMap, KafkaEventMap, KafkaAvroEventMap, create_kafka_connection_parameters, create_kafka_avro_connection_parameters, create_rmq_connection_parameters
from helpers.app import get_abs_static_dir, create_message_bus, create_rpc, handle_app_shutdown, register_static_dir_route_handler, register_readiness_handler
from repos.dbUser import DBUserRepo
from repos.dbRole import DBRoleRepo
from auth import register_auth_route_handler, register_auth_event_handler, register_auth_rpc_handler, TokenOAuth2AuthModel, JWTTokenModel, DefaultUserModel, UserSeederModel, RoleModel
from schemas.user import UserData

import os

error_threshold = int(os.getenv('APP_ERROR_THRESHOLD', '10'))

# -- 🐇 Rabbitmq setting
rmq_connection_parameters = create_rmq_connection_parameters(
    host = os.getenv('APP_RABBITMQ_HOST', 'localhost'),
    user = os.getenv('APP_RABBITMQ_USER', 'root'),
    password = os.getenv('APP_RABBITMQ_PASS', ''),
    virtual_host = os.getenv('APP_RABBITMQ_VHOST', '/'),
    heartbeat=30
)
rmq_event_map = RMQEventMap({})

# -- 🪠 Kafka setting
kafka_connection_parameters = create_kafka_connection_parameters(
    bootstrap_servers = os.getenv('APP_KAFKA_BOOTSTRAP_SERVERS', 'localhost:29092'),
    sasl_mechanism = os.getenv('APP_KAFKA_SASL_MECHANISM', 'PLAIN'),
    sasl_plain_username = os.getenv('APP_KAFKA_SASL_PLAIN_USERNAME', 'root'),
    sasl_plain_password = os.getenv('APP_KAFKA_SASL_PLAIN_PASSWORD', '')
)
kafka_event_map = KafkaEventMap({})

# -- 🪠 Kafka avro setting
kafka_avro_connection_parameters = create_kafka_avro_connection_parameters(
    bootstrap_servers = os.getenv('APP_KAFKA_BOOTSTRAP_SERVERS', 'localhost:29092'),
    schema_registry = os.getenv('APP_KAFKA_SCHEMA_REGISTRY', 'http://localhost:8035'),
    sasl_mechanism = os.getenv('APP_KAFKA_SASL_MECHANISM', 'PLAIN'),
    sasl_plain_username = os.getenv('APP_KAFKA_SASL_PLAIN_USERNAME', 'root'),
    sasl_plain_password = os.getenv('APP_KAFKA_SASL_PLAIN_PASSWORD', '')
)
kafka_avro_event_map = KafkaAvroEventMap({})

# -- 🚌 Message bus and RPC initialization
mb_type = os.getenv('APP_MESSAGE_BUS_TYPE', 'local')
rpc_type = os.getenv('APP_RPC_TYPE', 'local')
mb = create_message_bus(mb_type, rmq_connection_parameters, rmq_event_map, kafka_connection_parameters, kafka_event_map, kafka_avro_connection_parameters, kafka_avro_event_map)
rpc = create_rpc(rpc_type, rmq_connection_parameters, rmq_event_map)

# -- 🛢️ Database engine initialization
db_url = os.getenv('APP_SQLALCHEMY_DATABASE_URL', 'sqlite://')
engine = create_engine(db_url, echo=True)
role_repo = DBRoleRepo(engine=engine, create_all=True)
user_repo = DBUserRepo(engine=engine, create_all=True)

# -- 👤 User initialization
role_model = RoleModel(role_repo)
guest_username = os.getenv('APP_GUEST_USERNAME', 'guest')
root_permission = os.getenv('APP_ROOT_PERMISSION', 'root')
user_model = DefaultUserModel(user_repo, guest_username)
user_seeder_model = UserSeederModel(user_model)
user_seeder_model.seed(UserData(
    username = os.getenv('APP_ROOT_USERNAME', 'root'),
    email = os.getenv('APP_ROOT_INITIAL_EMAIL', 'root@root.com'),
    phone_number = os.getenv('APP_ROOT_INITIAL_PHONE_NUMBER', '+621234567890'),
    password = os.getenv('APP_ROOT_INITIAL_PASSWORD', 'toor'),
    active = True,
    permissions = [root_permission],
    full_name = os.getenv('APP_ROOT_INITIAL_FULL_NAME', 'root')
))
token_model = JWTTokenModel(
    user_model = user_model,
    access_token_secret_key = os.getenv('APP_ACCESS_TOKEN_SECRET_KEY', '123'),
    access_token_algorithm = os.getenv('APP_ACCESS_TOKEN_ALGORITHM', 'HS256'),
    access_token_expire_minutes = int(os.getenv('APP_ACCESS_TOKEN_EXPIRE_MINUTES', '30'))
)
access_token_url = os.getenv('APP_ACCESS_TOKEN_URL', '/token/')
oauth2_scheme = OAuth2PasswordBearer(tokenUrl = access_token_url, auto_error = False)
auth_model = TokenOAuth2AuthModel(user_model, role_model, token_model, oauth2_scheme, root_permission)

# -- ⚡FastAPI initialization
app = FastAPI(title='ztplAppName')

# -- 📖 Register core handlers
enable_route_handler = os.getenv('APP_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_event_handler = os.getenv('APP_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler = os.getenv('APP_ENABLE_RPC_HANDLER', '1') != '0'
static_url = os.getenv('APP_STATIC_URL', '/static')
static_dir = get_abs_static_dir(os.getenv('APP_STATIC_DIR', ''))
handle_app_shutdown(app, mb, rpc)
register_readiness_handler(app, mb, rpc, error_threshold)
register_static_dir_route_handler(app, static_url, static_dir, static_route_name='static')
if enable_route_handler:
    register_auth_route_handler(app, mb, rpc, access_token_url, auth_model)
if enable_event_handler:
    register_auth_event_handler(mb)
if enable_rpc_handler:
    register_auth_rpc_handler(rpc, role_model, user_model, token_model)
