import os
import json

cors_allow_origins = json.loads(os.getenv('APP_CORS_ALLOW_ORIGINS', '["*"]'))
cors_allow_origin_regex = os.getenv('APP_CORS_ALLOW_ORIGIN_REGEX', '')
cors_allow_methods = json.loads(os.getenv('APP_CORS_ALLOW_METHODS', '["*"]'))
cors_allow_headers = json.loads(os.getenv('APP_CORS_ALLOW_HEADERS', '["*"]'))
cors_allow_credentials = os.getenv('APP_CORS_ALLOW_CREDENTIALS', 'false') == 'true'
cors_expose_headers = os.getenv('APP_CORS_EXPOSE_HEADERS', 'false') == 'true'
cors_max_age = int(os.getenv('APP_CORS_MAX_AGE', '600'))
