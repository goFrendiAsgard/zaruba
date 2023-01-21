from typing import Any, List, Mapping, Optional
from sqlalchemy.engine import Engine
from sqlalchemy.orm import InstrumentedAttribute
from sqlalchemy import (
    Boolean, Column, DateTime, ForeignKey, Integer, String, Text
)
from schema.ztpl_app_crud_entity import (
    ZtplAppCrudEntity, ZtplAppCrudEntityData
)
from module.ztpl_app_module_name.ztpl_app_crud_entity.repo.ztpl_app_crud_entity_repo import (
    ZtplAppCrudEntityRepo
)
from repo import Base, BaseMixin, DBRepo


# Note: 🤖 Don't delete the following statement
class DBZtplAppCrudEntityEntity(Base, BaseMixin):
    __tablename__ = "ztpl_app_crud_entities"  # Note: 🤖 Don't delete this line
    dummy_column = Column(String(36)) # This will be replaced


class DBZtplAppCrudEntityRepo(
    DBRepo[DBZtplAppCrudEntityEntity, ZtplAppCrudEntity, ZtplAppCrudEntityData],
    ZtplAppCrudEntityRepo
):
    schema_class = ZtplAppCrudEntity
    db_entity_class = DBZtplAppCrudEntityEntity

    def get_keyword_fields(self) -> List[InstrumentedAttribute]:
        '''
        Return list of fields for keyword filtering
        '''
        return [
            DBZtplAppCrudEntityEntity.ztplAppCrudFirstField
        ]
   
    def from_schema_data_to_db_entity_dict(
        self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData
    ) -> Mapping[str, Any]:
        '''
        Convert ZtplAppCrudEntityData into dictionary
        The result of this convertion is used for inserting/updating DBZtplAppCrudEntityEntity.
        '''
        ztpl_app_crud_entity_dict = super().from_schema_data_to_db_entity_dict(
            ztpl_app_crud_entity_data
        )
        return ztpl_app_crud_entity_dict

    def from_db_entity_to_schema(
        self, db_ztpl_app_crud_entity: DBZtplAppCrudEntityEntity
    ) -> ZtplAppCrudEntity:
        '''
        Convert DBZtplAppCrudEntityEntity into ZtplAppCrudEntity
        The result of this convertion is usually returned to external layer (e.g., service)
        '''
        ztpl_app_crud_entity = super().from_db_entity_to_schema(db_ztpl_app_crud_entity)
        return ztpl_app_crud_entity
