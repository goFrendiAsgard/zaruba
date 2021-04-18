from dotenv import dotenv_values
from .project import MainProject, ServiceProject

import os
import shutil

def test_service_project_generate():
    dir_name = './playground/service_project'
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
    service_project.generate(dir_name=dir_name, service_name=service_name, image_name='myImage', container_name='myContainer', location='./test_resources/app', start_command='node main.js', ports=[])
    # assert env
    service_project.save_env(dir_name, service_name)
    envs: Mapping[str, str] = dotenv_values(os.path.join(dir_name, 'template.env'))
    assert envs['MY_SERVICE_PORT'] == '3000'
    # reload
    generated_project = ServiceProject()
    generated_project.load(dir_name=dir_name, service_name=service_name)
    # runMyService
    assert generated_project.get(['tasks', 'runMyService', 'extend']) == 'core.startService'
    assert generated_project.get(['tasks', 'runMyService', 'location']) == '../../../test_resources/app'
    assert generated_project.get(['tasks', 'runMyService', 'configRef']) == 'myService'
    assert generated_project.get(['tasks', 'runMyService', 'envRef']) == 'myService'
    assert generated_project.get(['tasks', 'runMyService', 'lconfRef']) == 'myService'
    # runMyServiceContainer
    assert generated_project.get(['tasks', 'runMyServiceContainer', 'extend']) == 'core.startDockerContainer'
    assert generated_project.get(['tasks', 'runMyServiceContainer', 'configRef']) == 'myServiceContainer'
    assert generated_project.get(['tasks', 'runMyServiceContainer', 'lconfigRef']) == 'myService'
    assert generated_project.get(['tasks', 'runMyServiceContainer', 'envRef']) == 'myService'
    assert generated_project.get(['tasks', 'runMyServiceContainer', 'dependencies', 0]) == 'buildMyServiceImage'
    # stopMyServiceContainer
    assert generated_project.get(['tasks', 'stopMyServiceContainer', 'extend']) == 'core.stopDockerContainer'
    assert generated_project.get(['tasks', 'stopMyServiceContainer', 'configRef']) == 'myServiceContainer'
    # removeMyServiceContainer
    assert generated_project.get(['tasks', 'removeMyServiceContainer', 'extend']) == 'core.removeDockerContainer'
    assert generated_project.get(['tasks', 'removeMyServiceContainer', 'configRef']) == 'myServiceContainer'
    # buildMyServiceContainer
    assert generated_project.get(['tasks', 'buildMyServiceImage', 'extend']) == 'core.buildDockerImage'
    assert generated_project.get(['tasks', 'buildMyServiceImage', 'timeout']) == '1h'
    assert generated_project.get(['tasks', 'buildMyServiceImage', 'configRef']) == 'myServiceContainer'
    # pushMyServiceImage
    assert generated_project.get(['tasks', 'pushMyServiceImage', 'extend']) == 'core.pushDockerImage'
    assert generated_project.get(['tasks', 'pushMyServiceImage', 'timeout']) == '1h'
    assert generated_project.get(['tasks', 'pushMyServiceImage', 'configRef']) == 'myServiceContainer'
    assert generated_project.get(['tasks', 'pushMyServiceImage', 'dependencies', 0]) == 'buildMyServiceImage'
    # config
    assert generated_project.get(['configs', 'myService', 'start']) == 'node main.js'
    assert generated_project.get(['configs', 'myServiceContainer', 'containerName']) == 'myContainer'
    assert generated_project.get(['configs', 'myServiceContainer', 'imageName']) == 'myimage'
    # lconfig
    assert generated_project.get(['lconfigs', 'myService', 'ports', 0]) == '{{ .GetEnv "PORT" }}'
    # envs
    assert generated_project.get(['envs', 'myService', 'PORT', 'from']) == 'MY_SERVICE_PORT'
    assert generated_project.get(['envs', 'myService', 'PORT', 'default']) == '3000'
    # assert main project
    main_project = generated_project.main_project
    assert len(main_project.get(['includes'])) == 2
    assert main_project.get(['includes', 0]) == '${ZARUBA_HOME}/scripts/core.zaruba.yaml'
    assert main_project.get(['includes', 1]) == 'zaruba-tasks/myService.zaruba.yaml'
    assert main_project.get(['tasks', 'run', 'dependencies', 0]) == 'runMyService'
    assert main_project.get(['tasks', 'runContainer', 'dependencies', 0]) == 'runMyServiceContainer'
    assert main_project.get(['tasks', 'stopContainer', 'dependencies', 0]) == 'stopMyServiceContainer'
    assert main_project.get(['tasks', 'removeContainer', 'dependencies', 0]) == 'removeMyServiceContainer'
    assert main_project.get(['tasks', 'buildImage', 'dependencies', 0]) == 'buildMyServiceImage'
    assert main_project.get(['tasks', 'pushImage', 'dependencies', 0]) == 'pushMyServiceImage'