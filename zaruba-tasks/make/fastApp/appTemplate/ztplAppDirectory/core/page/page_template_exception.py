from schema.menu_context import MenuContext


class PageTemplateException(Exception):
    '''
    Exception that should be rendered as a
    human-friendly error page.
    '''
    status_code: int
    detail: str
    menu_context: MenuContext

    def __init__(
        self, status_code: int, detail: str, menu_context: MenuContext
    ):
        '''
        Init a PageTemplateException.

        PageTemplateException should be rendered as a
        human-friendly error page.
        '''
        self.status_code = status_code
        self.detail = detail
        self.menu_context = menu_context
