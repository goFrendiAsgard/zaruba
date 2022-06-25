
@app.ztplAppHttpMethod('ztplAppUrl', response_class=HTMLResponse)
def ztplAppHttpMethod_ztpl_app_url(current_user: User = Depends(auth_service.everyone())) -> HTMLResponse:
    # NOTE: To make this page require authentication, you can replace current_user:  User = Depends(auth_service.everyone()) with:
    #   current_user:  User = Depends(auth_service.authenticated())
    # or
    #   current_user:  User = Depends(auth_service.has_any_permission('your_permission'))
    try:
        # NOTE: To send event, do this
        #   mb.publish('event_name', {'some': 'object'})
        # NOTE: To call rpc, do this
        #   rpc.call('rpc_name', parameter1, parameter2,...)
        greetings = 'hello {}'.format(current_user.username)
        return HTMLResponse(content=greetings, status_code=200)
    except:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
