from enum import IntEnum

class AuthType(IntEnum):
    ANYONE = 0
    NON_USER = 1
    USER = 2
    HAS_PERMISSION = 3