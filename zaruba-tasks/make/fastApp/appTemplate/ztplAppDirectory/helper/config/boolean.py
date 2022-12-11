import os

def get_boolean_env(env_name: str, default_value: bool) -> bool:
    env_value = os.getenv(env_name, '')
    if env_value == '': # not provided
        return default_value
    lower_env_value = env_value.lower()
    if lower_env_value in ('0', 'no', 'false'): # clearly false
        return False
    if lower_env_value in ('1', 'yes', 'true'): # clearly true
        return True
    return default_value # other value