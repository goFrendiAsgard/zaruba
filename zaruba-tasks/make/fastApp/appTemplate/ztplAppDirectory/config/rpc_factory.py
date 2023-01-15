from helper.transport import RMQRPC, LocalRPC, RPC
from config.rmq import rmq_connection_parameters, rmq_event_map
from transport import AppRPC


def create_app_rpc(rpc_type: str) -> AppRPC:
    '''
    Return a new AppRPC based on specified rpc_type.

    Keyword arguments:
    - rpc_type -- RPC type (e.g., rmq, local)
    '''
    rpc = create_rpc(rpc_type)
    return AppRPC(rpc)


def create_rpc(rpc_type: str) -> RPC:
    '''
    Return a new RPC based on specified rpc_type.

    Keyword arguments:
    - rpc_type -- RPC type (e.g., rmq, local)
    '''
    if rpc_type == 'rmq':
        return RMQRPC(rmq_connection_parameters, rmq_event_map)
    return LocalRPC()
