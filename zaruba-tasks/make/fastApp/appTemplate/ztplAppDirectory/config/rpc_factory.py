from helper.transport import RPC, RMQRPC, LocalRPC
from config.rmq import rmq_connection_parameters, rmq_event_map
from transport import AppRPC


def create_rpc(rpc_type: str) -> AppRPC:
    if rpc_type == 'rmq':
        return AppRPC(RMQRPC(rmq_connection_parameters, rmq_event_map))
    return AppRPC(LocalRPC())

