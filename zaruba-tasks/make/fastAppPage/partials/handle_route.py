# ZtplAppUrlTitle page
menu_service.add_menu(name='ztplAppModuleName:ztplAppUrl', title='ZtplAppUrlTitle', url='ztpl-normalized-app-url', auth_type=AuthType.ANYONE, parent_name='ztplAppModuleName')
@app.get('ztpl-normalized-app-url', response_class=HTMLResponse)
async def get_ztpl_app_url(request: Request, context: MenuContext = Depends(menu_service.has_access('ztplAppModuleName:ztplAppUrl'))) -> HTMLResponse:
    '''
    Serve (get) ztpl-normalized-app-url
    '''
    try:
        return page_template.TemplateResponse('default_page.html', context={
            'request': request,
            'context': context,
            'content_path': 'modules/ztplAppModuleName/ztpl_app_url.html'
        }, status_code=200)
    except:
        print(traceback.format_exc(), file=sys.stderr) 
        return page_template.TemplateResponse('default_error.html', context={
            'request': request,
            'status_code': 500,
            'detail': 'Internal server error'
        }, status_code=500)
