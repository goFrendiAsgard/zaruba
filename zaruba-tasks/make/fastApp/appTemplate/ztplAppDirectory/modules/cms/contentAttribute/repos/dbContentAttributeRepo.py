from typing import List, Optional
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String, Text
from schemas.contentAttribute import ContentAttribute, ContentAttributeData
from modules.cms.contentAttribute.repos.contentAttributeRepo import ContentAttributeRepo
from repos import Base

import uuid
import datetime

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class DBContentAttributeEntity(Base):
    __tablename__ = "content_attributes"
    id = Column(String(36), primary_key=True, index=True)
    content_id = Column(String(36), index=True)
    key = Column(String(255), index=True)
    value = Column(Text(), index=True, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow) # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class DBContentAttributeRepo(ContentAttributeRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)


    def find_by_id(self, id: str) -> Optional[ContentAttribute]:
        db = Session(self.engine)
        content_attribute: ContentAttribute
        try:
            db_content_attribute = db.query(DBContentAttributeEntity).filter(DBContentAttributeEntity.id == id).first()
            if db_content_attribute is None:
                return None
            content_attribute = ContentAttribute.from_orm(db_content_attribute)
        finally:
            db.close()
        return content_attribute


    def find(self, keyword: str, limit: int, offset: int) -> List[ContentAttribute]:
        db = Session(self.engine)
        content_attributes: List[ContentAttribute] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_content_attributes = db.query(DBContentAttributeEntity).filter(DBContentAttributeEntity.content_id.like(keyword_filter)).offset(offset).limit(limit).all()
            content_attributes = [ContentAttribute.from_orm(db_result) for db_result in db_content_attributes]
        finally:
            db.close()
        return content_attributes


    def count(self, keyword: str) -> int:
        db = Session(self.engine)
        content_attribute_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            content_attribute_count = db.query(DBContentAttributeEntity).filter(DBContentAttributeEntity.content_id.like(keyword_filter)).count()
        finally:
            db.close()
        return content_attribute_count


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def insert(self, content_attribute_data: ContentAttributeData) -> Optional[ContentAttribute]:
        db = Session(self.engine)
        content_attribute: ContentAttribute
        try:
            new_content_attribute_id = str(uuid.uuid4())
            db_content_attribute = DBContentAttributeEntity(
                id=new_content_attribute_id,
                content_id=content_attribute_data.content_id,
                key=content_attribute_data.key,
                value=content_attribute_data.value,
                created_at=datetime.datetime.utcnow(), # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
                created_by=content_attribute_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=content_attribute_data.updated_by,
            )
            db.add(db_content_attribute)
            db.commit()
            db.refresh(db_content_attribute) 
            content_attribute = ContentAttribute.from_orm(db_content_attribute)
        finally:
            db.close()
        return content_attribute


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def update(self, id: str, content_attribute_data: ContentAttributeData) -> Optional[ContentAttribute]:
        db = Session(self.engine)
        content_attribute: ContentAttribute
        try:
            db_content_attribute = db.query(DBContentAttributeEntity).filter(DBContentAttributeEntity.id == id).first()
            if db_content_attribute is None:
                return None
            db_content_attribute.content_id = content_attribute_data.content_id
            db_content_attribute.key = content_attribute_data.key
            db_content_attribute.value = content_attribute_data.value
            db_content_attribute.updated_at = datetime.datetime.utcnow() # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
            db_content_attribute.updated_by = content_attribute_data.updated_by
            db.add(db_content_attribute)
            db.commit()
            db.refresh(db_content_attribute) 
            content_attribute = ContentAttribute.from_orm(db_content_attribute)
        finally:
            db.close()
        return content_attribute


    def delete(self, id: str) -> Optional[ContentAttribute]:
        db = Session(self.engine)
        content_attribute: ContentAttribute
        try:
            db_content_attribute = db.query(DBContentAttributeEntity).filter(DBContentAttributeEntity.id == id).first()
            if db_content_attribute is None:
                return None
            db.delete(db_content_attribute)
            db.commit()
            content_attribute = ContentAttribute.from_orm(db_content_attribute)
        finally:
            db.close()
        return content_attribute


    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'
