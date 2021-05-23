from helper import cli
import uuid

@cli
def get_uuid4():
    print(str(uuid.uuid4()))


if __name__ == '__main__':
    get_uuid4()
    