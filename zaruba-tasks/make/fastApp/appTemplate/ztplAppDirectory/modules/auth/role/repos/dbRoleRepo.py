from typing import Any, List, Optional
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.role import Role, RoleData
from modules.auth.role.repos.roleRepo import RoleRepo
from repos import Base

import uuid
import datetime
import json


class DBRoleEntity(Base):
    __tablename__ = "roles"
    id = Column(String(36), primary_key=True, index=True)
    name = Column(String(20), index=True)
    json_permissions = Column(String(20), index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


class DBRoleRepo(RoleRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'

    def _from_db_result(self, db_result: Any) -> Role:
        role = Role.from_orm(db_result)
        role.permissions = json.loads(db_result.json_permissions)
        return role

    def find_by_id(self, id: str) -> Optional[Role]:
        db = Session(self.engine)
        role: Role
        try:
            db_result = db.query(DBRoleEntity).filter(DBRoleEntity.id == id).first()
            if db_result is None:
                return None
            role = self._from_db_result(db_result)
        finally:
            db.close()
        return role

    def find_by_name(self, name: str) -> Optional[Role]:
        db = Session(self.engine)
        role: Role
        try:
            db_result = db.query(DBRoleEntity).filter(DBRoleEntity.name == name).first()
            if db_result is None:
                return None
            role = self._from_db_result(db_result)
        finally:
            db.close()
        return role

    def find(self, keyword: str, limit: int, offset: int) -> List[Role]:
        db = Session(self.engine)
        roles: List[Role] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_results = db.query(DBRoleEntity).filter(DBRoleEntity.name.like(keyword_filter)).offset(offset).limit(limit).all()
            roles = [self._from_db_result(db_result) for db_result in db_results]
        finally:
            db.close()
        return roles

    def count(self, keyword: str) -> int:
        db = Session(self.engine)
        role_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            role_count = db.query(DBRoleEntity).filter(DBRoleEntity.name.like(keyword_filter)).count()
        finally:
            db.close()
        return role_count

    def insert(self, role_data: RoleData) -> Optional[Role]:
        db = Session(self.engine)
        role: Role
        try:
            new_role_id=str(uuid.uuid4())
            db_role = DBRoleEntity(
                id=new_role_id,
                name=role_data.name,
                json_permissions=json.dumps(role_data.permissions),
                created_at=datetime.datetime.utcnow(),
                created_by=role_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=role_data.updated_by
            )
            db.add(db_role)
            db.commit()
            db.refresh(db_role) 
            role = self._from_db_result(db_role)
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
            db_role.updated_at = datetime.datetime.utcnow()
            db_role.updated_by = role_data.updated_by
            db.add(db_role)
            db.commit()
            db.refresh(db_role) 
            role = self._from_db_result(db_role)
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
            role = self._from_db_result(db_role)
        finally:
            db.close()
        return role

