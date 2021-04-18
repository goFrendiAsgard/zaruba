from .config import YamlConfig


def get_yaml_config() -> YamlConfig:
    yaml_config = YamlConfig()
    yaml_config.load('./test_resources/yaml_config.yaml')
    return yaml_config


def test_yaml_config_append_lconfig():
    yaml_config = get_yaml_config()
    yaml_config.append(['lconfigs', 'ports'], '8080')
    yaml_config.append(['lconfigs', 'ports'], '3000')
    assert yaml_config.get(['lconfigs', 'ports', 0]) == '8080'
    assert yaml_config.get(['lconfigs', 'ports', 1]) == '3000'

