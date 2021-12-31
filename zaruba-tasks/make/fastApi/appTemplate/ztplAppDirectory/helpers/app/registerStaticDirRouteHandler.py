from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles

def register_static_dir_route_handler(app: FastAPI, static_url: str, static_dir: str, static_route_name: str='static'):
    if static_dir != '':
        app.mount(static_url, StaticFiles(directory=static_dir), name=static_route_name)
        print('Register static directory route')
