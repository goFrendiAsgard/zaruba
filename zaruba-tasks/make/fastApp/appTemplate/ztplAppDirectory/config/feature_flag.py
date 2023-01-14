from helper.config import get_boolean_env

enable_auth_module: bool = get_boolean_env(
    'APP_ENABLE_AUTH_MODULE', True
)
enable_cms_module: bool = get_boolean_env(
    'APP_ENABLE_CMS_MODULE', True
)
enable_log_module: bool = get_boolean_env(
    'APP_ENABLE_LOG_MODULE', True
)
enable_route_handler: bool = get_boolean_env(
    'APP_ENABLE_ROUTE_HANDLER', True
)
enable_ui: bool = get_boolean_env(
    'APP_ENABLE_UI', True
)
enable_api: bool = get_boolean_env(
    'APP_ENABLE_API', True
)
enable_error_page: bool = get_boolean_env(
    'APP_ENABLE_ERROR_PAGE', True
)
enable_event_handler: bool = get_boolean_env(
    'APP_ENABLE_EVENT_HANDLER', True
)
enable_rpc_handler: bool = get_boolean_env(
    'APP_ENABLE_RPC_HANDLER', True
)
