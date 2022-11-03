import os

site_name: str = os.getenv('APP_UI_SITE_NAME', 'App')
tagline: str = os.getenv('APP_UI_TAGLINE', 'Quod est superius est sicut quod inferius')
footer: str = os.getenv('APP_UI_FOOTER', 'App &copy; 2022-now')
renew_access_token_interval: int = int(os.getenv('APP_RENEW_ACCESS_TOKEN_INTERVAL', '300'))