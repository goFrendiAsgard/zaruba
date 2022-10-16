import os


def get_absolute_dir(raw_static_dir: str) -> str:
    return os.path.abspath(raw_static_dir) if raw_static_dir != '' else ''

public_dir = get_absolute_dir(os.getenv('APP_PUBLIC_DIR', 'public'))
page_dir = get_absolute_dir(os.getenv('APP_UI_PAGE_DIR', 'pages'))
