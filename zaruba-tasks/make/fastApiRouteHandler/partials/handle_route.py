
@app.ztplAppHttpMethod('ztplAppUrl', response_class=HTMLResponse)
def ztplAppHttpMethod_ztpl_app_url(current_user = Depends(auth_model.everyone)) -> HTMLResponse:
    try:
        # NOTE: To send event, do this
        #   mb.call('event_name', {})
        # NOTE: To call rpc, do this
        #   rpc.call('rpc_name', [parameter])
        greetings = 'hello {}'.format(current_user.username)
        return HTMLResponse(content=greetings, status_code=200)
    except:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
