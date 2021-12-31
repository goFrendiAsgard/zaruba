from typing import List
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import or_, Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.user import User, UserData
from repos.user import UserRepo

import bcrypt
import uuid
import datetime

Base = declarative_base()

class DBUserEntity(Base):
    __tablename__ = "user"
    id = Column(String(36), primary_key=True, index=True)
    username = Column(String(50), index=True, unique=True, nullable=False)
    email = Column(String(50), index=True, unique=True, nullable=False)
    roles = Column(String(20), index=True, nullable=False)
    active = Column(Boolean(), index=True, nullable=False, default=False)
    hashed_password = Column(String(20), index=False)
    full_name = Column(String(20), index=True, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow)


class DBUserRepo(UserRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)

    def _hash_password(self, password: str) -> str:
        return bcrypt.hashpw(password.encode(), bcrypt.gensalt())

    def _is_valid_password(self, password: str, hashed_password: str) -> bool:
        return bcrypt.checkpw(password.encode(), hashed_password)

    def find_by_username(self, username: str) -> User:
        db = Session(self.engine)
        user: User
        try:
            db_result = db.query(DBUserEntity).filter(DBUserEntity.username == username).first()
            if db_result is None:
                return None
            user = User.from_orm(db_result)
        finally:
            db.close()
        return user
    
    def find_by_password(self, identity: str, password: str) -> User:
        db = Session(self.engine)
        user: User
        try:
            db_user = db.query(DBUserEntity).filter(
                    or_(
                        DBUserEntity.username == identity,
                        DBUserEntity.email == identity,
                    )
                ).first()
            if not self._is_valid_password(password, db_user.hashed_password):
                return None
            user = User.from_orm(db_user)
        finally:
            db.close()
        return user

    def find_by_id(self, id: str) -> User:
        db = Session(self.engine)
        user: User
        try:
            db_result = db.query(DBUserEntity).filter(DBUserEntity.id == id).first()
            if db_result is None:
                return None
            user = User.from_orm(db_result)
        finally:
            db.close()
        return user
    
    def find_by_password(self, identity: str, password: str) -> User:
        db = Session(self.engine)
        user: User
        try:
            db_user = db.query(DBUserEntity).filter(
                    or_(
                        DBUserEntity.username == identity,
                        DBUserEntity.email == identity,
                    )
                ).first()
            if not self._is_valid_password(password, db_user.hashed_password):
                return None
            user = User.from_orm(db_user)
        finally:
            db.close()
        return user

    def find(self, keyword: str, limit: int, offset: int) -> List[User]:
        db = Session(self.engine)
        users: List[User] = []
        try:
            keyword = '%{}%'.format(keyword) if keyword != '' else '%'
            db_users = db.query(DBUserEntity).filter(
                    or_(
                        DBUserEntity.username.like(keyword),
                        DBUserEntity.email.like(keyword),
                    )
                ).offset(offset).limit(limit).all()
            users = [User.from_orm(db_user) for db_user in db_users]
        finally:
            db.close()
        return users

    def insert(self, user_data: UserData) -> User:
        db = Session(self.engine)
        new_user: User
        try:
            db_user = DBUserEntity(
                id=str(uuid.uuid4()),
                username=user_data.username,
                email=user_data.email,
                roles=user_data.roles,
                active=user_data.active,
                hashed_password=self._hash_password(user_data.password),
                full_name=user_data.full_name,
                created_at=datetime.datetime.utcnow()
            )
            db.add(db_user)
            db.commit()
            db.refresh(db_user) 
            new_user = User.from_orm(db_user)
        finally:
            db.close()
        return new_user

    def update(self, id: str, user_data: UserData) -> User:
        db = Session(self.engine)
        updated_user: User
        try:
            db_user = db.query(DBUserEntity).filter(DBUserEntity.id == id).first()
            if db_user is None:
                return None
            db_user.username = user_data.username
            db_user.email = user_data.email
            db_user.roles = user_data.roles
            db_user.active = user_data.active
            db_user.full_name = user_data.full_name
            db_user.updated_at = datetime.datetime.utcnow()
            if user_data.password:
                db_user.hashed_password = self._hash_password(user_data.password)
            db.add(db_user)
            db.commit()
            db.refresh(db_user) 
            updated_user = User.from_orm(db_user)
        finally:
            db.close()
        return updated_user

    def delete(self, id: str) -> User:
        db = Session(self.engine)
        deleted_user: User
        try:
            db_user = db.query(DBUserEntity).filter(DBUserEntity.id == id).first()
            if db_user is None:
                return None
            db.delete(db_user)
            db.commit()
            deleted_user = User.from_orm(db_user)
        finally:
            db.close()
        return deleted_user
