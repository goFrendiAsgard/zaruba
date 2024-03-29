
@app.ztplAppHttpMethod(
    'ztpl-normalized-app-url',
    response_class=HTMLResponse
)
async def ztplAppHttpMethod_ztpl_app_url(
    current_user: Optional[User] = Depends(auth_service.anyone())
) -> HTMLResponse:
    '''
    Handle (ztplAppHttpMethod) ztpl-normalized-app-url
    To enforce authorization, you can use any of these dependencies:
        - auth_service.anyone()
        - auth_service.is_visitor()
        - auth_service.is_user()
        - auth_service.has_permission('permission')
    To publish an event, you can use:
        mb.publish('event_name', {'some': 'value'})
    To send RPC, you can use: 
        rpc.call('rpc_name', 'parameter1', 'parameter2')
    '''
    try:
        if not current_user:
            current_user = User.parse_obj(auth_service.get_guest_user())
        greetings = 'hello {}'.format(current_user.username)
        return HTMLResponse(content=greetings, status_code=200)
    except Exception:
        logging.error('Non HTTPException error', exc_info=True)
        raise HTTPException(
            status_code=500,
            detail='Internal server serror'
        )
