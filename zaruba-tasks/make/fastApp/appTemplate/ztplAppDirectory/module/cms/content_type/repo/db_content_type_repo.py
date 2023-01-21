from typing import Any, List, Mapping, Optional
from sqlalchemy.engine import Engine
from sqlalchemy.orm import InstrumentedAttribute
from sqlalchemy import (
    Boolean, Column, DateTime, ForeignKey, Integer, String, Text
)
from schema.content_type import (
    ContentType, ContentTypeData, ContentTypeAttribute
)
from module.cms.content_type.repo.content_type_repo import ContentTypeRepo
from repo import Base, BaseMixin, DBRepo

import jsons


# Note: 🤖 Don't delete the following statement
class DBContentTypeEntity(Base, BaseMixin):
    __tablename__ = "content_types"  # Note: 🤖 Don't delete this line
    name = Column(String(255), index=True)
    template = Column(String(255), index=True)
    json_attributes = Column(Text(), index=False, default='[]')


# Note: 🤖 Don't delete the following statement
class DBContentTypeRepo(
    DBRepo[DBContentTypeEntity, ContentType, ContentTypeData],
    ContentTypeRepo
):

    schema_class = ContentType
    db_entity_class = DBContentTypeEntity

    def find_by_name(self, name: str) -> Optional[ContentType]:
        db = self.create_db_sesion()
        try:
            search_filter = DBContentTypeEntity.name == name
            return self.fetch_one_by_filter(db, search_filter)
        finally:
            db.close()

    def get_keyword_fields(self) -> List[InstrumentedAttribute]:
        return [
            DBContentTypeEntity.name,
        ]

    def from_db_entity_to_schema(
        self, db_content_type: DBContentTypeEntity
    ) -> ContentType:
        content_type = super().from_db_entity_to_schema(db_content_type)
        content_type.attributes = [
            ContentTypeAttribute(**attribute)
            for attribute in jsons.loads(db_content_type.json_attributes)
        ]
        return content_type

    def from_schema_data_to_db_entity_dict(
        self, content_type_data: ContentTypeData
    ) -> Mapping[str, Any]:
        content_type_dict = super().from_schema_data_to_db_entity_dict(
            content_type_data
        )
        content_type_dict['json_attributes'] = jsons.dumps([
            attribute.dict() for attribute in content_type_data.attributes
        ])
        return content_type_dict