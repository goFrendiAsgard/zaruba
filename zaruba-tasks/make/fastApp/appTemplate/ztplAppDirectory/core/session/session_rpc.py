from typing import Any, Optional, Mapping
from transport import AppMessageBus, AppRPC
from core.session.session_service import SessionService
from core.security.service.auth_service import AuthService
from schema.user import User

import logging


def register_session_rpc(
    mb: AppMessageBus,
    rpc: AppRPC,
    auth_service: AuthService,
    session_service: SessionService
):

    @rpc.handle('create_cred_token')
    def create_cred_token(
        identity: str,
        password: str,
        current_user_data: Optional[Mapping[str, Any]] = None
    ) -> str:
        current_user = _get_user(current_user_data)
        return session_service.create_cred_token(
            identity, password, current_user
        )

    @rpc.handle('renew_cred_token')
    def renew_cred_token(
        token: str,
        current_user_data: Optional[Mapping[str, Any]] = None
    ) -> str:
        current_user = _get_user(current_user_data)
        return session_service.renew_cred_token(token, current_user)

    @rpc.handle('get_user_by_cred_token')
    def get_user_by_cred_token(
        token: str,
        current_user_data: Optional[Mapping[str, Any]] = None
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user(current_user_data)
        user = session_service.get_user_by_cred_token(token, current_user)
        return None if user is None else user.dict()

    def _get_user(user_data: Optional[Mapping[str, Any]]) -> Optional[User]:
        if user_data is None:
            return None
        return User.parse_obj(user_data)

    logging.info('Register core.session RPC handler')
