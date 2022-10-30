from typing import Optional, Tuple
from modules.ztplAppModuleName.ztplAppCrudEntity.repos.ztplAppCrudEntityRepo import ZtplAppCrudEntityRepo
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from modules.ztplAppModuleName.ztplAppCrudEntity.ztplAppCrudEntityService import ZtplAppCrudEntityService
from modules.ztplAppModuleName.ztplAppCrudEntity.repos.dbZtplAppCrudEntityRepo import DBZtplAppCrudEntityRepo
from helpers.transport import LocalRPC, LocalMessageBus
from sqlalchemy import create_engine

def create_ztpl_app_crud_entity_data():
    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
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


def init_test_ztpl_app_crud_entity_service_components() -> Tuple[ZtplAppCrudEntityService, DBZtplAppCrudEntityRepo, LocalMessageBus, LocalRPC]:
    engine = create_engine('sqlite://', echo=False)
    ztpl_app_crud_entity_repo = DBZtplAppCrudEntityRepo(engine=engine, create_all=True)
    mb = LocalMessageBus()
    rpc = LocalRPC()
    ztpl_app_crud_entity_service = ZtplAppCrudEntityService(mb, rpc, ztpl_app_crud_entity_repo)
    return ztpl_app_crud_entity_service, ztpl_app_crud_entity_repo, mb, rpc
