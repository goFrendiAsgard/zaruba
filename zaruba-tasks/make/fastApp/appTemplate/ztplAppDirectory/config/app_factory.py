from typing import List
from fastapi import FastAPI, Request, HTTPException, Depends
from fastapi.responses import HTMLResponse, JSONResponse
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from starlette.exceptions import HTTPException as StarletteHTTPException
from core import MenuService, PageTemplateException
from transport import AppMessageBus, AppRPC
from schema.menu_context import MenuContext
from schema.auth_type import AuthType

import re
import logging


def create_app(
    mb: AppMessageBus,
    rpc: AppRPC,
    menu_service: MenuService,
    page_template: Jinja2Templates,
    enable_ui: bool = True,
    enable_error_page: bool = True,
    cors_allow_credentials: bool = False,
    cors_allow_headers: List[str] = ['*'],
    cors_allow_methods: List[str] = ['*'],
    cors_allow_origin_regex: str = '',
    cors_allow_origins: List[str] = [],
    cors_expose_headers: bool = False,
    cors_max_age: int = 600,
    public_dir: str = 'public',
    error_threshold: int = 10,
    site_name: str = 'App',
    public_url: str = '/public',
    readiness_url: str = '/readiness'
) -> FastAPI:
    '''
    ⚛️ Create a FastAPI app instance with:
    - CORS middleware (as defined by the arguments).
    - route handlers for:
        - root ('/').
        - readiness (as defined by readiness_url by the arguments).
        - public directory (as defined by public_url by the arguments).
    - app shutdown event handler.

    Keyword arguments:
    - mb -- Instance of transport.AppMessageBus.
    - rpc -- Instance of transport.AppRPC.
    - menu_service -- Instance of core.MenuService.
    - page_template -- Instance of fastapi.templating.Jinja2Template.
    - enable_ui -- Whether UI is enabled or not (default: True).
    - enable_error_page -- Whether error page is enabled or not
        (default: True).
    - cors_allow_credentials -- Whether CORS credential is allowed or not
        (default: False).
    - cors_allow_headers -- List of CORS headers allowed (default: ['*']).
    - cors_allow_methods -- List of CORS methods allowed (default: ['*']).
    - cors_origin_regex -- Regex pattern for allow origin (default: '').
    - cors_expose_headers -- Whether CORS header is exposed or not
        (default: False).
    - cors_max_age -- CORS max age (default: 600).
    - public_dir -- Public directory location (default: 'public').
    - error_threshold -- Error threshold before readiness failed (default: 10).
    - site_name -- Site name (default: 'App').
    - public_url -- Public URL (default: '/public').
    - readiness_url -- Public URL (default: '/readiness').

    Feel free to add more things as you need.
    '''
    app = FastAPI(title=site_name)

    # 🔀 Apply CORS middleware
    app.add_middleware(
        CORSMiddleware,
        allow_origins=cors_allow_origins,
        allow_origin_regex=cors_allow_origin_regex,
        allow_methods=cors_allow_methods,
        allow_headers=cors_allow_headers,
        allow_credentials=cors_allow_credentials,
        expose_headers=cors_expose_headers,
        max_age=cors_max_age,
    )

    # 🥑 Handle readiness request
    @app.get(readiness_url)
    def handle_readiness():
        '''
        Handle readiness request.

        Respond with:
        - 200 status code if everything works as intended.
        - 500 status code if there is anything wrong with the system.
        '''
        if mb.is_failing():
            raise HTTPException(
                status_code=500, detail='messagebus is failing')
        if rpc.is_failing():
            raise HTTPException(status_code=500, detail='RPC is failing')
        if mb.get_error_count() > error_threshold:
            raise HTTPException(
                status_code=500, detail='messagebus error exceeding threshold')
        if rpc.get_error_count() > error_threshold:
            raise HTTPException(
                status_code=500, detail='RPC error exceeding threshold')
        return HTMLResponse(content='ready', status_code=200)

    # 🔚 Handle application shutdown
    @app.on_event('shutdown')
    def on_shutdown():
        '''
        Handle shutdown event.

        By default will also kill MessageBus and RPC.
        '''
        mb.shutdown()
        rpc.shutdown()
    logging.info('Register app shutdown handler')

    # 📢 serve public static directory (js, css, html, images, etc)
    if public_dir != '':
        app.mount(
            public_url,
            StaticFiles(directory=public_dir),
            name='static-resources'
        )
        logging.info('Register static directory route')

    # 🏠 Serve home page
    if enable_ui:
        menu_service.add_menu(
            name='home',
            title='Home',
            url='/',
            auth_type=AuthType.ANYONE
        )

        @app.get('/', response_class=HTMLResponse)
        async def get_home(
            request: Request,
            context: MenuContext = Depends(
                menu_service.has_access('home')
            )
        ) -> HTMLResponse:
            '''
            Handle root ('/') URL.
            '''
            try:
                return page_template.TemplateResponse(
                    'default_page.html',
                    context={
                        'request': request,
                        'context': context,
                        'content_path': 'home.html'
                    },
                    status_code=200
                )
            except Exception:
                logging.error('Non HTTPException error', exc_info=True)
                return page_template.TemplateResponse(
                    'default_error.html',
                    context={
                        'request': request,
                        'status_code': 500,
                        'detail': 'Internal server error'
                    },
                    status_code=500
                )

    # ❌ Handle any PageTemplateException
    if enable_ui and enable_error_page:
        @app.exception_handler(PageTemplateException)
        def handle_page_template_exception(
            request: Request, exception: PageTemplateException
        ):
            '''
            Handle PageTemplateException.
            '''
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

    # ❌ Handle any StarletteHTTPException
    if enable_ui and enable_error_page:
        @app.exception_handler(StarletteHTTPException)
        def handle_starlette_template_exception(
            request: Request, exception: StarletteHTTPException
        ):
            '''
            Handle StarletteHTTPException.
            '''
            url = request.url.path
            if re.search('/api/', url):
                # Do not override error generated by /api/
                return JSONResponse(
                    content={'detail': exception.detail},
                    status_code=exception.status_code
                )
            return page_template.TemplateResponse(
                'default_error.html',
                context={
                    'request': request,
                    'status_code': exception.status_code,
                    'detail': exception.detail,
                    'context': None
                },
                status_code=exception.status_code
            )

    return app
