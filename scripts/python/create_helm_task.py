from helper import cli
from helper.project import HelmProject


@cli
def create_helm_task(template_location: str):
    dir_name = '.'
    helm_project = HelmProject(template_location)
    helm_project.generate(dir_name)


if __name__ == '__main__':
    create_helm_task()