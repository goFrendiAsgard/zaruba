import os

guest_username = os.getenv('APP_GUEST_USERNAME', 'guest')


root_username = os.getenv('APP_ROOT_INITIAL_USERNAME', 'root')
root_initial_email = os.getenv('APP_ROOT_INITIAL_EMAIL', 'root@innistrad.com')
root_initial_phone_number = os.getenv('APP_ROOT_INITIAL_PHONE_NUMBER', '+621234567890')
root_initial_password = os.getenv('APP_ROOT_INITIAL_PASSWORD', 'Alch3mist')
root_initial_fullname = os.getenv('APP_ROOT_INITIAL_FULL_NAME', 'root')
root_permission = os.getenv('APP_ROOT_PERMISSION', 'root')

access_token_secret_key = os.getenv('APP_ACCESS_TOKEN_SECRET_KEY', '123')
access_token_algorithm = os.getenv('APP_ACCESS_TOKEN_ALGORITHM', 'HS256')
access_token_expire = float(os.getenv('APP_ACCESS_TOKEN_EXPIRE', '1800'))
