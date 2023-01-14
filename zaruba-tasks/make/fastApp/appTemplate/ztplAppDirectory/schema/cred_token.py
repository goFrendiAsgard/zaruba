from pydantic import BaseModel


class CreateCredTokenRequest(BaseModel):
    username: str
    password: str


class CreateCredTokenResponse(BaseModel):
    cred_token: str
    token_type: str


class RenewCredTokenRequest(BaseModel):
    cred_token: str


class RenewCredTokenResponse(BaseModel):
    cred_token: str
    token_type: str
