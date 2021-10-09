
@rpc.handle('ztplAppRpcName')
def handle_zaruba_rpc_name(parameter: str) -> str:
    print('handle RPC call zarubaRPCName with parameter: {}'.format(parameter))
    return parameter
