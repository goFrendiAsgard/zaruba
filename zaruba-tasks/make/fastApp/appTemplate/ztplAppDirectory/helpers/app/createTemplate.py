from fastapi.templating import Jinja2Templates

import os

def escape_template(string: str) -> str:
    return '{{' + string + '}}'

def create_templates(directory: str, guest_username: str, site_name: str, backend_url: str, public_url: str, renew_access_token_url: str, renew_access_token_interval: int) -> Jinja2Templates:
    templates = Jinja2Templates(directory=directory)
    templates.env.globals['guest_username'] = guest_username
    templates.env.globals['site_name'] = site_name
    templates.env.globals['backend_url'] = backend_url
    templates.env.globals['public_url'] = public_url
    templates.env.globals['renew_access_token_url'] = renew_access_token_url
    templates.env.globals['renew_access_token_interval'] = renew_access_token_interval
    templates.env.globals['vue'] = escape_template
    templates.env.globals['getenv'] = os.getenv
    return templates
