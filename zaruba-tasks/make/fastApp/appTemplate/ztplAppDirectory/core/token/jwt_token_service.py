from typing import Optional
from jose import jwt
from datetime import datetime, timedelta
from schema.user import User
from module.auth.user.user_service import UserService
from core.token.token_service import TokenService

import sys


class JWTTokenService(TokenService):

    def __init__(
        self, user_service: UserService,
        cred_token_secret_key: str,
        cred_token_algorithm: str, cred_token_expire: float
    ):
        self.user_service = user_service
        self.cred_token_secret_key = cred_token_secret_key
        self.cred_token_algorithm = cred_token_algorithm
        self.cred_token_expire = cred_token_expire

    def create_cred_token(
        self,
        user: User,
        current_user: Optional[User] = None
    ) -> str:
        cred_token_expires = timedelta(seconds=self.cred_token_expire)
        to_encode = {"sub": user.id}
        expire = datetime.now() + cred_token_expires
        to_encode.update({"exp": expire})
        encoded_jwt = jwt.encode(
            to_encode,
            self.cred_token_secret_key,
            algorithm=self.cred_token_algorithm
        )
        return encoded_jwt

    def get_user_by_cred_token(
        self, token: str,
        current_user: Optional[User] = None
    ) -> Optional[User]:
        if not token:
            return None
        try:
            payload = jwt.decode(
                token,
                self.cred_token_secret_key,
                algorithms=[self.cred_token_algorithm]
            )
            user_id = payload.get("sub")
            return self.user_service.find_by_id(user_id, current_user)
        except Exception:
            print('Error while getting user by token {token}'.format(
                token=token), file=sys.stderr
            )
            return None
