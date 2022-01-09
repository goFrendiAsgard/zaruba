from typing import List
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.role import Role, RoleData
from repos.role import RoleRepo

import uuid
import datetime
import json

Base = declarative_base()

class DBRoleEntity(Base):
    __tablename__ = "roles"
    id = Column(String(36), primary_key=True, index=True)
    name = Column(String(20), index=True)
    json_permissions = Column(String(20), index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow)


class DBRoleRepo(RoleRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def find_by_id(self, id: str) -> Role:
        db = Session(self.engine)
        result: Role
        try:
            db_result = db.query(DBRoleEntity).filter(DBRoleEntity.id == id).first()
            if db_result is None:
                return None
            result = Role.from_orm(db_result)
        finally:
            db.close()
        return result

    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        db = Session(self.engine)
        results: List[Role] = []
        try:
            keyword = '%{}%'.format(keyword) if keyword != '' else '%'
            db_results = db.query(DBRoleEntity).filter(DBRoleEntity.name.like(keyword)).offset(offset).limit(limit).all()
            results = [Role.from_orm(db_result) for db_result in db_results]
        finally:
            db.close()
        return results

    def insert(self, role_data: RoleData) -> Role:
        db = Session(self.engine)
        result: Role
        try:
            db_entity = DBRoleEntity(
                id=str(uuid.uuid4()),
                name=role_data.name,
                json_permissions=json.dumps(role_data.permissions),
                created_at=datetime.datetime.utcnow()
            )
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = Role.from_orm(db_entity)
        finally:
            db.close()
        return result

    def update(self, id: str, role_data: RoleData) -> Role:
        db = Session(self.engine)
        result: Role
        try:
            db_entity = db.query(DBRoleEntity).filter(DBRoleEntity.id == id).first()
            if db_entity is None:
                return None
            db_entity.name = role_data.name
            db_entity.json_permissions = json.dumps(role_data.permissions)
            db_entity.updated_at = datetime.datetime.utcnow()
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = Role.from_orm(db_entity)
        finally:
            db.close()
        return result

    def delete(self, id: str) -> Role:
        db = Session(self.engine)
        result: Role
        try:
            db_entity = db.query(DBRoleEntity).filter(DBRoleEntity.id == id).first()
            if db_entity is None:
                return None
            db.delete(db_entity)
            db.commit()
            result = Role.from_orm(db_entity)
        finally:
            db.close()
        return result

