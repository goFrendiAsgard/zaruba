from typing import Optional
from jose import JWTError, jwt
from datetime import datetime, timedelta
from pydantic import BaseModel
from schemas.user import User, UserData
from auth.userModel import UserModel

import abc

class TokenModel(abc.ABC):

    @abc.abstractmethod
    def create_user_token(self, string: str) -> str:
        pass

    @abc.abstractmethod
    def get_user_by_token(self, token: str) -> str:
        pass

class JWTTokenModel(TokenModel):

    def __init__(self, user_model: UserModel, access_token_secret_key: str, access_token_algorithm: str, access_token_expire_minutes: str):
        self.user_model = user_model
        self.access_token_secret_key = access_token_secret_key
        self.access_token_algorithm = access_token_algorithm
        self.access_token_expire_minutes = access_token_expire_minutes
    
    def create_user_token(self, user: User) -> str: 
        access_token_expires = timedelta(minutes=self.access_token_expire_minutes)
        to_encode = {"sub": user.id}
        expire = datetime.utcnow() + access_token_expires
        to_encode.update({"exp": expire})
        encoded_jwt = jwt.encode(to_encode, self.access_token_secret_key, algorithm=self.access_token_algorithm)
        return encoded_jwt

    def get_user_by_token(self, token: str) -> str:
        payload = jwt.decode(token, self.access_token_secret_key, algorithms=[self.access_token_algorithm])
        user_id = payload.get("sub")
        return self.user_model.find_by_id(user_id)