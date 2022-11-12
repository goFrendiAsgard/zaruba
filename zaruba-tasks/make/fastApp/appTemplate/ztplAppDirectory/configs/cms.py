import os

seed_default_content_type: bool = os.getenv('APP_SEED_DEFAULT_CONTENT_TYPE', '1') != '0'
default_content_type_name: str = os.getenv('APP_DEFAULT_CONTENT_TYPE_NAME', 'article')