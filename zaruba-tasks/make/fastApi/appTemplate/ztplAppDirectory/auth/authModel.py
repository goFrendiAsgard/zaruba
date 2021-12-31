from typing import Callable, List
from fastapi.security import OAuth2PasswordBearer, OAuth2
from fastapi import Depends, FastAPI, HTTPException, status
from starlette.requests import Request
from auth.userModel import UserModel
from schemas.user import User

class AuthModel():

    def __init__(self, user_model: UserModel, oauth2_scheme: OAuth2, root_role: str):
        self.user_model = user_model
        self.oauth2_scheme = oauth2_scheme
        self.root_role = root_role

    def raise_unauthorized_exception(self, detail: str):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )

    async def everyone(self, request: Request) -> User:
        token = await self.oauth2_scheme(request)
        current_user = self.user_model.find_by_token(token)
        if not current_user:
            return self.user_model.get_guest_user()
        return current_user

    def current_user_has_any_role(self, roles: List[str]) -> Callable[[Request], User]:
        async def has_role(token = Depends(self.oauth2_scheme)) -> User:
            current_user = self.user_model.find_by_token(token)
            if not current_user:
                self.raise_unauthorized_exception('Invalid token')
            if len(roles) == 0:
                return current_user
            if current_user.has_role(self.root_role):
                return current_user
            for role in roles:
                if current_user.has_role(role):
                    return current_user
            self.raise_unauthorized_exception('Insufficient privilege')
        return has_role