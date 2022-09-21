from fastapi import FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates

from modules.ui import PageTemplateException
from helpers.transport import MessageBus, RPC
from configs.cors import cors_allow_credentials, cors_allow_headers, cors_allow_methods, cors_allow_origin_regex, cors_allow_origins, cors_expose_headers, cors_max_age
from configs.dir import public_dir
from configs.error import error_threshold
from configs.ui import site_name
from configs.url import public_url_path 

def create_app(mb: MessageBus, rpc: RPC, page_template: Jinja2Templates) -> FastAPI:
    app = FastAPI(title=site_name)

    # apply CORS middleware
    app.add_middleware(
        CORSMiddleware,
        allow_origins = cors_allow_origins,
        allow_origin_regex = cors_allow_origin_regex,
        allow_methods = cors_allow_methods,
        allow_headers = cors_allow_headers,
        allow_credentials = cors_allow_credentials,
        expose_headers = cors_expose_headers,
        max_age = cors_max_age,
    )

    # handle any exception with PageTemplateException type
    @app.exception_handler(PageTemplateException)
    def handle_template_exception(request: Request, exception: PageTemplateException):
        menu_context = exception.menu_context
        return page_template.TemplateResponse(
            'default_error.html',
            context={
                'request': request,
                'status_code': exception.status_code,
                'detail': exception.detail, 
                'context': menu_context
            },
            status_code=exception.status_code
        )

    # handle readiness request
    @app.get('/readiness')
    def handle_readiness():
        if mb.is_failing():
            raise HTTPException(status_code=500, detail='Messagebus is failing')
        if rpc.is_failing():
            raise HTTPException(status_code=500, detail='RPC is failing')
        if mb.get_error_count() > error_threshold:
            raise HTTPException(status_code=500, detail='Messagebus error exceeding threshold')
        if rpc.get_error_count() > error_threshold:
            raise HTTPException(status_code=500, detail='RPC error exceeding threshold')
        return HTMLResponse(content='ready', status_code=200)

    # handle application shutdown
    @app.on_event('shutdown')
    def on_shutdown():
        mb.shutdown()
        rpc.shutdown()
    print('Register app shutdown handler')

    if public_dir != '':
        # serve public static directory (js, css, html, images, etc)
        app.mount(public_url_path, StaticFiles(directory=public_dir), name='static-resources')
        print('Register static directory route')

    return app