from typing import Optional
from jose import JWTError, jwt
from datetime import datetime, timedelta
from pydantic import BaseModel
from schemas.user import User, UserData

class Token(BaseModel):
    access_token: str
    token_type: str

class TokenData(BaseModel):
    username: Optional[str] = None

class TokenModel():

    def __init__(self, access_token_secret_key: str, access_token_algorithm: str, access_token_expire_minutes: str):
        self.access_token_secret_key = access_token_secret_key
        self.access_token_algorithm = access_token_algorithm
        self.access_token_expire_minutes = access_token_expire_minutes
    
    def create_token(self, string:str) -> str: 
        access_token_expires = timedelta(minutes=self.access_token_expire_minutes)
        to_encode = {"sub": string}
        expire = datetime.utcnow() + access_token_expires
        to_encode.update({"exp": expire})
        encoded_jwt = jwt.encode(to_encode, self.access_token_secret_key, algorithm=self.access_token_algorithm)
        return encoded_jwt

    def extract_from_token(self, token: str) -> str:
        payload = jwt.decode(token, self.access_token_secret_key, algorithms=[self.access_token_algorithm])
        return payload.get("sub")