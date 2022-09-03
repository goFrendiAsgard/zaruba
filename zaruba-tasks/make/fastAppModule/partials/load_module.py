
################################################
# -- 🧩 ZtplAppModuleName module
################################################
enable_ztpl_app_module_name_module = os.getenv('APP_ENABLE_ZTPL_APP_MODULE_NAME_MODULE', '1') != '0'
if enable_ztpl_app_module_name_module:
    if enable_route_handler:
        register_ztpl_app_module_name_route_handler(app, mb, rpc, auth_service, menu_service, templates, enable_ui, enable_api)
    if enable_event_handler:
        register_ztpl_app_module_name_event_handler(mb)
    if enable_rpc_handler:
        register_ztpl_app_module_name_rpc_handler(rpc)
