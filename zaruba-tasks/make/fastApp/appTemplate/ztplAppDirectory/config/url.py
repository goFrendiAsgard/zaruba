import os
from config.port import http_port
from helper.config import url_path

# Auth URLs
create_oauth_cred_token_url: str = url_path(os.getenv(
    'APP_CREATE_OAUTH_CRED_TOKEN_URL',
    '/api/v1/create-oauth-access-token'
))
create_cred_token_url: str = url_path(os.getenv(
    'APP_CREATE_CRED_TOKEN_URL',
    '/api/v1/create-access-token'
))
renew_cred_token_url: str = url_path(os.getenv(
    'APP_RENEW_CRED_TOKEN_URL',
    '/api/v1/refresh-access-token'
))

# Readiness URLs
readiness_url: str = url_path(os.getenv(
    'APP_READINESS_URL',
    '/readiness'
))

# Public URLs
public_url: str = url_path(os.getenv(
    'APP_PUBLIC_URL',
    '/public'
))

# Backend Address
backend_address: str = os.getenv(
    'APP_BACKEND_ADDRESS',
    'http://localhost:{}'.format(http_port)
).rstrip('/')
