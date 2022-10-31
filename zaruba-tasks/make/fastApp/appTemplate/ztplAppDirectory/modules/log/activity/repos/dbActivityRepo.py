from typing import Any, List, Optional
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.activity import Activity, ActivityData
from modules.log.activity.repos.activityRepo import ActivityRepo
from repos import Base

import uuid
import datetime
import jsons

# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class DBActivityEntity(Base):
    __tablename__ = "activities"
    id = Column(String(36), primary_key=True, index=True)
    user_id = Column(String(36), index=True)
    activity = Column(String(255), index=True)
    object = Column(String(255), index=True, nullable=True)
    row_id = Column(String(255), index=True)
    json_row = Column(String(255), index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow) # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
    created_by = Column(String(36), nullable=True)
    updated_at = Column(DateTime, nullable=True)
    updated_by = Column(String(36), nullable=True)


# Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
class DBActivityRepo(ActivityRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)


    def _get_keyword_filter(self, keyword: str) -> str:
        return '%{}%'.format(keyword) if keyword != '' else '%'


    def _from_db_result(self, db_result: Any) -> Activity:
        activity = Activity.from_orm(db_result)
        activity.row = json.loads(db_result.json_row)
        return activity


    def find_by_id(self, id: str) -> Optional[Activity]:
        db = Session(self.engine)
        activity: Activity
        try:
            db_activity = db.query(DBActivityEntity).filter(DBActivityEntity.id == id).first()
            if db_activity is None:
                return None
            activity = Activity.from_orm(db_activity)
        finally:
            db.close()
        return activity


    def find(self, keyword: str, limit: int, offset: int) -> List[Activity]:
        db = Session(self.engine)
        activities: List[Activity] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_activities = db.query(DBActivityEntity).filter(DBActivityEntity.activity.like(keyword_filter)).offset(offset).limit(limit).all()
            activities = [Activity.from_orm(db_result) for db_result in db_activities]
        finally:
            db.close()
        return activities


    def count(self, keyword: str) -> int:
        db = Session(self.engine)
        activity_count = 0
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            activity_count = db.query(DBActivityEntity).filter(DBActivityEntity.activity.like(keyword_filter)).count()
        finally:
            db.close()
        return activity_count


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def insert(self, activity_data: ActivityData) -> Optional[Activity]:
        db = Session(self.engine)
        activity: Activity
        try:
            new_activity_id = str(uuid.uuid4())
            db_activity = DBActivityEntity(
                id=new_activity_id,
                user_id=activity_data.user_id,
                activity=activity_data.activity,
                object=activity_data.object,
                row_id=activity_data.row_id,
                json_row=jsons.dumps(activity_data.row),
                created_at=datetime.datetime.utcnow(), # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
                created_by=activity_data.created_by,
                updated_at=datetime.datetime.utcnow(),
                updated_by=activity_data.updated_by,
            )
            db.add(db_activity)
            db.commit()
            db.refresh(db_activity) 
            activity = Activity.from_orm(db_activity)
        finally:
            db.close()
        return activity


    # Note: 💀 Don't delete the following line, Zaruba use it for pattern matching
    def update(self, id: str, activity_data: ActivityData) -> Optional[Activity]:
        db = Session(self.engine)
        activity: Activity
        try:
            db_activity = db.query(DBActivityEntity).filter(DBActivityEntity.id == id).first()
            if db_activity is None:
                return None
            db_activity.user_id = activity_data.user_id
            db_activity.activity = activity_data.activity
            db_activity.object = activity_data.object
            db_activity.row_id = activity_data.row_id
            db_activity.row = jsons.dumps(activity_data.row)
            db_activity.updated_at = datetime.datetime.utcnow() # Note: 💀 Don't delete this line, Zaruba use it for pattern matching
            db_activity.updated_by = activity_data.updated_by
            db.add(db_activity)
            db.commit()
            db.refresh(db_activity) 
            activity = Activity.from_orm(db_activity)
        finally:
            db.close()
        return activity


    def delete(self, id: str) -> Optional[Activity]:
        db = Session(self.engine)
        activity: Activity
        try:
            db_activity = db.query(DBActivityEntity).filter(DBActivityEntity.id == id).first()
            if db_activity is None:
                return None
            db.delete(db_activity)
            db.commit()
            activity = Activity.from_orm(db_activity)
        finally:
            db.close()
        return activity

