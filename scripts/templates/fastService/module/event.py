from sqlalchemy.orm import sessionmaker, Session
from sqlalchemy.engine import Engine
import database, schema, transport

def init(mb: transport.MessageBus, engine: Engine, DBSession: sessionmaker):
    print('Init {} event/rpc handlers'.format('module'))
    pass