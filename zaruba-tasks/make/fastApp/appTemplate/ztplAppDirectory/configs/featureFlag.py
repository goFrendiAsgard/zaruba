import os

enable_auth_module: bool = os.getenv('APP_ENABLE_AUTH_MODULE', '1') != '0'
enable_route_handler: bool = os.getenv('APP_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_ui: bool = os.getenv('APP_ENABLE_UI', '1') != '0'
enable_api: bool = os.getenv('APP_ENABLE_API', '1') != '0'
enable_error_page: bool = os.getenv('APP_ENABLE_ERROR_PAGE', '1') != '0'
enable_event_handler: bool = os.getenv('APP_ENABLE_EVENT_HANDLER', '1') != '0'
enable_rpc_handler: bool = os.getenv('APP_ENABLE_RPC_HANDLER', '1') != '0'