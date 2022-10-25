from enum import IntEnum

class AuthType(IntEnum):
    ANYONE = 0
    VISITOR = 1
    USER = 2
    HAS_PERMISSION = 3