import os

site_name = os.getenv('APP_UI_SITE_NAME', 'App')
tagline = os.getenv('APP_UI_TAGLINE', 'Quod est superius est sicut quod inferius')
footer = os.getenv('APP_UI_FOOTER', 'App &copy; 2022-now')
renew_access_token_interval = int(os.getenv('APP_RENEW_ACCESS_TOKEN_INTERVAL', '300'))