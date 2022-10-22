
@app.ztplAppHttpMethod('ztplAppUrl', response_class=HTMLResponse)
async def ztplAppHttpMethod_ztpl_app_url(current_user: Optional[User] = Depends(auth_service.everyone())) -> HTMLResponse:
    '''
    Handle (ztplAppHttpMethod) ztplAppUrl
    To enforce authorization, you can use any of these dependencies as parameter:
        - auth_service.everyone()
        - auth_service.is_unauthenticated()
        - auth_service.is_authenticated()
        - auth_service.is_authorized('permission')
    To publish an event, you can use:
        mb.publish('event_name', {'some': 'value'})
    To send RPC, you can use: 
        rpc.call('rpc_name', 'parameter1', 'parameter2')
    '''
    try:
        if not current_user:
            current_user = User.parse_obj(rpc.call('get_guest_user'))
        greetings = 'hello {}'.format(current_user.username)
        return HTMLResponse(content=greetings, status_code=200)
    except:
        print(traceback.format_exc(), file=sys.stderr) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
