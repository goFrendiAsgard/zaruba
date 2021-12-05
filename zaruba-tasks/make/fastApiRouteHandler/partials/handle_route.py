
@app.ztplAppHttpMethod('ztplAppUrl')
def ztplAppHttpMethod_ztpl_app_url() -> str:
    try:
        # To send event: 
        #   mb.call('event_name', {})
        # To call rpc:
        #   rpc.call('rpc_name', [parameter])
        return 'OK'
    except Exception as error:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')
