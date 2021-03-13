from helper import cli
import helper.generator as generator


@cli
def show_service_name(service_location: str):
    print(generator.get_service_name(service_location))

if __name__ == '__main__':
    show_service_name()