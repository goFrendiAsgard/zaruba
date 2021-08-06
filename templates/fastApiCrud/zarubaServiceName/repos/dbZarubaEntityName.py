from typing import List
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.zarubaEntityName import ZarubaEntityName, ZarubaEntityNameData
from repos.zarubaEntityName import ZarubaEntityNameRepo

import uuid
import datetime

Base = declarative_base()

class DBZarubaEntityNameEntity(Base):
    __tablename__ = "zaruba_entity_name"
    id = Column(String(36), primary_key=True, index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow)


class DBZarubaEntityNameRepo(ZarubaEntityNameRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)


    def find_by_id(self, id: str) -> ZarubaEntityName:
        db = Session(self.engine)
        result: ZarubaEntityName
        try:
            db_result = db.query(DBZarubaEntityNameEntity).filter(DBZarubaEntityNameEntity.id == id).first()
            if db_result is None:
                return None
            result = ZarubaEntityName.from_orm(db_result)
        finally:
            db.close()
        return result

    
    def find(self, keyword: str, limit: int, offset: int) -> List[ZarubaEntityName]:
        db = Session(self.engine)
        results: List[ZarubaEntityName] = []
        try:
            keyword = '%{}%'.format(keyword) if keyword != '' else '%'
            db_results = db.query(DBZarubaEntityNameEntity).filter(DBZarubaEntityNameEntity.zaruba_first_field_name.like(keyword)).offset(offset).limit(limit).all()
            results = [ZarubaEntityName.from_orm(db_result) for db_result in db_results]
        finally:
            db.close()
        return results

    
    def insert(self, zaruba_entity_name_data: ZarubaEntityNameData) -> ZarubaEntityName:
        db = Session(self.engine)
        result: ZarubaEntityName
        try:
            db_entity = DBZarubaEntityNameEntity(
                id=str(uuid.uuid4()),
                created_at=datetime.datetime.utcnow()
            )
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = ZarubaEntityName.from_orm(db_entity)
        finally:
            db.close()
        return result
    
    def update(self, id: str, zaruba_entity_name_data: ZarubaEntityNameData) -> ZarubaEntityName:
        db = Session(self.engine)
        result: ZarubaEntityName
        try:
            db_entity = db.query(DBZarubaEntityNameEntity).filter(DBZarubaEntityNameEntity.id == id).first()
            if db_entity is None:
                return None
            db_entity.updated_at = datetime.datetime.utcnow()
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = ZarubaEntityName.from_orm(db_entity)
        finally:
            db.close()
        return result

 
    def delete(self, id: str) -> ZarubaEntityName:
        db = Session(self.engine)
        result: ZarubaEntityName
        try:
            db_entity = db.query(DBZarubaEntityNameEntity).filter(DBZarubaEntityNameEntity.id == id).first()
            if db_entity is None:
                return None
            db.delete(db_entity)
            db.commit()
            result = ZarubaEntityName.from_orm(db_entity)
        finally:
            db.close()
        return result

