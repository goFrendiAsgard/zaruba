from typing import Any, List, Mapping
from sqlalchemy.orm import sessionmaker, Session
from module import schema, crud

import database, transport

# 💡 HINT: 
#
#   * If you need other components beside `mb` and `DBSession`, please:
#       * Add them as parameter of `init` function
#       * Declare the component at `main.py`
#   * Whenever possible, don't put business logic here. Instead, try to:
#     Encapsulate your business logic into another class/function 
#     so that you can import it here
#   * Make sure every event/rpc handler only receive or return:
#       * Primitive data types (str, int, float, boolean)
#       * List
#       * Dictionary
#     List and dictionary can be nested but cannot contain custom object
#     This is necessary because RMQMessageBus will serialize your data into JSON Object
#   * If you really need to pass non primitive object (e.g: DBSession),
#     please inject them by using decorator (see the example below)
#
#
# 📝 EXAMPLE:
#
#   def init(mb: transport.MessageBus, DBSession: sessionmaker):
#
#       @transport.handle('hit')
#       def handle_hit_event(msg: Mapping[str, Any]):
#           print('Receiving message: ', msg)
#
#       @transport.handle_rpc('add')
#       def handle_add_rpc(num1: int, num2: int):
#           return num1 + num2
#
#       @transport.handle('loggedHit')
#       @database.handle(DBSession)
#       def handle_logged_hit(db: Session, msg: Mapping[str, Any]):
#           # log_to_db(db, msg) 
#           print('Receiving message: ', msg)
#           


def init(mb: transport.MessageBus, DBSession: sessionmaker):
    print('Init {} event/rpc handlers'.format('module'))