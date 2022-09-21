import os

site_name = os.getenv('APP_UI_SITE_NAME', 'App')
renew_access_token_interval = int(os.getenv('APP_RENEW_ACCESS_TOKEN_INTERVAL', '300'))
