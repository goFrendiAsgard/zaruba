from .project import MainProject, DockerProject
import shutil

def test_docker_project_generate():
    dir_name = './playground/docker_project'
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    main_project = MainProject()
    main_project.generate(dir_name)
    # generate service project
    docker_project = DockerProject()
    docker_project.load_from_template('./test_resources/docker.zaruba.yaml')
    docker_project.generate(dir_name=dir_name, service_name='myService', image_name='myImage', container_name='myContainer')
    # reload
    generated_project = DockerProject()
    generated_project.load(dir_name=dir_name, service_name='myService')
    # runMyService
    assert generated_project.get(['tasks', 'runMyService', 'extend']) == 'core.startDockerContainer'
    assert generated_project.get(['tasks', 'runMyService', 'configRef']) == 'myService'
    assert generated_project.get(['tasks', 'runMyService', 'envRef']) == 'myService'
    assert generated_project.get(['tasks', 'runMyService', 'lconfigRef']) == 'myService'
    # stopMyServiceContainer
    assert generated_project.get(['tasks', 'stopMyServiceContainer', 'extend']) == 'core.stopDockerContainer'
    assert generated_project.get(['tasks', 'stopMyServiceContainer', 'configRef']) == 'myService'
    # removeMyServiceContainer
    assert generated_project.get(['tasks', 'removeMyServiceContainer', 'extend']) == 'core.removeDockerContainer'
    assert generated_project.get(['tasks', 'removeMyServiceContainer', 'configRef']) == 'myService'
    # config
    assert generated_project.get(['configs', 'myService', 'port::3306']) == 3306
    assert generated_project.get(['configs', 'myService', 'useImagePrefix']) == False
    assert generated_project.get(['configs', 'myService', 'containerName']) == 'myContainer'
    assert generated_project.get(['configs', 'myService', 'imageName']) == 'my-image'
    # assert main project
    main_project = generated_project.main_project
    assert len(main_project.get(['includes'])) == 2
    assert main_project.get(['includes', 0]) == '${ZARUBA_HOME}/scripts/core.zaruba.yaml'
    assert main_project.get(['includes', 1]) == 'zaruba-tasks/myService.zaruba.yaml'
    assert main_project.get(['tasks', 'run', 'dependencies', 0]) == 'runMyService'
    assert main_project.get(['tasks', 'stopContainer', 'dependencies', 0]) == 'stopMyServiceContainer'
    assert main_project.get(['tasks', 'removeContainer', 'dependencies', 0]) == 'removeMyServiceContainer'