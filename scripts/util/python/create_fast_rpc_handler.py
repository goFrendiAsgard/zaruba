from helper import cli
from helper.codegen import FastApiModule

@cli
def create_fast_rpc_handler(template_location: str, service_name: str, module_name: str, event_name: str):
    dir_name = '.'
    module = FastApiModule(service_name, module_name)
    module.load_from_template(template_location)
    module.add_rpc_handler(dir_name, event_name)


if __name__ == '__main__':
    create_fast_rpc_handler()