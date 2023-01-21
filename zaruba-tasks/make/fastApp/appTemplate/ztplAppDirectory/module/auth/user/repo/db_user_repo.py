from typing import Any, List, Optional, Mapping
from sqlalchemy.ext.hybrid import hybrid_property
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session, InstrumentedAttribute
from sqlalchemy import (
    and_, or_, Boolean, Column, DateTime, String, Text
)
from schema.user import User, UserWithoutPassword, UserData
from module.auth.user.repo.user_repo import UserRepo
from repo import Base, BaseMixin, DBRepo

import bcrypt
import jsons


class DBUserEntity(Base, BaseMixin):
    __tablename__ = "users"  # Note: 🤖 Don't delete this line
    username = Column(String(50), index=True, unique=True, nullable=False)
    email = Column(String(50), index=True, unique=True, nullable=True)
    phone_number = Column(String(20), index=True, unique=True, nullable=True)
    json_role_ids = Column(Text(), nullable=False, default='[]')
    json_permissions = Column(Text(), nullable=False, default='[]')
    active = Column(Boolean(), index=True, nullable=False, default=False)
    hashed_password = Column(String(60), index=False, nullable=False)
    full_name = Column(String(50), index=True, nullable=True)


class DBUserRepo(
    DBRepo[DBUserEntity, UserWithoutPassword, UserData],
    UserRepo
):
    schema_class = UserWithoutPassword
    db_entity_class = DBUserEntity

    def find_by_username(self, username: str) -> Optional[UserWithoutPassword]:
        db = self.create_db_sesion()
        try:
            search_filter = DBUserEntity.username == username
            return self.fetch_one_by_filter(db, search_filter)
        finally:
            db.close()

    def find_by_email(self, email: str) -> Optional[UserWithoutPassword]:
        db = self.create_db_sesion()
        try:
            search_filter = DBUserEntity.email == email
            return self.fetch_one_by_filter(db, search_filter)
        finally:
            db.close()

    def find_by_phone_number(
        self, phone_number: str
    ) -> Optional[UserWithoutPassword]:
        db = self.create_db_sesion()
        try:
            search_filter = DBUserEntity.phone_number == phone_number
            return self.fetch_one_by_filter(db, search_filter)
        finally:
            db.close()

    def find_by_identity_and_password(
        self, identity: str, password: str
    ) -> Optional[UserWithoutPassword]:
        db = self.create_db_sesion()
        try:
            search_filter = or_(
                and_(
                    DBUserEntity.username == identity,
                    DBUserEntity.username != '',
                    DBUserEntity.username is not None,
                ),
                and_(
                    DBUserEntity.email == identity,
                    DBUserEntity.email != '',
                    DBUserEntity.email is not None,
                ),
                and_(
                    DBUserEntity.phone_number == identity,
                    DBUserEntity.phone_number != '',
                    DBUserEntity.phone_number is not None,
                )
            )
            db_user = db.query(DBUserEntity).filter(search_filter).first()
            if (
                db_user is not None
                and self.is_valid_password(password, db_user.hashed_password)
            ):
                return self.from_db_entity_to_schema(db_user)
            return None
        finally:
            db.close()

    def get_keyword_fields(self) -> List[InstrumentedAttribute]:
        return [
            DBUserEntity.username,
            DBUserEntity.email,
            DBUserEntity.phone_number,
            DBUserEntity.full_name,
        ]

    def from_db_entity_to_schema(
        self, db_user: DBUserEntity
    ) -> UserWithoutPassword:
        user = super().from_db_entity_to_schema(db_user)
        user.permissions = jsons.loads(db_user.json_permissions)
        user.role_ids = jsons.loads(db_user.json_role_ids)
        return user

    def from_schema_data_to_db_entity_dict(
        self, user_data: UserData
    ) -> Mapping[str, Any]:
        db_user_dict = super().from_schema_data_to_db_entity_dict(user_data)
        db_user_dict['hashed_password'] = self.hash_password(
            user_data.password
        )
        db_user_dict['json_permissions'] = jsons.dumps(user_data.permissions)
        db_user_dict['json_role_ids'] = jsons.dumps(user_data.role_ids)
        return db_user_dict

    def hash_password(self, password: str) -> str:
        return bcrypt.hashpw(
            password.encode('utf-8'), bcrypt.gensalt()
        ).decode('utf-8')

    def is_valid_password(self, password: str, hashed_password: str) -> bool:
        return bcrypt.checkpw(
            password.encode('utf-8'),
            hashed_password.encode('utf-8')
        )
