from typing import Optional
from jose import jwt
from datetime import datetime, timedelta
from schemas.user import User
from modules.auth.user.userService import UserService

import abc
import sys

class TokenService(abc.ABC):

    @abc.abstractmethod
    def create_access_token(self, user: User, current_user: Optional[User]) -> str:
        pass

    @abc.abstractmethod
    def get_user_by_token(self, token: str, current_user: Optional[User]) -> Optional[User]:
        pass


class JWTTokenService(TokenService):

    def __init__(self, user_service: UserService, access_token_secret_key: str, access_token_algorithm: str, access_token_expire: float):
        self.user_service = user_service
        self.access_token_secret_key = access_token_secret_key
        self.access_token_algorithm = access_token_algorithm
        self.access_token_expire = access_token_expire


    def create_access_token(self, user: User, current_user: Optional[User] = None) -> str: 
        access_token_expires = timedelta(seconds=self.access_token_expire)
        to_encode = {"sub": user.id}
        expire = datetime.now() + access_token_expires
        to_encode.update({"exp": expire})
        encoded_jwt = jwt.encode(to_encode, self.access_token_secret_key, algorithm=self.access_token_algorithm)
        return encoded_jwt


    def get_user_by_token(self, token: str, current_user: Optional[User] = None) -> Optional[User]:
        if not token:
            return None
        try:
            payload = jwt.decode(token, self.access_token_secret_key, algorithms=[self.access_token_algorithm])
            user_id = payload.get("sub")
            return self.user_service.find_by_id(user_id, current_user)
        except:
            print('Error while getting user by token {token}'.format(token=token), file=sys.stderr)
            return None