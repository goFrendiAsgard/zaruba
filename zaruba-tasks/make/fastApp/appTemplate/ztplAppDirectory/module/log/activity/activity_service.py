from typing import Optional
from helper.transport import RPC, MessageBus
from core import AuthService
from schema.user import User
from schema.activity import Activity, ActivityData, ActivityResult
from module.log.activity.repo.activity_repo import ActivityRepo
from fastapi import HTTPException


class ActivityService():

    def __init__(
        self,
        mb: MessageBus,
        rpc: RPC,
        auth_service: AuthService,
        activity_repo: ActivityRepo
    ):
        self.mb = mb
        self.rpc = rpc
        self.auth_service = auth_service
        self.activity_repo = activity_repo

    def find(
        self,
        keyword: str,
        limit: int,
        offset: int,
        current_user: Optional[User] = None
    ) -> ActivityResult:
        count = self.activity_repo.count(keyword)
        rows = self.activity_repo.find(keyword, limit, offset)
        return ActivityResult(count=count, rows=rows)

    def find_by_id(
        self,
        id: str,
        current_user: Optional[User] = None
    ) -> Optional[Activity]:
        activity = self._find_by_id_or_error(id)
        return activity

    def insert(
        self,
        activity_data: ActivityData
    ) -> Optional[Activity]:
        system_user = self.auth_service.get_system_user()
        activity_data.created_by = system_user.id
        activity_data.updated_by = system_user.id
        activity_data = self._validate_data(activity_data)
        return self.activity_repo.insert(activity_data)

    def _find_by_id_or_error(
        self,
        id: Optional[str] = None
    ) -> Optional[Activity]:
        activity = self.activity_repo.find_by_id(id)
        if activity is None:
            raise HTTPException(
                status_code=404,
                detail='activity id not found: {}'.format(id)
            )
        return activity

    def _validate_data(
        self,
        activity_data: ActivityData,
        id: Optional[str] = None
    ) -> ActivityData:
        return activity_data
