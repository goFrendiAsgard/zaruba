
@rpc.handle('ztplAppRpcName')
def ztpl_app_rpc_name(parameter: str) -> str:
    print('handle RPC call ztplAppRpcName with parameter: {}'.format(parameter))
    return parameter
