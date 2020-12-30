from typing import Any, List, Mapping
from sqlalchemy.orm import sessionmaker, Session
from sqlalchemy.engine import Engine
import database, schema, transport

# 💡 HINT: 
#
#   * If you need other components beside `mb`, `engine` and `DBSession`, please:
#       * Add them as parameter of `init` function
#       * Declare the component at `main.py`
#   * Make sure every event/rpc handler only receive or return:
#       * Primitive data types (str, int, float, boolean)
#       * List
#       * Dictionary
#     List and dictionary can be nested but cannot contain custom object
#     This is necessary because RMQMessageBus will serialize your data into JSON Object
#
#
# 📝 EXAMPLE:
#
#   def init(mb: transport.MessageBus, engine: Engine, DBSession: sessionmaker):
#
#       @transport.handle('hit')
#       def handle_hit_event(msg: Mapping[str, Any]):
#           print('Receiving message: ', msg)
#
#       @transport.handle_rpc('add')
#       def handle_add_rpc(num1: int, num2: int):
#           return num1 + num2


def init(mb: transport.MessageBus, engine: Engine, DBSession: sessionmaker):
    print('Init {} event/rpc handlers'.format('module'))