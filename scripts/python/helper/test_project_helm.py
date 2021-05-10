from .project import MainProject, HelmProject
import shutil

def test_helm_project_generate():
    dir_name = './playground/helm_deployment_project'
    try:
        shutil.rmtree(dir_name)
    except OSError:
        pass
    main_project = MainProject()
    main_project.generate(dir_name)
    # generate service project
    helm_project = HelmProject()
    helm_project.generate(dir_name=dir_name)
    # reload
    generated_project = HelmProject()
    generated_project.load(dir_name=dir_name)
    # assert generated project
    assert generated_project.exist(['environments', 'default'])
    assert generated_project.exist(['environments', 'stable'])
    assert generated_project.exist(['repositories'])
    assert generated_project.exist(['releases'])
    # assert main project
    main_project = generated_project.main_project
    assert main_project.get(['tasks', 'helmApply', 'extend']) == 'core.helmApply'
    assert main_project.get(['tasks', 'helmApply', 'location']) == 'helm-deployments'
    assert main_project.get(['tasks', 'helmDestroy', 'extend']) == 'core.helmDestroy'
    assert main_project.get(['tasks', 'helmDestroy', 'location']) == 'helm-deployments'