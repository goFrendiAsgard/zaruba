from  .codegen import FastApiService

import shutil


def test_fastapi_service_generate():
    dir_name = './playground/fast_api_project'
    template_dir_name = '../templates/fastApiService'
    service_name = 'servo'
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    service = FastApiService(service_name)
    service.load_from_template(template_dir_name)
    service.generate(dir_name)
    # load generated service
    generated_service = FastApiService(service_name)
    generated_service.load(dir_name)
    main_content = generated_service.get_content('servo/main.py')
    assert main_content.find('SERVO_RABBITMQ_HOST') > -1
