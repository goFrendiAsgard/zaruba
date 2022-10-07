
@app.get('ztplAppUrl', response_class=HTMLResponse)
async def get_ztpl_app_url(request: Request, context: MenuContext = Depends(menu_service.is_authorized('ztplAppModuleName:ztplAppUrl'))) -> HTMLResponse:
    '''
    Handle (get) ztplAppUrl
    '''
    try:
        return page_template.TemplateResponse('default_page.html', context={
            'request': request,
            'context': context,
            'content_path': 'ztplAppModuleName/ztpl_app_url.html'
        }, status_code=200)
    except:
        print(traceback.format_exc(), file=sys.stderr) 
        return page_template.TemplateResponse('default_error.html', context={
            'request': request,
            'status_code': 500,
            'detail': 'Internal server error'
        }, status_code=500)
