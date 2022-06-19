from fastapi.templating import Jinja2Templates

def escape_template(string: str) -> str:
    return '{{' + string + '}}'

def create_templates(directory: str, guest_username: str, site_name: str) -> Jinja2Templates:
    templates = Jinja2Templates(directory=directory)
    templates.env.globals['guest_username'] = guest_username
    templates.env.globals['site_name'] = site_name
    templates.env.globals['vue'] = escape_template
    return templates
