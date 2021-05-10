from helper import cli
from helper.project import HelmServiceProject

# USAGE
# python create_service_deployment.py

@cli
def create_service_deployment(service_name: str):
    dir_name = '.'
    if service_name == '':
        raise 'Service name should be given'
    helm_service_project = HelmServiceProject()
    helm_service_project.generate(dir_name, service_name)


if __name__ == '__main__':
    create_service_deployment()