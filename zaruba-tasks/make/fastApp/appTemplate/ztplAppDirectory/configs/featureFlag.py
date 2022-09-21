import os

enable_auth_module = os.getenv('APP_ENABLE_AUTH_MODULE', '1') != '0'
enable_route_handler = os.getenv('APP_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_ui = os.getenv('APP_ENABLE_UI', '1') != '0'
enable_api = os.getenv('APP_ENABLE_API', '1') != '0'
enable_event_handler = os.getenv('APP_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler = os.getenv('APP_ENABLE_RPC_HANDLER', '1') != '0'
seed_root_user = os.getenv('APP_SEED_ROOT_USER', '1') != '0'