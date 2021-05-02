from helper import cli
import helper.text as get_service_name


@cli
def show_service_name(service_location: str):
    print(get_service_name(service_location))

if __name__ == '__main__':
    show_service_name()