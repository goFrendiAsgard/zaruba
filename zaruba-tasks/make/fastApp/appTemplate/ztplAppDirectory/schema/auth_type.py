from enum import IntEnum


class AuthType(IntEnum):
    '''
    Authentication type.

    - ANYONE: Anyone is authorized
    - VISITOR: Only unauthenticated user is authorized
    - USER: Only authenticated user is authorized
    - HAS_PERMISSION: Only authenticated user with
        specific permision is authorized
    '''
    ANYONE = 0
    VISITOR = 1
    USER = 2
    HAS_PERMISSION = 3
