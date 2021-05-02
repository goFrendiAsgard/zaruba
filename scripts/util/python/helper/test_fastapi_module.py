from  .codegen import FastApiService, FastApiModule

import shutil


def test_fastapi_module_generate():
    dir_name = './playground/fast_api_project_with_module'
    service_template_dir_name = '../../templates/fastApiService'
    module_template_dir_name = '../../templates/fastApiModule'
    service_name = 'servo'
    module_name = 'modulo'
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
    # Load generated module
    generated_module = FastApiModule(service_name, module_name)
    generated_module.load(dir_name)
    controller_content = generated_module.get_content('servo/modulo/controller.py')
    assert controller_content.find('Handle events for modulo') > -1
    assert controller_content.find('Handle routes for modulo') > -1
    # load generated service
    generated_service = FastApiService(service_name)
    generated_service.load(dir_name)
    main_content = generated_service.get_content('servo/main.py')
    assert main_content.find('SERVO_') > -1
    assert main_content.find('from modulo.controller import') > -1


def test_fastapi_module_handle_route():
    dir_name = './playground/fast_api_project_with_module_and_route'
    service_template_dir_name = '../../templates/fastApiService'
    module_template_dir_name = '../../templates/fastApiModule'
    service_name = 'servo'
    module_name = 'modulo'
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
    module.add_route(dir_name, http_method='get', url='/helloWorld')
    # Load generated module
    generated_module = FastApiModule(service_name, module_name)
    generated_module.load(dir_name)
    controller_content = generated_module.get_content('servo/modulo/controller.py')
    assert controller_content.find("@self.app.get('/helloWorld')") > -1
    assert controller_content.find('def get_hello_world') > -1


def test_fastapi_module_handle_event():
    dir_name = './playground/fast_api_project_with_module_and_event'
    service_template_dir_name = '../../templates/fastApiService'
    module_template_dir_name = '../../templates/fastApiModule'
    service_name = 'servo'
    module_name = 'modulo'
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
    module.add_event_handler(dir_name, event_name='bigSale')
    # Load generated module
    generated_module = FastApiModule(service_name, module_name)
    generated_module.load(dir_name)
    controller_content = generated_module.get_content('servo/modulo/controller.py')
    assert controller_content.find("@self.mb.handle_event('bigSale')") > -1
    assert controller_content.find('def handle_event_big_sale') > -1


def test_fastapi_module_handle_rpc():
    dir_name = './playground/fast_api_project_with_module_and_rpc'
    service_template_dir_name = '../../templates/fastApiService'
    module_template_dir_name = '../../templates/fastApiModule'
    service_name = 'servo'
    module_name = 'modulo'
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
    module.add_rpc_handler(dir_name, event_name='buyStock')
    # Load generated module
    generated_module = FastApiModule(service_name, module_name)
    generated_module.load(dir_name)
    controller_content = generated_module.get_content('servo/modulo/controller.py')
    assert controller_content.find("@self.mb.handle_rpc('buyStock')") > -1
    assert controller_content.find('def handle_rpc_buy_stock') > -1
