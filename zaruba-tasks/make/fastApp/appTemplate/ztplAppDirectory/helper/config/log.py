from typing import Mapping
import os
import logging


def get_log_level_from_env(
    env_name,
    default_value: int = logging.INFO
) -> int:
    '''
    Get environment value as logging level value
    '''
    env_value = os.getenv(env_name, '')
    lower_env_value = env_value.lower()
    result_map: Mapping[str, int] = {
        'critical': logging.CRITICAL,
        'fatal': logging.FATAL,
        'error': logging.ERROR,
        'warning': logging.WARNING,
        'warn': logging.WARN,
        'info': logging.INFO,
        'debug': logging.DEBUG,
        'notest': logging.NOTSET
    }
    if lower_env_value in result_map:
        return result_map[lower_env_value]
    logging.warning(' '.join([
        'Env {}'.format(env_name),
        'contains invalid log level value {},'.format(env_value),
        'use default value: {}'.format(default_value)
    ]))
    return default_value
