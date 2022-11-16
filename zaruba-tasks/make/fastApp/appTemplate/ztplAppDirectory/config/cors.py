from typing import List
import os
import json

cors_allow_origins: List[str] = json.loads(os.getenv('APP_CORS_ALLOW_ORIGINS', '["*"]'))
cors_allow_origin_regex: str = os.getenv('APP_CORS_ALLOW_ORIGIN_REGEX', '')
cors_allow_methods: List[str] = json.loads(os.getenv('APP_CORS_ALLOW_METHODS', '["*"]'))
cors_allow_headers: List[str] = json.loads(os.getenv('APP_CORS_ALLOW_HEADERS', '["*"]'))
cors_allow_credentials: bool = os.getenv('APP_CORS_ALLOW_CREDENTIALS', '0') == '1'
cors_expose_headers: bool = os.getenv('APP_CORS_EXPOSE_HEADERS', '0') == '1'
cors_max_age: int = int(os.getenv('APP_CORS_MAX_AGE', '600'))
