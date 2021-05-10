from helper import cli
from helper.project import MainProject


@cli
def update_env():
    dir_name = '.'
    main_project = MainProject()
    main_project.load(dir_name)
    main_project.update_env(dir_name)


if __name__ == '__main__':
    update_env()