from typing import List, Optional
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
    created_at = Column(DateTime, default=datetime.datetime.now)
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, default=datetime.datetime.now)
    updated_by = Column(String(36), nullable=True)


class DBRoleRepo(RoleRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def find_by_id(self, id: str) -> Optional[Role]:
        db = Session(self.engine)
        role: Role
        try:
            db_result = db.query(DBRoleEntity).filter(DBRoleEntity.id == id).first()
            if db_result is None:
                return None
            role = Role.from_orm(db_result)
        finally:
            db.close()
        return role

    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        db = Session(self.engine)
        roles: List[Role] = []
        try:
            keyword = '%{}%'.format(keyword) if keyword != '' else '%'
            db_results = db.query(DBRoleEntity).filter(DBRoleEntity.name.like(keyword)).offset(offset).limit(limit).all()
            roles = [Role.from_orm(db_result) for db_result in db_results]
        finally:
            db.close()
        return roles

    def insert(self, role_data: RoleData) -> Optional[Role]:
        db = Session(self.engine)
        role: Role
        try:
            new_role_id=str(uuid.uuid4())
            db_role = DBRoleEntity(
                id=new_role_id,
                name=role_data.name,
                json_permissions=json.dumps(role_data.permissions),
                created_at=datetime.datetime.now()
            )
            db.add(db_role)
            db.commit()
            db.refresh(db_role) 
            role = Role.from_orm(db_role)
        finally:
            db.close()
        return role

    def update(self, id: str, role_data: RoleData) -> Optional[Role]:
        db = Session(self.engine)
        role: Role
        try:
            db_role = db.query(DBRoleEntity).filter(DBRoleEntity.id == id).first()
            if db_role is None:
                return None
            db_role.name = role_data.name
            db_role.json_permissions = json.dumps(role_data.permissions)
            db_role.updated_at = datetime.datetime.now()
            db.add(db_role)
            db.commit()
            db.refresh(db_role) 
            role = Role.from_orm(db_role)
        finally:
            db.close()
        return role

    def delete(self, id: str) -> Optional[Role]:
        db = Session(self.engine)
        role: Role
        try:
            db_role = db.query(DBRoleEntity).filter(DBRoleEntity.id == id).first()
            if db_role is None:
                return None
            db.delete(db_role)
            db.commit()
            role = Role.from_orm(db_role)
        finally:
            db.close()
        return role

