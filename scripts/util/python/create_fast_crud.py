from typing import List
from helper import cli
from helper.codegen import FastApiCrud

@cli
def create_fast_crud(template_location: str, service_name: str, module_name: str, entity_name: str, field_names: str):
    field_name_list = field_names.split(',') if field_names != '' else []
    dir_name = '.'
    crud = FastApiCrud(service_name, module_name, entity_name, field_names=field_name_list)
    crud.load_from_template(template_location)
    crud.generate(dir_name)


if __name__ == '__main__':
    create_fast_crud()