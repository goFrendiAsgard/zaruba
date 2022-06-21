from enum import IntEnum

class AuthType(IntEnum):
    EVERYONE = 0
    UNAUTHENTICATED = 1
    AUTHENTICATED = 2
    AUTHORIZED = 3