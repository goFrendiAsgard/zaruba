from helper import cli
from helper.config import YamlConfig

@cli
def add_link(source: str, destination: str, file_name='./default.values.yaml'):
    yaml_config = YamlConfig()
    yaml_config.load(file_name)
    yaml_config.set(['link::{}'.format(destination)], source)
    yaml_config.save(file_name)
    

if __name__ == '__main__':
    add_link()
