from typing import Any, Optional, Mapping
from helper.transport import RPC, MessageBus
from schema.activity import ActivityData
from module.log.activity import ActivityService
from core import AuthService

import logging


def register_activity_event(
    mb: MessageBus,
    rpc: RPC,
    auth_service: AuthService,
    activity_service: ActivityService
):
    @mb.handle('new_activity')
    def insert_activity(
        activity_data: Mapping[str, Any]
    ) -> Optional[Mapping[str, Any]]:
        activity = ActivityData.parse_obj(activity_data)
        new_activity = activity_service.insert(activity)
        return None if new_activity is None else new_activity.dict()

    logging.info('Register log.activity event handler')
