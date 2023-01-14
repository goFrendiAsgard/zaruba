import os
from helper.config import get_boolean_env

db_url: str = os.getenv(
    'APP_SQLALCHEMY_DATABASE_URL', 'sqlite:///database.db'
)
db_create_all: bool = get_boolean_env(
    'APP_DB_CREATE_ALL', True
)
