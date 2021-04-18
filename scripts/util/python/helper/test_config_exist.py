from .config import YamlConfig


def get_yaml_config() -> YamlConfig:
    yaml_config = YamlConfig()
    yaml_config.load('./test_resources/yaml_config.yaml')
    return yaml_config


def test_yaml_config_exist_tasks_hello_description():
    yaml_config = get_yaml_config()
    assert yaml_config.exist(['tasks', 'hello', 'description']) == True


def test_yaml_config_exist_tasks_hello_description_text():
    yaml_config = get_yaml_config()
    assert yaml_config.exist(['tasks', 'hello', 'description', 'text']) == False


def test_yaml_config_exist_tasks_hello_start():
    yaml_config = get_yaml_config()
    assert yaml_config.exist(['tasks', 'hello', 'start']) == True


def test_yaml_config_exist_tasks_hello_check():
    yaml_config = get_yaml_config()
    assert yaml_config.exist(['tasks', 'hello', 'check']) == False

def test_yaml_config_exist_tasks_hello_start_0():
    yaml_config = get_yaml_config()
    assert yaml_config.exist(['tasks', 'hello', 'start', 0]) == True


def test_yaml_config_exist_tasks_hello_start_3():
    yaml_config = get_yaml_config()
    yaml_config.exist(['tasks', 'hello', 'start', 3]) == False