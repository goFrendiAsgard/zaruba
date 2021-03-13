from helper import cli
import helper.generator as generator


@cli
def show_env_prefix(service_location: str) -> str:
    print(generator.get_env_prefix(service_location))


if __name__ == '__main__':
    show_env_prefix()