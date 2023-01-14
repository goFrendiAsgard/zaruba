from typing import Optional, Tuple
from module.ztpl_app_module_name.ztpl_app_crud_entity.repo.ztpl_app_crud_entity_repo import ZtplAppCrudEntityRepo
from schema.ztpl_app_crud_entity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from module.ztpl_app_module_name.ztpl_app_crud_entity.ztpl_app_crud_entity_service import ZtplAppCrudEntityService
from module.ztpl_app_module_name.ztpl_app_crud_entity.repo.db_ztpl_app_crud_entity_repo import DBZtplAppCrudEntityRepo
from helper.transport import LocalRPC, LocalMessageBus, MessageBus
from transport import AppMessageBus, AppRPC
from sqlalchemy import create_engine

def create_ztpl_app_crud_entity_data() -> ZtplAppCrudEntityData:
    # Note: 🤖 Don't delete the following statement
    dummy_ztpl_app_crud_entity_data = ZtplAppCrudEntityData(
        created_by=''
    )
    return dummy_ztpl_app_crud_entity_data


def insert_ztpl_app_crud_entity_data(ztpl_app_crud_entity_repo: ZtplAppCrudEntityRepo, index: Optional[int] = None) -> ZtplAppCrudEntity:
    ztpl_app_crud_entity_data = create_ztpl_app_crud_entity_data()
    ztpl_app_crud_entity_data.ztplAppCrudFirstField = 'ztplAppCrudEntity' if index is None else 'ztplAppCrudEntity-{index}'.format(index=index)
    ztpl_app_crud_entity_data.created_by = 'original_user'
    ztpl_app_crud_entity_data.updated_by = 'original_user'
    return ztpl_app_crud_entity_repo.insert(ztpl_app_crud_entity_data)


def create_mb() -> AppMessageBus:
    mb = AppMessageBus(LocalMessageBus())
    # handle new_activity event
    @mb.handle('new_activity')
    def handle_new_activity(activity_data):
        print('New Activity', activity_data)
    # return mb
    return mb


def init_test_ztpl_app_crud_entity_service_components() -> Tuple[ZtplAppCrudEntityService, DBZtplAppCrudEntityRepo, AppMessageBus, AppRPC]:
    engine = create_engine('sqlite://', echo=False)
    ztpl_app_crud_entity_repo = DBZtplAppCrudEntityRepo(engine=engine, create_all=True)
    mb = create_mb()
    rpc = AppRPC(LocalRPC())
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mb, rpc, ztpl_app_crud_entity_repo)
    return ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, mb, rpc
