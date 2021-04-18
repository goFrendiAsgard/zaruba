from .project import MainProject, ServiceProject, HelmProject, HelmServiceProject

import os
import shutil

def test_main_project_generate():
    dir_name = './playground/main_project'
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    main_project = MainProject()
    main_project.generate(dir_name)
    # load generated project
    generated_main_project = MainProject()
    generated_main_project.load(dir_name)
    assert generated_main_project.get(['includes', 0]) == '${ZARUBA_HOME}/scripts/core.zaruba.yaml'
    assert generated_main_project.exist(['tasks', 'run'])
    assert generated_main_project.exist(['tasks', 'runContainer'])
    assert generated_main_project.exist(['tasks', 'stopContainer'])
    assert generated_main_project.exist(['tasks', 'removeContainer'])
    assert generated_main_project.exist(['tasks', 'buildImage'])
    assert generated_main_project.exist(['tasks', 'pushImage'])


def test_main_project_update_env():
    dir_name = './playground/main_project_update_env'
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    main_project = MainProject()
    main_project.generate(dir_name)
    # prepare app_location
    app_location = os.path.join(dir_name, 'app')
    shutil.copytree('./test_resources/app', app_location)
    # generate service project
    service_project = ServiceProject()
    service_project.load_from_template('./test_resources/service.zaruba.yaml')
    service_project.generate(dir_name=dir_name, service_name='app', image_name='myImage', container_name='myContainer', location=app_location, start_command='node main.js', ports=[])
    # generate helm project
    helm_project = HelmProject()
    helm_project.generate(dir_name=dir_name)
    # generate helm service project
    helm_service_project = HelmServiceProject()
    helm_service_project.generate(dir_name, 'app')
    # update env file
    f_write = open(os.path.join(app_location, 'template.env'), 'a')
    f_write.write('\nFOO=BAR\n')
    f_write.close()
    # reload main project and update env
    main_project.load(dir_name)
    main_project.update_env(dir_name)
