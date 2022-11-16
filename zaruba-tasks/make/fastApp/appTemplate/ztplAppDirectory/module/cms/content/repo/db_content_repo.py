from typing import List, Optional, Mapping
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session, relationship
from sqlalchemy import or_, Boolean, Column, DateTime, ForeignKey, Integer, String, Text
from schema.content import Content, ContentData
from module.cms.content.repo.content_repo import ContentRepo
from repo import Base

import uuid
import datetime


def generate_primary_key() -> str:
    return str(uuid.uuid4())


# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class DBContentEntity(Base):
    __tablename__ = "contents"
    id = Column(String(36), primary_key=True, index=True)
    title = Column(String(255), index=True)
    content_type_id = Column(String(36), ForeignKey('content_types.id'), index=True)
    description = Column(Text(), index=True, nullable=True)
    content_attributes = relationship('DBContentAttributeEntity', back_populates='content', cascade='all, delete-orphan')
    created_at = Column(DateTime, default=datetime.datetime.utcnow) # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


class DBContentAttributeEntity(Base):
    __tablename__ = "content_attributes"
    id = Column(String(36), primary_key=True, index=True, default=generate_primary_key)
    content_id = Column(String(36), ForeignKey('contents.id'), index=True, nullable=False)
    key = Column(String(255), index=True, nullable=False)
    value = Column(Text(), index=True, nullable=True)
    content = relationship('DBContentEntity', back_populates='content_attributes')
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class DBContentRepo(ContentRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)


    def find_by_id(self, id: str) -> Optional[Content]:
        db = Session(self.engine, expire_on_commit=False)
        content: Content
        try:
            db_content = db.query(DBContentEntity).filter(DBContentEntity.id == id).first()
            if db_content is None:
                return None
            content = self._from_db_content(db_content)
        finally:
            db.close()
        return content


    def find(self, keyword: str, limit: int, offset: int) -> List[Content]:
        db = Session(self.engine, expire_on_commit=False)
        contents: List[Content] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            subquery = db.query(DBContentAttributeEntity.content_id).filter(
                DBContentAttributeEntity.value.like(keyword_filter)
            ).subquery()
            db_contents = db.query(DBContentEntity).filter(
                or_(
                    DBContentEntity.title.like(keyword_filter),
                    DBContentEntity.id.in_(subquery)
                )
            ).offset(offset).limit(limit).all()
            contents = [self._from_db_content(db_content) for db_content in db_contents]
        finally:
            db.close()
        return contents


    def count(self, keyword: str) -> int:
        db = Session(self.engine, expire_on_commit=False)
        content_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            subquery = db.query(DBContentAttributeEntity.content_id).filter(
                DBContentAttributeEntity.value.like(keyword_filter)
            ).subquery()
            content_count = db.query(DBContentEntity).filter(
                or_(
                    DBContentEntity.title.like(keyword_filter),
                    DBContentEntity.id.in_(subquery)
                )
            ).count()
        finally:
            db.close()
        return content_count


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def insert(self, content_data: ContentData) -> Optional[Content]:
        db = Session(self.engine, expire_on_commit=False)
        content: Content
        try:
            content_attributes = self._map_to_attributes(content_data.attributes)
            new_content_id = str(uuid.uuid4())
            db_content = DBContentEntity(
                id=new_content_id,
                title=content_data.title,
                content_attributes=content_attributes,
                content_type_id=content_data.content_type_id,
                description=content_data.description,
                created_at=datetime.datetime.utcnow(), # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
                created_by=content_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=content_data.updated_by,
            )
            db.add(db_content)
            db.commit()
            db.refresh(db_content) 
            content = self._from_db_content(db_content)
        finally:
            db.close()
        return content


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def update(self, id: str, content_data: ContentData) -> Optional[Content]:
        db = Session(self.engine, expire_on_commit=False)
        content: Content
        try:
            db_content = db.query(DBContentEntity).filter(DBContentEntity.id == id).first()
            if db_content is None:
                return None
            content_attributes = self._map_to_attributes(content_data.attributes)
            db_content.title = content_data.title
            db_content.content_attributes = content_attributes
            db_content.content_type_id = content_data.content_type_id
            db_content.description = content_data.description
            db_content.updated_at = datetime.datetime.utcnow() # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
            db_content.updated_by = content_data.updated_by
            db.add(db_content)
            db.commit()
            db.refresh(db_content) 
            content = self._from_db_content(db_content)
        finally:
            db.close()
        return content


    def delete(self, id: str) -> Optional[Content]:
        db = Session(self.engine, expire_on_commit=False)
        content: Content
        try:
            db_content = db.query(DBContentEntity).filter(DBContentEntity.id == id).first()
            if db_content is None:
                return None
            db.delete(db_content)
            db.commit()
            content = self._from_db_content(db_content)
        finally:
            db.close()
        return content


    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'


    def _from_db_content(self, db_content: DBContentEntity) -> Content:
        content = Content.from_orm(db_content)
        content.attributes = self._attributes_to_map(db_content.content_attributes)
        return content

 
    def _map_to_attributes(self, attribute_map: Mapping[str, str]) -> List[DBContentAttributeEntity]:
        attributes: List[DBContentAttributeEntity] = []
        for key, value in attribute_map.items():
            attributes.append(DBContentAttributeEntity(key=key, value=value))
        return attributes


    def _attributes_to_map(self, attributes: List[DBContentAttributeEntity]) -> Mapping[str, str]:
        attribute_map: Mapping[str, str] = {}
        for attribute in attributes:
            attribute_map[attribute.key] = attribute.value
        return attribute_map
