from helper import cli
from helper.codegen import FastApiModule

@cli
def create_fast_service(template_location: str, service_name: str, module_name: str):
    dir_name = '.'
    module = FastApiModule(service_name, module_name)
    module.load_from_template(template_location)
    module.generate(dir_name)


if __name__ == '__main__':
    create_fast_service()