from schemas.menuContext import MenuContext
from schemas.menu import Menu
from fastapi import FastAPI, Request
from fastapi.templating import Jinja2Templates

class TemplateException(Exception):
    status_code: int
    detail: str
    menu_context: MenuContext

    def __init__(self, status_code: int, detail: str, menu_context: MenuContext):
        self.status_code = status_code
        self.detail = detail
        self.menu_context = menu_context


def register_template_exception_handler(app: FastAPI, templates: Jinja2Templates):
    @app.exception_handler(TemplateException)
    def handle_template_exception(request: Request, exception: TemplateException):
        menu_context = exception.menu_context
        return templates.TemplateResponse(
            'default_error.html',
            context={
                'request': request,
                'status_code': exception.status_code,
                'detail': exception.detail, 
                'context': menu_context
            },
            status_code=exception.status_code
        )