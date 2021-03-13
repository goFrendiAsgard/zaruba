from helper import cli
import helper.generator as generator
import helper.decoration as decoration

@cli
def add_link(source: str, destination: str, file_name='./default.values.yaml'):
    config = generator.read_config(file_name)
    print('{yellow}Add link from "{source}" to "{destination}" on "{file_name}"{normal}'.format(yellow=decoration.yellow, normal=decoration.normal, source=source, destination=destination, file_name=file_name))
    config['link::{}'.format(destination)] = source
    generator.write_config(file_name, config)

if __name__ == '__main__':
    add_link()
