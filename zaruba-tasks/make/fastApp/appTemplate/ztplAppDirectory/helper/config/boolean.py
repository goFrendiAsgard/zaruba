from typing import Mapping
import os
import logging


def get_boolean_env(env_name: str, default_value: bool) -> bool:
    '''
    Get environment value as boolean
    '''
    env_value = os.getenv(env_name, '')
    lower_env_value = env_value.lower()
    result_map: Mapping[str, bool] = {
        '0': False,
        'no': False,
        'false': False,
        '1': True,
        'yes': True,
        'true': True,
    }
    if lower_env_value in result_map:
        return result_map[lower_env_value]
    logging.warning(' '.join([
        'Env {}'.format(env_name),
        'contains invalid boolean value {},'.format(env_value),
        'use default value: {}'.format(default_value)
    ]))
    return default_value
