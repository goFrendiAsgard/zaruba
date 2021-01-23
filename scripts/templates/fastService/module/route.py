from typing import List
from fastapi import FastAPI, HTTPException
from module import schema

import traceback
import transport

# 💡 HINT: 
#
#   * If you need other components beside `app` and `mb`, please:
#       * Add them as parameter of `init` function
#       * Declare the component at `main.py`
#   * Whenever possible, don't put business logic here. Instead, try to:
#       * Invoke RPC call (i.e: `mb.call_rpc(rpc_name, *args)`) or
#       * Encapsulate your business logic into another class/function 
#         so that you can import it here
#   * Visit fastapi documentation at: https://fastapi.tiangolo.com/tutorial/first-steps/
#
#
# 📝 EXAMPLE:
#
#   import time
#
#   def init(app: FastAPI, mb: transport.MessageBus):
#
#       @app.get('/favorite-number')
#       def get_favorite_number():
#           # publish hit event to messagebus
#           hit_time = time.strftime('%Y-%m-%d %H:%M:%S')
#           mb.publish('hit', {'url': '/favorite-number', 'time': hit_time})
#           # return favorite number (ref: https://bigbangtheory.fandom.com/wiki/73)
#           num = mb.call_rpc('add', 70, 3)
#           return {'favorite_number': num}


def init(app: FastAPI, mb: transport.MessageBus):
    print('Init {} route handlers'.format('module'))
