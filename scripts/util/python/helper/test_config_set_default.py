from .config import YamlConfig


def get_yaml_config() -> YamlConfig:
    yaml_config = YamlConfig({})
    yaml_config.load('./test_resources/yaml_config.yaml')
    return yaml_config


def test_yaml_config_set_default_config_hello_name():
    yaml_config = get_yaml_config()
    yaml_config.set_default(['config', 'hello', 'name'], 'koga')
    assert yaml_config.get(['config', 'hello', 'name']) == 'koga'


def test_yaml_config_set_default_tasks_hello_description_text():
    yaml_config = get_yaml_config()
    yaml_config.set_default(['tasks', 'hello', 'description'], 'a description')
    yaml_config.get(['tasks', 'hello', 'description']) == 'show hello world'
   