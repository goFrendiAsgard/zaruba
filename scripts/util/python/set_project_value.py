from helper import cli
from helper.config import YamlConfig

@cli
def set_project_value(key, value, file_name='./default.values.yaml'):
    yaml_config = YamlConfig()
    yaml_config.load(file_name)
    yaml_config.set([key], value)
    yaml_config.save(file_name)


if __name__ == '__main__':
    set_project_value()

