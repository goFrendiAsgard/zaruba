from fastapi import FastAPI, HTTPException
import transport, schema

# 💡 HINT: 
#
#   * If you need other components beside `app` and `mb`, please:
#       * Add them as parameter of `init` function
#       * Declare the component at `main.py`
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
