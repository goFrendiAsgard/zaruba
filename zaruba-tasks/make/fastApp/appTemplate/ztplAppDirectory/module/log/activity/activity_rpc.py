from typing import Any, Optional, Mapping
from helper.transport import RPC, MessageBus
from schema.activity import Activity, ActivityData
from schema.user import User
from module.log.activity.activity_service import ActivityService
from core import AuthService

import sys

def register_activity_rpc(mb: MessageBus, rpc: RPC, auth_service: AuthService, activity_service: ActivityService):

    @rpc.handle('find_activity')
    def find_activities(keyword: str, limit: int, offset: int, current_user_data: Optional[Mapping[str, Any]]) -> Mapping[str, Any]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        activity_result = activity_service.find(keyword, limit, offset, current_user)
        return activity_result.dict()


    @rpc.handle('find_activity_by_id')
    def find_activity_by_id(id: str, current_user_data: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        current_user = None if current_user_data is None else User.parse_obj(current_user_data)
        activity = activity_service.find_by_id(id, current_user)
        return None if activity is None else activity.dict()


    @rpc.handle('insert_activity')
    def insert_activity(activity_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        activity = ActivityData.parse_obj(activity_data) 
        new_activity = activity_service.insert(activity)
        return None if new_activity is None else new_activity.dict()


    print('Handle RPC for log.activity', file=sys.stderr)