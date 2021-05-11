from dotenv import dotenv_values
from .project import MainProject, ServiceProject, HelmProject, HelmServiceProject

import os
import shutil

def test_helm_service_project_generate():
    helm_template_location = '../templates/helmDeployments'
    dir_name = './playground/helm_deployment_project'
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    main_project = MainProject()
    main_project.generate(dir_name)
    service_name = 'myService'
    # generate service project
    service_project = ServiceProject()
    service_project.load_from_template('./test_resources/service.zaruba.yaml')
    service_project.generate(dir_name=dir_name, service_name=service_name, image_name='myImage', container_name='myContainer', location='./test_resources/app', start_command='node main.js', port_list=[], env_list=[], runner_version='')
    # generate helm project
    helm_project = HelmProject(helm_template_location)
    helm_project.generate(dir_name=dir_name)
    # generate helm service project
    helm_service_project = HelmServiceProject()
    helm_service_project.generate(dir_name, service_name)
    # assert generated project
    generated_project = HelmServiceProject()
    generated_project.load(dir_name, service_name)
    assert generated_project.get(['app', 'ports', 0, 'containerPort']) == 3000
    assert generated_project.get(['app', 'ports', 0, 'servicePort']) == 3000
    assert generated_project.get(['app', 'name']) == 'my-service'
    assert generated_project.get(['app', 'container', 'image']) == 'my-image'
    assert generated_project.get(['app', 'container', 'env', 0, 'name']) == 'PORT'
    assert generated_project.get(['app', 'container', 'env', 0, 'value']) == '3000'
    # assert helm project
    helm_project = generated_project.helm_project
    assert helm_project.get(['releases', 3, 'name']) == 'my-service'
    assert helm_project.get(['releases', 3, 'chart']) == './charts/app'
    assert helm_project.get(['releases', 3, 'values', 0]) == './values/my-service.yaml.gotmpl'