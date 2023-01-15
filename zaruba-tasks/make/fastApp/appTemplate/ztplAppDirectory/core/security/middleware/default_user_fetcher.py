from typing import Optional, Callable
from fastapi.security import OAuth2
from fastapi import Depends, Cookie, HTTPException, status
from starlette.requests import Request
from schema.user import User
from helper.transport.rpc import RPC
from core.security.middleware.user_fetcher import UserFetcher

import sys
import traceback


class DefaultUserFetcher(UserFetcher):
    '''
    Default user fetcher.

    You can use DefaultUserFetcher to create a user-fetcher function.
    '''

    def __init__(self, rpc: RPC, oauth2_scheme: OAuth2):
        '''
        Initiate DefaultUserFetcher.
        You can use DefaultUserFetcher to create a user-fetcher function.
        '''
        self.rpc = rpc
        self.oauth2_scheme = oauth2_scheme

    def get_user_fetcher(
        self,
        throw_error: bool = True
    ) -> Callable[[Request], Optional[User]]:
        '''
        Create a function to fetch user
        '''
        async def fetch_user(
            bearer_token=Depends(self.oauth2_scheme),
            app_cred_token=Cookie(default=None)
        ) -> Optional[User]:
            cred_token = self._get_cred_token(bearer_token, app_cred_token)
            if cred_token is None:
                return None
            try:
                user_data = self.rpc.call('get_user_by_cred_token', cred_token)
                return None if user_data is None else User.parse_obj(user_data)
            except Exception:
                print(
                    'Error while fetching user with token {token}'.format(
                        token=cred_token
                    ),
                    file=sys.stderr
                )
                print(traceback.format_exc(), file=sys.stderr)
                return self._raise_error_or_return_none(
                    throw_error,
                    status.HTTP_500_INTERNAL_SERVER_ERROR,
                    'Cannot fetch user'
                )
        return fetch_user

    def _get_cred_token(
        self,
        bearer_token: str,
        app_cred_token: str
    ) -> Optional[str]:
        if bearer_token is not None:
            return bearer_token
        return app_cred_token

    def _raise_error_or_return_none(
        self,
        throw_error: bool,
        status_code: int,
        detail: str
    ) -> None:
        if not throw_error:
            return None
        raise HTTPException(
            status_code=status_code,
            detail=detail,
            headers={'WWW-Authenticate': 'Bearer'},
        )
