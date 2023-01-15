from helper.config import get_absolute_dir
import os

public_dir: str = get_absolute_dir(os.getenv(
    'APP_PUBLIC_DIR', 'public'
))
page_dir: str = get_absolute_dir(os.getenv(
    'APP_UI_PAGE_DIR', 'pages'
))
