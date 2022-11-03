from typing import Optional, Tuple
from schemas.activity import Activity, ActivityData
from core.security.service.test_authService_util import init_test_auth_service_components
from modules.log.activity.repos.activityRepo import ActivityRepo
from modules.log.activity.activityService import ActivityService
from modules.log.activity.repos.dbActivityRepo import DBActivityRepo
from helpers.transport import LocalRPC, LocalMessageBus
from sqlalchemy import create_engine
from transport import AppMessageBus, AppRPC

def create_activity_data():
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    dummy_activity_data = ActivityData(
        activity='',
        user_id='',
        object='',
        row_id='',
        row='',
        created_by=''
    )
    return dummy_activity_data


def insert_activity_data(activity_repo: ActivityRepo, index: Optional[int] = None) -> Activity:
    activity_data = create_activity_data()
    activity_data.activity = 'activity' if index is None else 'activity-{index}'.format(index=index)
    activity_data.created_by = 'original_user'
    activity_data.updated_by = 'original_user'
    return activity_repo.insert(activity_data)


def init_test_activity_service_components() -> Tuple[ActivityService, DBActivityRepo, AppMessageBus, AppRPC]:
    auth_service, _, _, _ = init_test_auth_service_components()
    engine = create_engine('sqlite://', echo=False)
    activity_repo = DBActivityRepo(engine=engine, create_all=True)
    mb = AppMessageBus(LocalMessageBus())
    rpc = AppRPC(LocalRPC())
    activity_service = ActivityService(mb, rpc, auth_service, activity_repo)
    return activity_service, auth_service, activity_repo, mb, rpc
