from helper.config import get_log_level_from_env

import logging

log_level = get_log_level_from_env('APP_LOG_LEVEL', logging.INFO)
