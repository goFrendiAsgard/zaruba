from typing import List
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.ztplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from repos.ztplAppCrudEntity import ZtplAppCrudEntityRepo

import uuid
import datetime

Base = declarative_base()

class DBZtplAppCrudEntityEntity(Base):
    __tablename__ = "ztpl_app_crud_entities"
    id = Column(String(36), primary_key=True, index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow)


class DBZtplAppCrudEntityRepo(ZtplAppCrudEntityRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def find_by_id(self, id: str) -> ZtplAppCrudEntity:
        db = Session(self.engine)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            db_ztpl_app_crud_entity = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.id == id).first()
            if db_ztpl_app_crud_entity is None:
                return None
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(db_ztpl_app_crud_entity)
        finally:
            db.close()
        return ztpl_app_crud_entity

    def find(self, keyword: str, limit: int, offset: int) -> List[ZtplAppCrudEntity]:
        db = Session(self.engine)
        ztpl_app_crud_entities: List[ZtplAppCrudEntity] = []
        try:
            keyword = '%{}%'.format(keyword) if keyword != '' else '%'
            db_ztpl_app_crud_entities = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.ztplAppCrudFirstField.like(keyword)).offset(offset).limit(limit).all()
            ztpl_app_crud_entities = [ZtplAppCrudEntity.from_orm(db_result) for db_result in db_ztpl_app_crud_entities]
        finally:
            db.close()
        return ztpl_app_crud_entities

    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        db = Session(self.engine)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            new_ztpl_app_crud_entity_id = str(uuid.uuid4())
            db_ztpl_app_crud_entity = DBZtplAppCrudEntityEntity(
                id=new_ztpl_app_crud_entity_id,
                created_at=datetime.datetime.utcnow()
            )
            db.add(db_ztpl_app_crud_entity)
            db.commit()
            db.refresh(db_ztpl_app_crud_entity) 
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(db_ztpl_app_crud_entity)
        finally:
            db.close()
        return ztpl_app_crud_entity

    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        db = Session(self.engine)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            db_ztpl_app_crud_entity = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.id == id).first()
            if db_ztpl_app_crud_entity is None:
                return None
            db_ztpl_app_crud_entity.updated_at = datetime.datetime.utcnow()
            db.add(db_ztpl_app_crud_entity)
            db.commit()
            db.refresh(db_ztpl_app_crud_entity) 
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(db_ztpl_app_crud_entity)
        finally:
            db.close()
        return ztpl_app_crud_entity

    def delete(self, id: str) -> ZtplAppCrudEntity:
        db = Session(self.engine)
        ztpl_app_crud_entity: ZtplAppCrudEntity
        try:
            db_ztpl_app_crud_entity = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.id == id).first()
            if db_ztpl_app_crud_entity is None:
                return None
            db.delete(db_ztpl_app_crud_entity)
            db.commit()
            ztpl_app_crud_entity = ZtplAppCrudEntity.from_orm(db_ztpl_app_crud_entity)
        finally:
            db.close()
        return ztpl_app_crud_entity

