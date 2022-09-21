import os
from configs.port import http_port

create_oauth_access_token_url_path = os.getenv('APP_CREATE_OAUTH_ACCESS_TOKEN_URL_PATH', '/api/v1/create-oauth-access-token/')
create_access_token_url_path = os.getenv('APP_CREATE_ACCESS_TOKEN_URL_PATH', '/api/v1/create-access-token/')
renew_access_token_url_path = os.getenv('APP_RENEW_ACCESS_TOKEN_URL_PATH', '/api/v1/refresh-access-token/')
public_url_path = os.getenv('APP_PUBLIC_URL_PATH', '/public')

backend_url = os.getenv('APP_UI_BACKEND_URL', 'http://localhost:{}'.format(http_port))