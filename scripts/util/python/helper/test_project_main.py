from .project import MainProject
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
