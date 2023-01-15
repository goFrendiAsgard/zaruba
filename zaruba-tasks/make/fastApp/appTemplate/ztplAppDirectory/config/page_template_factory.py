from fastapi.templating import Jinja2Templates
from config.dir import page_dir
from config.ui import site_name, tagline, footer, renew_cred_token_interval
from config.url import (
    backend_address, public_url, renew_cred_token_url
)

import os


def escape_template(string: str) -> str:
    return '{{' + string + '}}'


def create_page_template() -> Jinja2Templates:
    '''
    Initiate Jinja2Templates and set it's default env.globals.
    '''
    tpl = Jinja2Templates(directory=page_dir)
    tpl.env.globals['site_name'] = site_name
    tpl.env.globals['tagline'] = tagline
    tpl.env.globals['footer'] = footer
    tpl.env.globals['backend_address'] = backend_address
    tpl.env.globals['public_url'] = public_url
    tpl.env.globals['renew_cred_token_url'] = renew_cred_token_url
    tpl.env.globals['renew_cred_token_interval'] = renew_cred_token_interval
    tpl.env.globals['vue'] = escape_template
    tpl.env.globals['getenv'] = os.getenv
    return tpl
