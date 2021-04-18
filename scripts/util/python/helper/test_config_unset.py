from .config import YamlConfig


def get_yaml_config() -> YamlConfig:
    yaml_config = YamlConfig()
    yaml_config.load('./test_resources/yaml_config.yaml')
    return yaml_config


def test_yaml_config_unset_tasks_hello_description():
    yaml_config = get_yaml_config()
    yaml_config.unset(['tasks', 'hello', 'description'])
    try:
        yaml_config.get(['tasks', 'hello', 'description'])
        assert False, "Error Expected"
    except Exception as e:
        assert str(e) == "`ordereddict([('start', ['bash', '-c', 'echo hello world'])])` has no key `description`"


def test_yaml_config_unset_env_invalid():
    yaml_config = get_yaml_config()
    yaml_config.set(['env', 'key', 'default'], 'value')
    yaml_config.unset(['config', 'invalid', 'default'])


def test_yaml_config_unset_env():
    yaml_config = get_yaml_config()
    yaml_config.unset(['env'])