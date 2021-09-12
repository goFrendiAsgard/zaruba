
if enable_http_handler:
    zaruba_module_name_http_controller(app, mb, rpc)

if enable_event_handler:
    zaruba_module_name_event_controller(mb)

if enable_rpc_handler:
    zaruba_module_name_rpc_controller(rpc)
