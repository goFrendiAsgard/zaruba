from typing import Optional, Tuple
from module.auth.role.role_service import RoleService
from module.auth.role.repo.db_role_repo import DBRoleRepo
from module.auth.role.repo.role_repo import RoleRepo
from schema.role import Role, RoleData
from helper.transport import LocalRPC, LocalMessageBus
from sqlalchemy import create_engine
from transport import AppMessageBus, AppRPC


def create_role_data() -> RoleData:
    # Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
    dummy_role_data = RoleData(
        name='',
        permissions=[],
        created_by=''
    )
    return dummy_role_data


def insert_role_data(role_repo: RoleRepo, index: Optional[int] = None) -> Role:
    role_data = create_role_data()
    role_data.name = 'original' if index is None else 'original-{index}'.format(index=index)
    role_data.created_by = 'original_user'
    role_data.updated_by = 'original_user'
    return role_repo.insert(role_data)


def create_mb() -> AppMessageBus:
    mb = AppMessageBus(LocalMessageBus())
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_role_service_components() -> Tuple[RoleService, DBRoleRepo, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    role_repo = DBRoleRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    role_service = RoleService(mb, rpc, role_repo)
    return role_service, role_repo, mb, rpc
