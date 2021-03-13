from helper import cli
import helper.generator as generator
import helper.decoration as decoration

@cli
def set_project_value(key, value, file_name='./default.values.yaml'):
    config = generator.read_config(file_name)
    print('{yellow}Set "{key}" into "{value}" on "{file_name}"{normal}'.format(yellow=decoration.yellow, normal=decoration.normal, key=key, value=value, file_name=file_name))
    config[key] = value
    generator.write_config(file_name, config)


if __name__ == '__main__':
    set_project_value()

