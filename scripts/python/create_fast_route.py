from helper import cli
from helper.codegen import FastApiModule

@cli
def create_fast_route(template_location: str, service_name: str, module_name: str, http_method: str, url: str):
    dir_name = '.'
    module = FastApiModule(service_name, module_name)
    module.load_from_template(template_location)
    module.add_route(dir_name, http_method, url)


if __name__ == '__main__':
    create_fast_route()