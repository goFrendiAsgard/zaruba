from typing import List, Optional
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String, Text
from schema.content_type import ContentType, ContentTypeData, ContentTypeAttribute
from module.cms.content_type.repo.content_type_repo import ContentTypeRepo
from repo import Base

import uuid
import jsons
import datetime

# Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
class DBContentTypeEntity(Base):
    __tablename__ = "content_types"
    id = Column(String(36), primary_key=True, index=True)
    name = Column(String(255), index=True)
    template = Column(String(255), index=True)
    json_attributes = Column(Text(), index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow) # Note: 💀 Don't delete this line; Zaruba uses it for pattern matching
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


# Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
class DBContentTypeRepo(ContentTypeRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def find_by_name(self, name: str) -> Optional[ContentType]:
        db = Session(self.engine, expire_on_commit=False)
        content_type: ContentType
        try:
            db_content_type = db.query(DBContentTypeEntity).filter(DBContentTypeEntity.name == name).first()
            if db_content_type is None:
                return None
            content_type = ContentType.from_orm(db_content_type)
        finally:
            db.close()
        return content_type


    def find_by_id(self, id: str) -> Optional[ContentType]:
        db = Session(self.engine, expire_on_commit=False)
        content_type: ContentType
        try:
            db_content_type = db.query(DBContentTypeEntity).filter(DBContentTypeEntity.id == id).first()
            if db_content_type is None:
                return None
            content_type = self._from_db_content_type(db_content_type)
        finally:
            db.close()
        return content_type


    def find(self, keyword: str, limit: int, offset: int) -> List[ContentType]:
        db = Session(self.engine, expire_on_commit=False)
        content_types: List[ContentType] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_content_types = db.query(DBContentTypeEntity).filter(DBContentTypeEntity.name.like(keyword_filter)).offset(offset).limit(limit).all()
            content_types = [self._from_db_content_type(db_content_type) for db_content_type in db_content_types]
        finally:
            db.close()
        return content_types


    def count(self, keyword: str) -> int:
        db = Session(self.engine, expire_on_commit=False)
        content_type_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            content_type_count = db.query(DBContentTypeEntity).filter(DBContentTypeEntity.name.like(keyword_filter)).count()
        finally:
            db.close()
        return content_type_count


    # Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
    def insert(self, content_type_data: ContentTypeData) -> Optional[ContentType]:
        db = Session(self.engine, expire_on_commit=False)
        content_type: ContentType
        try:
            new_content_type_id = str(uuid.uuid4())
            json_attributes = jsons.dumps([attribute.dict() for attribute in content_type_data.attributes])
            db_content_type = DBContentTypeEntity(
                id=new_content_type_id,
                name=content_type_data.name,
                template=content_type_data.template,
                json_attributes=json_attributes,
                created_at=datetime.datetime.utcnow(), # Note: 💀 Don't delete this line; Zaruba uses it for pattern matching
                created_by=content_type_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=content_type_data.updated_by,
            )
            db.add(db_content_type)
            db.commit()
            db.refresh(db_content_type) 
            content_type = self._from_db_content_type(db_content_type)
        finally:
            db.close()
        return content_type


    # Note: 💀 Don't delete the following line; Zaruba uses it for pattern matching
    def update(self, id: str, content_type_data: ContentTypeData) -> Optional[ContentType]:
        db = Session(self.engine, expire_on_commit=False)
        content_type: ContentType
        try:
            db_content_type = db.query(DBContentTypeEntity).filter(DBContentTypeEntity.id == id).first()
            if db_content_type is None:
                return None
            json_attributes=jsons.dumps([attribute.dict() for attribute in content_type_data.attributes])
            db_content_type.name = content_type_data.name
            db_content_type.template = content_type_data.template
            db_content_type.json_attributes = json_attributes
            db_content_type.updated_at = datetime.datetime.utcnow() # Note: 💀 Don't delete this line; Zaruba uses it for pattern matching
            db_content_type.updated_by = content_type_data.updated_by
            db.add(db_content_type)
            db.commit()
            db.refresh(db_content_type) 
            content_type = self._from_db_content_type(db_content_type)
        finally:
            db.close()
        return content_type


    def delete(self, id: str) -> Optional[ContentType]:
        db = Session(self.engine, expire_on_commit=False)
        content_type: ContentType
        try:
            db_content_type = db.query(DBContentTypeEntity).filter(DBContentTypeEntity.id == id).first()
            if db_content_type is None:
                return None
            db.delete(db_content_type)
            db.commit()
            content_type = self._from_db_content_type(db_content_type)
        finally:
            db.close()
        return content_type


    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'


    def _from_db_content_type(self, db_content_type: DBContentTypeEntity) -> ContentType:
        content_type = ContentType.from_orm(db_content_type)
        content_type.attributes = [ContentTypeAttribute(**attribute) for attribute in jsons.loads(db_content_type.json_attributes)]
        return content_type

