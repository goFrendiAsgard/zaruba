from typing import List, Optional
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String, Text
from schemas.content import Content, ContentData
from modules.cms.content.repos.contentRepo import ContentRepo
from repos import Base

import uuid
import datetime

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class DBContentEntity(Base):
    __tablename__ = "contents"
    id = Column(String(36), primary_key=True, index=True)
    title = Column(String(255), index=True)
    type_id = Column(String(36), index=True)
    description = Column(Text(), index=True, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow) # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


class DBContentAttributeEntity(Base):
    __tablename__ = "content_attributes"
    id = Column(String(36), primary_key=True, index=True)
    content_id = Column(String(36), index=True)
    key = Column(String(255), index=True)
    value = Column(Text(), index=True, nullable=True)
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
        db = Session(self.engine)
        content: Content
        try:
            db_content = db.query(DBContentEntity).filter(DBContentEntity.id == id).first()
            if db_content is None:
                return None
            content = Content.from_orm(db_content)
        finally:
            db.close()
        return content


    def find(self, keyword: str, limit: int, offset: int) -> List[Content]:
        db = Session(self.engine)
        contents: List[Content] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_contents = db.query(DBContentEntity).filter(DBContentEntity.title.like(keyword_filter)).offset(offset).limit(limit).all()
            contents = [Content.from_orm(db_result) for db_result in db_contents]
        finally:
            db.close()
        return contents


    def count(self, keyword: str) -> int:
        db = Session(self.engine)
        content_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            content_count = db.query(DBContentEntity).filter(DBContentEntity.title.like(keyword_filter)).count()
        finally:
            db.close()
        return content_count


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def insert(self, content_data: ContentData) -> Optional[Content]:
        db = Session(self.engine)
        content: Content
        try:
            new_content_id = str(uuid.uuid4())
            db_content = DBContentEntity(
                id=new_content_id,
                title=content_data.title,
                type_id=content_data.type_id,
                description=content_data.description,
                created_at=datetime.datetime.utcnow(), # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
                created_by=content_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=content_data.updated_by,
            )
            db.add(db_content)
            db.commit()
            db.refresh(db_content) 
            content = Content.from_orm(db_content)
        finally:
            db.close()
        return content


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def update(self, id: str, content_data: ContentData) -> Optional[Content]:
        db = Session(self.engine)
        content: Content
        try:
            db_content = db.query(DBContentEntity).filter(DBContentEntity.id == id).first()
            if db_content is None:
                return None
            db_content.title = content_data.title
            db_content.type_id = content_data.type_id
            db_content.description = content_data.description
            db_content.updated_at = datetime.datetime.utcnow() # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
            db_content.updated_by = content_data.updated_by
            db.add(db_content)
            db.commit()
            db.refresh(db_content) 
            content = Content.from_orm(db_content)
        finally:
            db.close()
        return content


    def delete(self, id: str) -> Optional[Content]:
        db = Session(self.engine)
        content: Content
        try:
            db_content = db.query(DBContentEntity).filter(DBContentEntity.id == id).first()
            if db_content is None:
                return None
            db.delete(db_content)
            db.commit()
            content = Content.from_orm(db_content)
        finally:
            db.close()
        return content


    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'
