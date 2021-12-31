
# -- ztplAppModuleName

if enable_route_handler:
    register_ztpl_app_module_name_route_handler(app, mb, rpc, auth_model)

if enable_event_handler:
    register_ztpl_app_module_name_event_handler(mb)

if enable_rpc_handler:
    register_ztpl_app_module_name_rpc_handler(rpc)
