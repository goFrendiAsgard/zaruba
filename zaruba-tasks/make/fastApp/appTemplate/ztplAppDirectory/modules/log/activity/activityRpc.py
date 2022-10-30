from typing import Any, Optional, Mapping
from helpers.transport import RPC, MessageBus
from schemas.activity import Activity, ActivityData
from schemas.user import User
from modules.log.activity.repos.activityRepo import ActivityRepo
from modules.log.activity.activityService import ActivityService

import sys

def register_activity_entity_rpc(mb: MessageBus, rpc: RPC, activity_repo: ActivityRepo):

    activity_service = ActivityService(mb, rpc, activity_repo)

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
    def insert_activity(activity_data: Mapping[str, Any], current_user_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        current_user = User.parse_obj(current_user_data)
        activity = ActivityData.parse_obj(activity_data) 
        new_activity = activity_service.insert(activity, current_user)
        return None if new_activity is None else new_activity.dict()


    print('Handle RPC for log.Activity', file=sys.stderr)