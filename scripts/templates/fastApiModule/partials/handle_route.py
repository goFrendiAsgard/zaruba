@self.app.zarubaHttpMethod('zarubaUrl')
def zarubaHttpMethod_zaruba_url():
    try:
        return 'OK'
    except HTTPException as error:
        raise error
    except Exception as error:
        print(traceback.format_exc()) 
        raise HTTPException(status_code=500, detail='Internal Server Error')