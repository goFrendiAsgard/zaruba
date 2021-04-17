from .config import YamlConfig
import os


def get_yaml_config() -> YamlConfig:
    yaml_config = YamlConfig({})
    yaml_config.load('./test_resources/yaml_config.yaml')
    return yaml_config


def test_yaml_config_save():
    yaml_config = get_yaml_config()
    yaml_config.set(['tasks', 'hello', 'description'], 'Say hello')
    file_name = './playground/yaml_config/new_yaml_config.yaml'
    try:
        os.remove(file_name)
        os.removedirs('./playground/yaml_config')
    except OSError:
        pass
    yaml_config.save(file_name)
    new_yaml_config = YamlConfig({})
    new_yaml_config.load(file_name)
    assert new_yaml_config.get(['tasks', 'hello', 'description']) == 'Say hello'