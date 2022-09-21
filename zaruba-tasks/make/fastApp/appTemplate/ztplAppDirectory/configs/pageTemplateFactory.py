from fastapi.templating import Jinja2Templates
from configs.dir import page_dir
from configs.auth import guest_username
from configs.ui import site_name, renew_access_token_interval
from configs.url import backend_url, public_url_path, renew_access_token_url_path

import os

def escape_template(string: str) -> str:
    return '{{' + string + '}}'

def create_page_template() -> Jinja2Templates:
    templates = Jinja2Templates(directory=page_dir)
    templates.env.globals['guest_username'] = guest_username
    templates.env.globals['site_name'] = site_name
    templates.env.globals['backend_url'] = backend_url
    templates.env.globals['public_url_path'] = public_url_path
    templates.env.globals['renew_access_token_url_path'] = renew_access_token_url_path
    templates.env.globals['renew_access_token_interval'] = renew_access_token_interval
    templates.env.globals['vue'] = escape_template
    templates.env.globals['getenv'] = os.getenv
    return templates
