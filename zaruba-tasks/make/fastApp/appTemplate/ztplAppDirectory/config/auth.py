import os
from helper.config import get_boolean_env

root_username: str = os.getenv(
    'APP_ROOT_INITIAL_USERNAME', 'root'
)
root_initial_email: str = os.getenv(
    'APP_ROOT_INITIAL_EMAIL', 'root@innistrad.com'
)
root_initial_phone_number: str = os.getenv(
    'APP_ROOT_INITIAL_PHONE_NUMBER', '+621234567890'
)
root_initial_password: str = os.getenv(
    'APP_ROOT_INITIAL_PASSWORD', 'Alch3mist'
)
root_initial_fullname: str = os.getenv(
    'APP_ROOT_INITIAL_FULL_NAME', 'root'
)
root_permission: str = os.getenv(
    'APP_ROOT_PERMISSION', 'root'
)

cred_token_secret_key: str = os.getenv(
    'APP_CRED_TOKEN_SECRET_KEY', '123'
)
cred_token_algorithm: str = os.getenv(
    'APP_CRED_TOKEN_ALGORITHM', 'HS256'
)
cred_token_expire: float = float(os.getenv(
    'APP_CRED_TOKEN_EXPIRE', '1800'
))
seed_root_user: bool = get_boolean_env(
    'APP_SEED_ROOT_USER', True
)
