
@app.ztplAppHttpMethod('ztplAppUrl', response_class=HTMLResponse)
def ztplAppHttpMethod_ztpl_app_url(current_user = Depends(auth_model.everyone())) -> HTMLResponse:
    # NOTE: To make this page require authentication, you can replace current_user = Depends(auth_model.everyone()) with:
    #   current_user = Depends(auth_model.authenticated())
    # or
    #   current_user = Depends(auth_model.has_any_permission('your_permission'))
    try:
        # NOTE: To send event, do this
        #   mb.call('event_name', {'some': 'object'})
        # NOTE: To call rpc, do this
        #   rpc.call('rpc_name', parameter1, parameter2,...)
        greetings = 'hello {}'.format(current_user.username)
        return HTMLResponse(content=greetings, status_code=200)
    except:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
