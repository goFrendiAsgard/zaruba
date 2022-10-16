from helpers.transport import RPC, RMQRPC, LocalRPC
from configs.rmq import rmq_connection_parameters, rmq_event_map


def create_rpc(rpc_type: str) -> RPC:
    if rpc_type == 'rmq':
        return RMQRPC(rmq_connection_parameters, rmq_event_map)
    return LocalRPC()

