from typing import List
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.ZtplAppCrudEntity import ZtplAppCrudEntity, ZtplAppCrudEntityData
from repos.ZtplAppCrudEntity import ZtplAppCrudEntityRepo

import uuid
import datetime

Base = declarative_base()

class DBZtplAppCrudEntityEntity(Base):
    __tablename__ = "ztpl_app_crud_entity"
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
        result: ZtplAppCrudEntity
        try:
            db_result = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.id == id).first()
            if db_result is None:
                return None
            result = ZtplAppCrudEntity.from_orm(db_result)
        finally:
            db.close()
        return result

    
    def find(self, keyword: str, limit: int, offset: int) -> List[ZtplAppCrudEntity]:
        db = Session(self.engine)
        results: List[ZtplAppCrudEntity] = []
        try:
            keyword = '%{}%'.format(keyword) if keyword != '' else '%'
            db_results = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.zaruba_first_field_name.like(keyword)).offset(offset).limit(limit).all()
            results = [ZtplAppCrudEntity.from_orm(db_result) for db_result in db_results]
        finally:
            db.close()
        return results

    
    def insert(self, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        db = Session(self.engine)
        result: ZtplAppCrudEntity
        try:
            db_entity = DBZtplAppCrudEntityEntity(
                id=str(uuid.uuid4()),
                created_at=datetime.datetime.utcnow()
            )
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = ZtplAppCrudEntity.from_orm(db_entity)
        finally:
            db.close()
        return result
    
    def update(self, id: str, ztpl_app_crud_entity_data: ZtplAppCrudEntityData) -> ZtplAppCrudEntity:
        db = Session(self.engine)
        result: ZtplAppCrudEntity
        try:
            db_entity = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.id == id).first()
            if db_entity is None:
                return None
            db_entity.updated_at = datetime.datetime.utcnow()
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = ZtplAppCrudEntity.from_orm(db_entity)
        finally:
            db.close()
        return result

 
    def delete(self, id: str) -> ZtplAppCrudEntity:
        db = Session(self.engine)
        result: ZtplAppCrudEntity
        try:
            db_entity = db.query(DBZtplAppCrudEntityEntity).filter(DBZtplAppCrudEntityEntity.id == id).first()
            if db_entity is None:
                return None
            db.delete(db_entity)
            db.commit()
            result = ZtplAppCrudEntity.from_orm(db_entity)
        finally:
            db.close()
        return result

