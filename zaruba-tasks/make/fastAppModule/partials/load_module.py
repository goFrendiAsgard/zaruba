
################################################
# -- 🧩 ZtplAppModuleName module
################################################
enable_ztpl_app_module_name_module = os.getenv('APP_ENABLE_ZTPL_APP_MODULE_NAME_MODULE', '1') != '0'
# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
if enable_ztpl_app_module_name_module:
    # API route
    if enable_route_handler and enable_api:
        register_ztpl_app_module_name_api_route(app, mb, rpc, auth_service)
    # UI route
    if enable_route_handler and enable_ui:
        register_ztpl_app_module_name_ui_route(app, mb, rpc, menu_service, page_template)
    # handle event
    if enable_event_handler:
        register_ztpl_app_module_name_event_handler(mb, rpc, auth_service)
    # serve RPC
    if enable_rpc_handler:
        # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
        register_ztpl_app_module_name_rpc_handler(mb, rpc, auth_service)
