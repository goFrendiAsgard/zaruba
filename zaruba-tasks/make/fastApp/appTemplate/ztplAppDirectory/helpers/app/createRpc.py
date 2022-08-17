from helpers.transport import RPC, RMQRPC, RMQEventMap, LocalRPC
import pika

def create_rpc(mb_type: str, rmq_connection_parameters: pika.ConnectionParameters, rmq_event_map: RMQEventMap) -> RPC:
    if mb_type == 'rmq':
        return RMQRPC(rmq_connection_parameters, rmq_event_map)
    return LocalRPC()

