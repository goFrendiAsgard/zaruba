from helper.config import get_boolean_env

import os

seed_default_content_type: bool = get_boolean_env(
    'APP_SEED_DEFAULT_CONTENT_TYPE', True
)
default_content_type_name: str = os.getenv(
    'APP_DEFAULT_CONTENT_TYPE_NAME', 'article'
)
