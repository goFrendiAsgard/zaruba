from typing import Any, List, Optional
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.ext.hybrid import hybrid_property
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import and_, or_, Boolean, Column, DateTime, ForeignKey, Integer, String, Text
from schemas.user import User, UserWithoutPassword, UserData
from modules.auth.user.repos.userRepo import UserRepo
from repos import Base

import bcrypt
import uuid
import datetime
import json

class DBUserEntity(Base):
    __tablename__ = "users"
    id = Column(String(36), primary_key=True, index=True)
    username = Column(String(50), index=True, unique=True, nullable=False)
    email = Column(String(50), index=True, unique=True, nullable=True)
    phone_number = Column(String(20), index=True, unique=True, nullable=True)
    json_role_ids = Column(Text(), nullable=False, default='[]')
    json_permissions = Column(Text(), nullable=False, default='[]')
    active = Column(Boolean(), index=True, nullable=False, default=False)
    hashed_password = Column(String(60), index=False, nullable=False)
    full_name = Column(String(50), index=True, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)

    @hybrid_property
    def permissions(self) -> List[str]:
        return json.loads(self.json_permissions)


class DBUserRepo(UserRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def _hash_password(self, password: str) -> str:
        return bcrypt.hashpw(password.encode('utf-8'), bcrypt.gensalt()).decode('utf-8')

    def _is_valid_password(self, password: str, hashed_password: str) -> bool:
        return bcrypt.checkpw(password.encode('utf-8'), hashed_password.encode('utf-8'))

    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'

    def _from_db_result(self, db_result: Any) -> UserWithoutPassword:
        user = UserWithoutPassword.from_orm(db_result)
        user.permissions = json.loads(db_result.json_permissions)
        user.role_ids = json.loads(db_result.json_role_ids)
        return user

    def find_by_username(self, username: str) -> Optional[User]:
        db = Session(self.engine)
        user: User
        try:
            db_user = db.query(DBUserEntity).filter(DBUserEntity.username == username).first()
            if db_user is None:
                return None
            user = User.from_orm(db_user)
        finally:
            db.close()
        return user

    def find_by_id(self, id: str) -> Optional[UserWithoutPassword]:
        db = Session(self.engine)
        user: UserWithoutPassword
        try:
            db_user = db.query(DBUserEntity).filter(DBUserEntity.id == id).first()
            if db_user is None:
                return None
            user = self._from_db_result(db_user)
        finally:
            db.close()
        return user
    
    def find_by_identity_and_password(self, identity: str, password: str) -> Optional[UserWithoutPassword]:
        db = Session(self.engine)
        user: UserWithoutPassword
        try:
            db_user = db.query(DBUserEntity).filter(
                    or_(
                        and_(DBUserEntity.username == identity, DBUserEntity.username != '', DBUserEntity.username is not None),
                        and_(DBUserEntity.email == identity, DBUserEntity.email != '', DBUserEntity.email is not None),
                        and_(DBUserEntity.phone_number == identity, DBUserEntity.phone_number != '', DBUserEntity.phone_number is not None)
                    )
                ).first()
            if not self._is_valid_password(password, db_user.hashed_password):
                return None
            user = self._from_db_result(db_user)
        finally:
            db.close()
        return user

    def find(self, keyword: str, limit: int, offset: int) -> List[UserWithoutPassword]:
        db = Session(self.engine)
        users: List[UserWithoutPassword] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_users = db.query(DBUserEntity).filter(
                    or_(
                        DBUserEntity.username.like(keyword_filter),
                        DBUserEntity.email.like(keyword_filter),
                        DBUserEntity.phone_number.like(keyword_filter),
                        DBUserEntity.full_name.like(keyword_filter),
                    )
                ).offset(offset).limit(limit).all()
            users = [self._from_db_result(db_user) for db_user in db_users]
        finally:
            db.close()
        return users

    def count(self, keyword: str) -> int:
        db = Session(self.engine)
        user_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            user_count = db.query(DBUserEntity).filter(
                    or_(
                        DBUserEntity.username.like(keyword_filter),
                        DBUserEntity.email.like(keyword_filter),
                        DBUserEntity.phone_number.like(keyword_filter),
                        DBUserEntity.full_name.like(keyword_filter),
                    )
                ).count()
        finally:
            db.close()
        return user_count

    def insert(self, user_data: UserData) -> Optional[UserWithoutPassword]:
        db = Session(self.engine)
        new_user: UserWithoutPassword
        try:
            new_user_id=str(uuid.uuid4())
            db_user = DBUserEntity(
                id=new_user_id,
                username=user_data.username,
                email=user_data.email,
                phone_number=user_data.phone_number,
                json_permissions=json.dumps(user_data.permissions),
                json_role_ids=json.dumps(user_data.role_ids),
                active=user_data.active,
                hashed_password=self._hash_password(user_data.password),
                full_name=user_data.full_name,
                created_at=datetime.datetime.utcnow(),
                created_by=user_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=user_data.updated_by
            )
            db.add(db_user)
            db.commit()
            db.refresh(db_user) 
            new_user = self._from_db_result(db_user)
        finally:
            db.close()
        return new_user

    def update(self, id: str, user_data: UserData) -> Optional[UserWithoutPassword]:
        db = Session(self.engine)
        updated_user: UserWithoutPassword
        try:
            db_user = db.query(DBUserEntity).filter(DBUserEntity.id == id).first()
            if db_user is None:
                return None
            db_user.username = user_data.username
            db_user.email = user_data.email
            db_user.phone_number = user_data.phone_number
            db_user.json_permissions = json.dumps(user_data.permissions)
            db_user.json_role_ids = json.dumps(user_data.role_ids)
            db_user.active = user_data.active
            db_user.full_name = user_data.full_name
            db_user.updated_at = datetime.datetime.utcnow()
            db_user.updated_by = user_data.updated_by
            if user_data.password:
                db_user.hashed_password = self._hash_password(user_data.password)
            db.add(db_user)
            db.commit()
            db.refresh(db_user) 
            updated_user = self._from_db_result(db_user)
        finally:
            db.close()
        return updated_user

    def delete(self, id: str) -> Optional[UserWithoutPassword]:
        db = Session(self.engine)
        deleted_user: UserWithoutPassword
        try:
            db_user = db.query(DBUserEntity).filter(DBUserEntity.id == id).first()
            if db_user is None:
                return None
            db.delete(db_user)
            db.commit()
            deleted_user = self._from_db_result(db_user)
        finally:
            db.close()
        return deleted_user

