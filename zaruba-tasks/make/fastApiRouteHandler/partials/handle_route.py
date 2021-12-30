
@app.ztplAppHttpMethod('ztplAppUrl', response_model=str)
def ztplAppHttpMethod_ztpl_app_url(token: str = Depends(oauth2_scheme)) -> str:
    try:
        # NOTE: To send event, do this
        #   mb.call('event_name', {})
        # NOTE: To call rpc, do this
        #   rpc.call('rpc_name', [parameter])
        return 'OK'
    except Exception as error:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
