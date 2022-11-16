import os

message_bus_type: str = os.getenv('APP_MESSAGE_BUS_TYPE', 'local')
