import os

def get_abs_static_dir(raw_static_dir: str) -> str:
    return os.path.abspath(raw_static_dir) if raw_static_dir != '' else ''
