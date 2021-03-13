from typing import Mapping
from helper import cli
import helper.generator as generator
import helper.decoration as decoration

@cli
def replace_all(location: str, **replace: Mapping[str, str]):
    print('{yellow}Replace content of "{location}"{normal}'.format(yellow=decoration.yellow, normal=decoration.normal, location=location))
    generator.replace_all(location, replace)


if __name__ == '__main__':
    replace_all()
   