from helper import cli
from helper.project import HelmProject


@cli
def create_helm_task():
    dir_name = '.'
    helm_project = HelmProject()
    helm_project.generate(dir_name)


if __name__ == '__main__':
    create_helm_task()