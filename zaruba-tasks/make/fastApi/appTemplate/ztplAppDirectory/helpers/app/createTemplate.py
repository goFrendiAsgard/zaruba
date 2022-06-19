from fastapi.templating import Jinja2Templates

def escape_template(string: str) -> str:
    return '{{' + string + '}}'

def create_templates(directory: str) -> Jinja2Templates:
    templates = Jinja2Templates(directory=directory)
    templates.env.globals['vue'] = escape_template
    return templates
