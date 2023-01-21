from typing import Any, List, Optional, Mapping
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session, relationship
from sqlalchemy import (
    or_, Column, DateTime, ForeignKey, String, Text
)
from schema.content import Content, ContentData
from module.cms.content.repo.content_repo import ContentRepo
from repo import Base, BaseMixin, DBRepo


# Note: 🤖 Don't delete the following statement
class DBContentEntity(Base, BaseMixin):
    __tablename__ = "contents"  # Note: 🤖 Don't delete this line
    title = Column(String(255), index=True)
    content_type_id = Column(
        String(36), ForeignKey('content_types.id'), index=True
    )
    description = Column(Text(), index=False, nullable=True)
    content_attributes = relationship(
        'DBContentAttributeEntity',
        back_populates='content',
        cascade='all, delete-orphan'
    )


class DBContentAttributeEntity(Base, BaseMixin):
    __tablename__ = "content_attributes"
    content_id = Column(
        String(36), ForeignKey('contents.id'), index=True, nullable=False
    )
    key = Column(String(255), index=True, nullable=False)
    value = Column(Text(), index=False, nullable=True)
    content = relationship(
        'DBContentEntity', back_populates='content_attributes'
    )


class DBContentRepo(
    DBRepo[DBContentEntity, Content, ContentData],
    ContentRepo
):

    schema_class = Content
    db_entity_class = DBContentEntity

    def get_search_filter(self, db: Session, keyword: str) -> Any:
        like_keyword = '%{}%'.format(keyword) if keyword != '' else '%'
        subquery = db.query(DBContentAttributeEntity.content_id).filter(
                DBContentAttributeEntity.value.like(like_keyword)
            ).subquery()
        return or_(
            DBContentEntity.title.like(like_keyword),
            DBContentEntity.id.in_(subquery)
        )

    def from_schema_data_to_db_entity_dict(
        self, content_data: ContentData
    ) -> Mapping[str, Any]:
        content_dict = super().from_schema_data_to_db_entity_dict(content_data)
        content_dict['content_attributes'] = self.from_dict_to_attributes(
            content_data.attributes
        )
        return content_dict

    def from_db_entity_to_schema(self, db_content: DBContentEntity) -> Content:
        content = super().from_db_entity_to_schema(db_content)
        content.attributes = self.from_attributes_to_dict(
            db_content.content_attributes
        )
        return content

    def from_dict_to_attributes(
        self, attribute_map: Mapping[str, str]
    ) -> List[DBContentAttributeEntity]:
        attributes: List[DBContentAttributeEntity] = []
        for key, value in attribute_map.items():
            attributes.append(DBContentAttributeEntity(key=key, value=value))
        return attributes

    def from_attributes_to_dict(
        self, attributes: List[DBContentAttributeEntity]
    ) -> Mapping[str, str]:
        attribute_map: Mapping[str, str] = {}
        for attribute in attributes:
            attribute_map[attribute.key] = attribute.value
        return attribute_map