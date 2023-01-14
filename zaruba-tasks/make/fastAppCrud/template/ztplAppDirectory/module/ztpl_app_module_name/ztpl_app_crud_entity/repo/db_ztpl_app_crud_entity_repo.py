from typing import List, Optional
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import (
    Boolean, Column, DateTime, ForeignKey, Integer, String, Text
)
from schema.ztpl_app_crud_entity import (
    ZtplAppCrudEntity, ZtplAppCrudEntityData
)
from module.ztpl_app_module_name.ztpl_app_crud_entity.repo.ztpl_app_crud_entity_repo import (
    ZtplAppCrudEntityRepo
)
from repo import Base

import uuid
import datetime

# Note: 🤖 Don't delete the following statement


class DBZtplAppCrudEntityEntity(Base):
    __tablename__ = "ztpl_app_crud_entities"
    id = Column(String(36), primary_key=True, index=True)
    # Note: 🤖 Don't delete this line
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


# Note: 🤖 Don't delete the following statement
class DBZtplAppCrudEntityRepo(ZtplAppCrudEntityRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def find_by_id(self, id: str) -> Optional[ZtplAppCrudEntity]:
        db = Session(self.engine, expire_on_commit=False)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            db_ztpl_app_crud_entity = db.query(
                DBZtplAppCrudEntityEntity
            ).filter(
                DBZtplAppCrudEntityEntity.id == id
            ).first()
            if db_ztpl_app_crud_entity is None:
                return None
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(
                db_ztpl_app_crud_entity
            )
        finally:
            db.close()
        return ztpl_app_crud_entity

    def find(
        self,
        keyword: str,
        limit: int,
        offset: int
    ) -> List[ZtplAppCrudEntity]:
        db = Session(self.engine, expire_on_commit=False)
        ztpl_app_crud_entities: List[ZtplAppCrudEntity] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_ztpl_app_crud_entities = db.query(
                DBZtplAppCrudEntityEntity
            ).filter(
                DBZtplAppCrudEntityEntity.ztplAppCrudFirstField.like(
                    keyword_filter
                )
            ).offset(offset).limit(limit).all()
            ztpl_app_crud_entities = [ZtplAppCrudEntity.from_orm(
                db_result) for db_result in db_ztpl_app_crud_entities]
        finally:
            db.close()
        return ztpl_app_crud_entities

    def count(self, keyword: str) -> int:
        db = Session(self.engine, expire_on_commit=False)
        ztpl_app_crud_entity_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            ztpl_app_crud_entity_count = db.query(
                DBZtplAppCrudEntityEntity
            ).filter(
                DBZtplAppCrudEntityEntity.ztplAppCrudFirstField.like(
                    keyword_filter
                )
            ).count()
        finally:
            db.close()
        return ztpl_app_crud_entity_count

    # Note: 🤖 Don't delete the following statement
    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        db = Session(self.engine, expire_on_commit=False)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            new_ztpl_app_crud_entity_id = str(uuid.uuid4())
            db_ztpl_app_crud_entity = DBZtplAppCrudEntityEntity(
                id=new_ztpl_app_crud_entity_id,
                created_at=datetime.datetime.utcnow(),  # Note: 🤖 Don't delete this line
                created_by=ztpl_app_crud_entity_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=ztpl_app_crud_entity_data.updated_by,
            )
            db.add(db_ztpl_app_crud_entity)
            db.commit()
            db.refresh(db_ztpl_app_crud_entity)
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(
                db_ztpl_app_crud_entity)
        finally:
            db.close()
        return ztpl_app_crud_entity

    # Note: 🤖 Don't delete the following statement
    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> Optional[ZtplAppCrudEntity]:
        db = Session(self.engine, expire_on_commit=False)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            db_ztpl_app_crud_entity = db.query(
                DBZtplAppCrudEntityEntity
            ).filter(
                DBZtplAppCrudEntityEntity.id == id
            ).first()
            if db_ztpl_app_crud_entity is None:
                return None
            # Note: 🤖 Don't delete this line
            db_ztpl_app_crud_entity.updated_at = datetime.datetime.utcnow()
            db_ztpl_app_crud_entity.updated_by = ztpl_app_crud_entity_data.updated_by
            db.add(db_ztpl_app_crud_entity)
            db.commit()
            db.refresh(db_ztpl_app_crud_entity)
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(
                db_ztpl_app_crud_entity)
        finally:
            db.close()
        return ztpl_app_crud_entity

    def delete(self, id: str) -> Optional[ZtplAppCrudEntity]:
        db = Session(self.engine, expire_on_commit=False)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            db_ztpl_app_crud_entity = db.query(
                DBZtplAppCrudEntityEntity
            ).filter(
                DBZtplAppCrudEntityEntity.id == id
            ).first()
            if db_ztpl_app_crud_entity is None:
                return None
            db.delete(db_ztpl_app_crud_entity)
            db.commit()
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(
                db_ztpl_app_crud_entity)
        finally:
            db.close()
        return ztpl_app_crud_entity

    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'
