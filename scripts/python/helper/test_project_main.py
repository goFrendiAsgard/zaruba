from dotenv import dotenv_values
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
    helm_template_location = '../templates/helmDeployments'
    dir_name = './playground/main_project_update_env'
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    main_project = MainProject()
    main_project.generate(dir_name)
    # prepare app_location
    service_name = 'app'
    app_location = os.path.join(dir_name, service_name)
    shutil.copytree('./test_resources/app', app_location)
    # generate service project
    service_project = ServiceProject()
    service_project.load_from_template('./test_resources/service.zaruba.yaml')
    service_project.generate(dir_name=dir_name, service_name=service_name, image_name='myImage', container_name='myContainer', location=app_location, start_command='node main.js', port_list=[], env_list=[], dependency_list=[], runner_version='')
    # generate helm project
    helm_project = HelmProject(helm_template_location)
    helm_project.generate(dir_name=dir_name)
    # generate helm service project
    helm_service_project = HelmServiceProject()
    helm_service_project.generate(dir_name, service_name)
    # update env file
    f_write = open(os.path.join(app_location, 'template.env'), 'a')
    f_write.write('\nFOO=BAR\n')
    f_write.close()
    # reload main project and update env
    main_project.load(dir_name)
    main_project.update_env(dir_name)
    # assert env
    service_project.save_env(dir_name, service_name)
    envs: Mapping[str, str] = dotenv_values(os.path.join(dir_name, 'template.env'))
    assert envs['APP_PORT'] == '3000'
    assert envs['APP_FOO'] == 'BAR'
    # assert service project
    service_project.load(dir_name, service_name)
    assert service_project.get(['envs', 'app', 'PORT', 'from']) == 'APP_PORT'
    assert service_project.get(['envs', 'app', 'PORT', 'default']) == '3000'
    assert service_project.get(['envs', 'app', 'FOO', 'from']) == 'APP_FOO'
    assert service_project.get(['envs', 'app', 'FOO', 'default']) == 'BAR'
    # assert helm service project
    helm_service_project.load(dir_name, service_name)
    assert helm_service_project.get(['app', 'container', 'env', 0, 'name']) == 'PORT'
    assert helm_service_project.get(['app', 'container', 'env', 0, 'value']) == '3000'
    assert helm_service_project.get(['app', 'container', 'env', 1, 'name']) == 'FOO'
    assert helm_service_project.get(['app', 'container', 'env', 1, 'value']) == 'BAR'

