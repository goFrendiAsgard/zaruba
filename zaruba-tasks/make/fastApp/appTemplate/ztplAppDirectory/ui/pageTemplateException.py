from schemas.menuContext import MenuContext
from schemas.menu import Menu
from fastapi import FastAPI, Request
from fastapi.templating import Jinja2Templates

class PageTemplateException(Exception):
    status_code: int
    detail: str
    menu_context: MenuContext

    def __init__(self, status_code: int, detail: str, menu_context: MenuContext):
        self.status_code = status_code
        self.detail = detail
        self.menu_context = menu_context
