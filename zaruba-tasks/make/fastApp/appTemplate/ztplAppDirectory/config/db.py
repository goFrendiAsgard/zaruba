import os

db_url: str = os.getenv('APP_SQLALCHEMY_DATABASE_URL', 'sqlite:///database.db')
db_create_all: bool = os.getenv('APP_DB_CREATE_ALL', '1') != '0'
