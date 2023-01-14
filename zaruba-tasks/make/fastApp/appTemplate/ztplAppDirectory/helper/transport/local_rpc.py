from typing import Any, Callable, Mapping
from helper.transport.rpc import RPC

import logging


class LocalRPC(RPC):

    def __init__(self):
        self.rpc_handler: Mapping[str, Callable[..., Any]] = {}
        self.error_count = 0

    def get_error_count(self) -> int:
        return self.error_count

    def is_failing(self) -> bool:
        return False

    def shutdown(self):
        pass

    def handle(self, rpc_name: str) -> Callable[..., Any]:
        def register_rpc_handler(rpc_handler: Callable[..., Any]):
            self.rpc_handler[rpc_name] = rpc_handler
        return register_rpc_handler

    def call(self, rpc_name: str, *args: Any) -> Any:
        if rpc_name not in self.rpc_handler:
            self.error_count += 1
            raise Exception(
                'RPC handler for "{}" is not found'.format(rpc_name)
            )
        logging.info(
            'Call RPC {} with arguments: {}'.format(rpc_name, args)
        )
        try:
            logging.info(
                'Handle RPC {} with arguments: {}'.format(rpc_name, args)
            )
            result = self.rpc_handler[rpc_name](*args)
            logging.info(
                'Reply RPC {} with arguments: {}, result: {}'.format(
                    rpc_name, args, result
                )
            )
            return result
        except Exception as exception:
            logging.error(
                'Error calling RPC {} with arguments: {}'.format(
                    rpc_name, args
                ),
                exc_info=True
            )
            raise exception
