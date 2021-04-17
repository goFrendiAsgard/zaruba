from .config import YamlConfig


def get_yaml_config() -> YamlConfig:
    yaml_config = YamlConfig({})
    yaml_config.load('./test_resources/yaml_config.yaml')
    return yaml_config


def test_yaml_config_set_config_hello_name():
    yaml_config = get_yaml_config()
    yaml_config.set(['config', 'hello', 'name'], 'koga')
    assert yaml_config.get(['config', 'hello', 'name']) == 'koga'


def test_yaml_config_set_tasks_hello_description_text():
    yaml_config = get_yaml_config()
    try:
        yaml_config.set(['tasks', 'hello', 'description', 'text'], 'a description')
        assert False, "Error Expected"
    except Exception as e:
        assert str(e) == "'str' object does not support item assignment"
   