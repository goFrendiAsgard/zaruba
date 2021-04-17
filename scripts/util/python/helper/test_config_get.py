from .config import YamlConfig


def get_yaml_config() -> YamlConfig:
    yaml_config = YamlConfig({})
    yaml_config.load('./test_resources/yaml_config.yaml')
    return yaml_config


def test_yaml_config_get_tasks_hello_description():
    yaml_config = get_yaml_config()
    assert yaml_config.get(['tasks', 'hello', 'description']) == 'say hello world'


def test_yaml_config_get_tasks_hello_description_text():
    yaml_config = get_yaml_config()
    try:
        yaml_config.get(['tasks', 'hello', 'description', 'text'])
        assert False, "Error Expected"
    except Exception as e:
        assert str(e) == "`say hello world` is neither list or dictionary"


def test_yaml_config_get_tasks_hello_start():
    yaml_config = get_yaml_config()
    assert len(yaml_config.get(['tasks', 'hello', 'start'])) == 3


def test_yaml_config_get_tasks_hello_check():
    yaml_config = get_yaml_config()
    try:
        yaml_config.get(['tasks', 'hello', 'check'])
        assert False, "Error Expected"
    except Exception as e:
        assert str(e) == "`ordereddict([('description', 'say hello world'), ('start', ['bash', '-c', 'echo hello world'])])` has no key `check`"


def test_yaml_config_get_tasks_hello_start_0():
    yaml_config = get_yaml_config()
    assert yaml_config.get(['tasks', 'hello', 'start', 0]) == 'bash'


def test_yaml_config_get_tasks_hello_start_3():
    yaml_config = get_yaml_config()
    try:
        yaml_config.get(['tasks', 'hello', 'start', 3])
        assert False, "Error Expected"
    except Exception as e:
        assert str(e) == "`['bash', '-c', 'echo hello world']` has no index `3`"