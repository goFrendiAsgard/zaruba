from  .codegen import FastApiService, FastApiModule, FastApiCrud

import shutil


def test_fastapi_crud_generate():
    dir_name = './playground/fast_api_project_with_crud'
    service_template_dir_name = '../templates/fastApiService'
    module_template_dir_name = '../templates/fastApiModule'
    crud_template_dir_name = '../templates/fastApiCrud'
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
    # load generated crud
    generated_crud = FastApiCrud(service_name, module_name, entity_name, field_names)
    generated_crud.load(dir_name)
    db_repo_content = generated_crud.get_content('servo/repos/dbEnto.py')
    assert db_repo_content.find('from schemas.ento import Ento, EntoData') > -1
    assert db_repo_content.find('class DBEntoRepo(') > -1
    assert db_repo_content.find('class DBEntoEntity(') > -1
    assert db_repo_content.find('unue = Column(String(20), index=True)') > -1
    assert db_repo_content.find('unue=ento_data.unue,') > -1
    assert db_repo_content.find('db_entity.unue = ento_data.unue') > -1
    db_content = generated_crud.get_content('servo/repos/ento.py')
    assert db_content.find('from schemas.ento import Ento, EntoData') > -1
    assert db_content.find('EntoRepo') > -1
    schema_content = generated_crud.get_content('servo/schemas/ento.py')
    assert schema_content.find('class EntoData') > -1
    assert schema_content.find('class Ento(EntoData)') > -1
    assert schema_content.find('unue: str') > -1
    handle_event_content = generated_crud.get_content('servo/modulo/handleEntoEvent.py')
    assert handle_event_content.find('from schemas.ento import Ento, EntoData') > -1
    assert handle_event_content.find('find_ento') > -1
    assert handle_event_content.find('find_ento_by_id') > -1
    assert handle_event_content.find('insert_ento') > -1
    assert handle_event_content.find('update_ento') > -1
    assert handle_event_content.find('delete_ento') > -1
    handle_route_content = generated_crud.get_content('servo/modulo/handleEntoRoute.py')
    assert handle_route_content.find('from schemas.ento import Ento, EntoData') > -1
    assert handle_route_content.find('find_ento') > -1
    assert handle_route_content.find('find_ento_by_id') > -1
    assert handle_route_content.find('insert_ento') > -1
    assert handle_route_content.find('update_ento') > -1
    assert handle_route_content.find('delete_ento') > -1
    # load generated module
    generated_module = FastApiModule(service_name, module_name)
    generated_module.load(dir_name)
    controller_content = generated_module.get_content('servo/modulo/controller.py')
    assert controller_content.find('from repos.ento import EntoRepo') > -1
    assert controller_content.find('from modulo.handleEntoRoute import handle_route as handle_ento_route') > -1
    assert controller_content.find('from modulo.handleEntoEvent import handle_event as handle_ento_event') > -1
    assert controller_content.find('handle_ento_event(self.mb, self.ento_repo') > -1
    assert controller_content.find('handle_ento_route(self.app, self.mb)') > -1
    # load service
    generated_service = FastApiService(service_name)
    generated_service.load(dir_name)
    main_content = generated_service.get_content('servo/main.py')
    assert main_content.find('from repos.dbEnto import DBEntoRepo') > -1
    assert main_content.find('ento_repo = DBEntoRepo(engine=engine, create_all=True)') > -1
    assert main_content.find('modulo_controller = ModuloController(app=app, mb=mb, enable_route=enable_route, enable_event=enable_event, ento_repo=ento_repo)') > -1
