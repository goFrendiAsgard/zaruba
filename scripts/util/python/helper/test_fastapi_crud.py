from  .codegen import FastApiService, FastApiModule, FastApiCrud

import shutil


def test_fastapi_crud_generate():
    dir_name = './playground/fast_api_project_with_crud'
    service_template_dir_name = '../../templates/fastApiService'
    module_template_dir_name = '../../templates/fastApiModule'
    crud_template_dir_name = '../../templates/fastApiCrud'
    service_name = 'servo'
    module_name = 'modulo'
    entity_name = 'ento'
    field_names = ['unue', 'dua']
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    service = FastApiService(service_name)
    service.load_from_template(service_template_dir_name)
    service.generate(dir_name)
    module = FastApiModule(service_name, module_name)
    module.load_from_template(module_template_dir_name)
    module.generate(dir_name)
    crud = FastApiCrud(service_name, module_name, entity_name, field_names)
    crud.load_from_template(crud_template_dir_name)
    crud.generate(dir_name)
    