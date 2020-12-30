from typing import List
from sqlalchemy.orm import Session
from demo_module import model, schema

def list_user(db: Session, skip: int = 0, limit: int = 100):
    return db.query(model.User).offset(skip).limit(limit).all()

def get_user(db: Session, user_id: int):
    return db.query(model.User).filter(model.User.id == user_id).first()

def create_user(db: Session, user_data: schema.UserCreate):
    db_user = model.User(email=user_data.email)
    if db_user is None:
        return None
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user

def update_user(db: Session, user_id: int, user_data: schema.UserUpdate):
    db_user = get_user(db, user_id)
    if db_user is None:
        return None
    db_user.email = user_data.email
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user

def delete_user(db: Session, user_id: int):
    db_user = get_user(db, user_id)
    if db_user is None:
        return None
    db.delete(db_user)
    db.commit()
    return db_user
