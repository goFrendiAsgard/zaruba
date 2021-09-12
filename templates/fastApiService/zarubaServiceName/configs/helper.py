from helpers.transport import MessageBus, RPC, RMQMessageBus, RMQRPC, RMQEventMap, LocalMessageBus, LocalRPC
import os

def get_abs_static_dir(raw_static_dir: str) -> str:
    return os.path.abspath(raw_static_dir) if raw_static_dir != '' else ''

def create_message_bus(mb_type: str, rmq_host: str, rmq_user: str, rmq_pass: str, rmq_vhost: str, rmq_event_map: RMQEventMap) -> MessageBus:
    if mb_type == 'rmq':
        return RMQMessageBus(rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)
    return LocalMessageBus()

def create_rpc(mb_type: str, rmq_host: str, rmq_user: str, rmq_pass: str, rmq_vhost: str, rmq_event_map: RMQEventMap) -> RPC:
    if mb_type == 'rmq':
        return RMQRPC(rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)
    return LocalRPC()

