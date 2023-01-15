import os


def get_absolute_dir(raw_static_dir: str) -> str:
    '''
    Get absolute directory path
    '''
    return os.path.abspath(raw_static_dir) if raw_static_dir != '' else ''
