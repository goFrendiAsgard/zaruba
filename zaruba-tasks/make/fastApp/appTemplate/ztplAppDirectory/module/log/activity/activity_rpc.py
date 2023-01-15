from typing import Any, Optional, Mapping
from helper.transport import RPC, MessageBus
from schema.activity import Activity, ActivityData
from schema.user import User
from module.log.activity.activity_service import ActivityService
from core import AuthService

import logging


def register_activity_rpc(
    mb: MessageBus,
    rpc: RPC,
    auth_service: AuthService,
    activity_service: ActivityService
):

    @rpc.handle('find_activity')
    def find_activities(
        keyword: str,
        limit: int,
        offset: int,
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Mapping[str, Any]:
        current_user = _get_user_from_dict(current_user_data)
        activity_result = activity_service.find(
            keyword, limit, offset, current_user
        )
        return activity_result.dict()

    @rpc.handle('find_activity_by_id')
    def find_activity_by_id(
        id: str, 
        current_user_data: Optional[Mapping[str, Any]]
    ) -> Optional[Mapping[str, Any]]:
        current_user = _get_user_from_dict(current_user_data)
        activity = activity_service.find_by_id(id, current_user)
        return _activity_as_dict(activity)

    @rpc.handle('insert_activity')
    def insert_activity(
        activity_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        activity = ActivityData.parse_obj(activity_data) 
        new_activity = activity_service.insert(activity)
        return _activity_as_dict(new_activity)

    def _get_user_from_dict(
        user_data: Optional[Mapping[str, Any]]
    ) -> Optional[User]:
        if user_data is None:
            return None
        return User.parse_obj(user_data)

    def _activity_as_dict(
        activity: Optional[Activity]
    ) -> Optional[Mapping[str, Any]]:
        if activity is None:
            return None
        return activity.dict()

    logging.info('Register log.activity RPC handler')
