from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles

def register_public_dir_route_handler(app: FastAPI, public_url: str, public_dir: str, public_route_name: str='static'):
    if public_dir != '':
        app.mount(public_url, StaticFiles(directory=public_dir), name=public_route_name)
        print('Register static directory route')
