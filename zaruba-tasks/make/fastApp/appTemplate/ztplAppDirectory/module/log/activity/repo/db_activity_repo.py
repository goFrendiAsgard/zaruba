from typing import Any, List, Optional, Mapping
from sqlalchemy import (
    or_, Boolean, Column, DateTime, ForeignKey, Integer, String, Text
)
from sqlalchemy.orm.attributes import InstrumentedAttribute
from schema.activity import Activity, ActivityData
from module.log.activity.repo.activity_repo import ActivityRepo
from repo import Base, BaseMixin, DBRepo

import jsons


# Note: 🤖 Don't delete the following statement
class DBActivityEntity(Base, BaseMixin):
    __tablename__ = "activities"  # Note: 🤖 Don't delete this line
    user_id = Column(String(36), index=True, nullable=False)
    activity = Column(String(255), index=True, nullable=False)
    object = Column(String(255), index=True, nullable=True)
    row_id = Column(String(255), index=True, nullable=True)
    json_row = Column(Text(), index=False, nullable=True)


class DBActivityRepo(
    DBRepo[DBActivityEntity, Activity, ActivityData],
    ActivityRepo
):

    schema_class = Activity
    db_entity_class = DBActivityEntity

    def get_keyword_fields(self) -> List[InstrumentedAttribute]:
        return [
            DBActivityEntity.activity,
            DBActivityEntity.object,
            DBActivityEntity.row_id,
            DBActivityEntity.user_id
        ]

    def from_schema_data_to_db_entity_dict(
        self, activity_data: ActivityData
    ) -> Mapping[str, Any]:
        activity_dict = super().from_schema_data_to_db_entity_dict(
            activity_data
        )
        activity_dict['json_row'] = jsons.dumps(activity_data.row)
        return activity_dict

    def from_db_entity_to_schema(
        self, db_activity: DBActivityEntity
    ) -> Activity:
        activity = super().from_db_entity_to_schema(db_activity)
        activity.row = jsons.loads(db_activity.json_row)
        return activity