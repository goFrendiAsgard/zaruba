@self.mb.handle_rpc('zarubaEventName')
def handle_rpc_zaruba_event_name(parameter: str) -> str:
    print('handle RPC call zarubaEventName with parameter: {}'.format(parameter))
    return parameter