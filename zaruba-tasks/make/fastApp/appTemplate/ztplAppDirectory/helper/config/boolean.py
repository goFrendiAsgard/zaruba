import os


def get_boolean_env(env_name: str, default_value: bool) -> bool:
    env_value = os.getenv(env_name, '')
    # env value is not provided
    if env_value == '':
        return default_value
    lower_env_value = env_value.lower()
    # env value is clearly false
    if lower_env_value in ('0', 'no', 'false'):
        return False
    # env value is clearly true
    if lower_env_value in ('1', 'yes', 'true'):
        return True
    # default value
    return default_value
