from helper import cli
from helper.codegen import FastApiService

@cli
def create_fast_service(template_location: str, service_name: str):
    dir_name = '.'
    service = FastApiService(service_name)
    service.load_from_template(template_location)
    service.generate(dir_name)


if __name__ == '__main__':
    create_fast_service()