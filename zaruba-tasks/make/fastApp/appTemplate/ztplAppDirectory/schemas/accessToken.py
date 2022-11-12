from pydantic import BaseModel

class CreateAccessTokenRequest(BaseModel):
    username: str
    password: str

class CreateAccessTokenResponse(BaseModel):
    access_token: str
    token_type: str

class RenewAccessTokenRequest(BaseModel):
    access_token: str

class RenewAccessTokenResponse(BaseModel):
    access_token: str
    token_type: str

